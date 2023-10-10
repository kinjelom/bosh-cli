package completion

import (
	"encoding/json"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/peterbourgon/diskv"
	"time"
)

type CacheItem struct {
	Timestamp time.Time `json:"timestamp"`
	Values    []string  `json:"values"`
}

type CompleteCache struct {
	store            *diskv.Diskv
	cacheKey         string
	expirationPeriod time.Duration
	logger           boshlog.Logger
	logTag           string
}

func NewCompleteCache(logger boshlog.Logger, queryCacheDirPath string, cacheKey string) *CompleteCache {
	c := &CompleteCache{
		store: diskv.New(diskv.Options{
			BasePath: queryCacheDirPath,
		}),
		cacheKey:         cacheKey,
		expirationPeriod: 60 * time.Second,
		logger:           logger,
		logTag:           "completion.CompleteCache",
	}
	return c
}

func (c *CompleteCache) PutValues(values []string) error {
	item := CacheItem{
		Timestamp: time.Now(),
		Values:    values,
	}
	data, err := json.Marshal(item)
	if err != nil {
		c.logger.Debug(c.logTag, "marshal values error: %v", err)
		return err
	}
	err = c.store.Write(c.cacheKey, data)
	if err != nil {
		c.logger.Debug(c.logTag, "write cache error: %v", err)
	}
	return err
}

func (c *CompleteCache) GetValues() (values []string, valid bool, err error) {
	data, err := c.store.Read(c.cacheKey)
	if err != nil {
		if err != nil {
			c.logger.Debug(c.logTag, "read cache error: %v", err)
		}
		return nil, false, err
	}

	var item CacheItem
	if err := json.Unmarshal(data, &item); err != nil {
		if err != nil {
			c.logger.Debug(c.logTag, "unmarshal cached values error: %v", err)
		}
		return nil, false, err
	}
	valid = time.Since(item.Timestamp) < c.expirationPeriod
	if !valid {
		_ = c.store.Erase(c.cacheKey)
	}
	return item.Values, valid, nil
}
