package http

import (
	"fmt"
	"strings"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-utils/httpclient"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"

	"github.com/cloudfoundry/bosh-agent/v2/agentclient"
	"github.com/cloudfoundry/bosh-agent/v2/agentclient/applyspec"
)

type AgentClient struct {
	AgentRequest        agentRequest
	getTaskDelay        time.Duration
	toleratedErrorCount int
	logger              boshlog.Logger
	logTag              string
}

func NewAgentClient(
	endpoint string,
	directorID string,
	getTaskDelay time.Duration,
	toleratedErrorCount int,
	httpClient *httpclient.HTTPClient,
	logger boshlog.Logger,
) agentclient.AgentClient {
	// if this were NATS, we would need the agentID, but since it's http, the endpoint is unique to the agent
	agentEndpoint := fmt.Sprintf("%s/agent", endpoint)
	agentRequest := agentRequest{
		directorID: directorID,
		endpoint:   agentEndpoint,
		httpClient: httpClient,
	}
	return &AgentClient{
		AgentRequest:        agentRequest,
		getTaskDelay:        getTaskDelay,
		toleratedErrorCount: toleratedErrorCount,
		logger:              logger,
		logTag:              "httpAgentClient",
	}
}

func (c *AgentClient) Ping() (string, error) {
	var response SimpleTaskResponse
	err := c.AgentRequest.Send("ping", []interface{}{}, &response)
	if err != nil {
		return "", bosherr.WrapError(err, "Sending ping to the agent")
	}

	return response.Value, nil
}

func (c *AgentClient) Stop() error {
	_, err := c.SendAsyncTaskMessage("stop", []interface{}{})
	return err
}

func (c *AgentClient) Drain(drainType string) (int64, error) {
	responseRaw, err := c.SendAsyncTaskMessage("drain", []interface{}{drainType, map[string]interface{}{}})
	if err != nil {
		return 0, err
	}

	responseValue := int64(responseRaw.(float64))

	return responseValue, err
}

func (c *AgentClient) Apply(spec applyspec.ApplySpec) error {
	_, err := c.SendAsyncTaskMessage("apply", []interface{}{spec})
	return err
}

func (c *AgentClient) Start() error {
	var response SimpleTaskResponse
	err := c.AgentRequest.Send("start", []interface{}{}, &response)
	if err != nil {
		return bosherr.WrapError(err, "Starting agent services")
	}

	if response.Value != "started" {
		return bosherr.Errorf("Failed to start agent services with response: '%s'", response.Exception.Message)
	}

	return nil
}

func (c *AgentClient) GetState() (agentclient.AgentState, error) {
	var response StateResponse

	getStateRetryable := boshretry.NewRetryable(func() (bool, error) {
		err := c.AgentRequest.Send("get_state", []interface{}{}, &response)
		if err != nil {
			return true, bosherr.WrapError(err, "Sending get_state to the agent")
		}
		return false, nil
	})

	attemptRetryStrategy := boshretry.NewAttemptRetryStrategy(c.toleratedErrorCount+1, c.getTaskDelay, getStateRetryable, c.logger)
	err := attemptRetryStrategy.Try()
	if err != nil {
		return agentclient.AgentState{}, bosherr.WrapError(err, "Sending get_state to the agent")
	}

	agentState := agentclient.AgentState{
		JobState:     response.Value.JobState,
		NetworkSpecs: response.Value.NetworkSpecs,
	}

	return agentState, err
}

func (c *AgentClient) ListDisk() ([]string, error) {
	var response ListResponse
	err := c.AgentRequest.Send("list_disk", []interface{}{}, &response)
	if err != nil {
		return []string{}, bosherr.WrapError(err, "Sending 'list_disk' to the agent")
	}

	return response.Value, nil
}

func (c *AgentClient) MountDisk(diskCID string) error {
	_, err := c.SendAsyncTaskMessage("mount_disk", []interface{}{diskCID})
	return err
}

func (c *AgentClient) UnmountDisk(diskCID string) error {
	_, err := c.SendAsyncTaskMessage("unmount_disk", []interface{}{diskCID})
	return err
}

func (c *AgentClient) MigrateDisk() error {
	_, err := c.SendAsyncTaskMessage("migrate_disk", []interface{}{})
	return err
}

func (c *AgentClient) RunScript(scriptName string, options map[string]interface{}) error {
	_, err := c.SendAsyncTaskMessage("run_script", []interface{}{scriptName, options})

	if err != nil && strings.Contains(err.Error(), "unknown message") {
		// ignore 'unknown message' errors for backwards compatibility with older stemcells
		c.logger.Warn(c.logTag, "Ignoring run_script 'unknown message' error from the agent: %s. Received while trying to run: %s", err.Error(), scriptName)
		return nil
	}

	return err
}

func (c *AgentClient) SetUpSSH(user string, publicKey string) (agentclient.SSHResult, error) {
	var response SSHResponse
	sshParams := map[string]string{"user": user, "public_key": publicKey}
	err := c.AgentRequest.Send("ssh", []interface{}{"setup", sshParams}, &response)

	if err != nil {
		return agentclient.SSHResult{}, err
	}

	if response.Value.Status != "success" {
		return agentclient.SSHResult{}, bosherr.Errorf("Unable to setup SSH account with the agent, status was: %s", response.Value.Status)
	}

	return agentclient.SSHResult{
		Command:       response.Value.Command,
		Status:        response.Value.Status,
		Ip:            response.Value.Ip,
		HostPublicKey: response.Value.HostPublicKey,
	}, nil
}

func (c *AgentClient) CleanUpSSH(user string) (agentclient.SSHResult, error) {
	var response SSHResponse
	sshParams := map[string]string{"user_regex": "^" + user}
	err := c.AgentRequest.Send("ssh", []interface{}{"cleanup", sshParams}, &response)

	if err != nil {
		return agentclient.SSHResult{}, err
	}

	if response.Value.Status != "success" {
		return agentclient.SSHResult{}, bosherr.Errorf("Unable to cleanup SSH account with the agent, status was: %s", response.Value.Status)
	}

	return agentclient.SSHResult{
		Command: response.Value.Command,
		Status:  response.Value.Status,
	}, nil
}

func (c *AgentClient) BundleLogs(owningUser string, logType string, filters []string) (agentclient.BundleLogsResult, error) {
	var response BundleLogsResponse
	err := c.AgentRequest.Send("bundle_logs", []interface{}{map[string]interface{}{
		"owning_user": owningUser,
		"log_type":    logType,
		"filters":     filters,
	}}, &response)
	if err != nil {
		return agentclient.BundleLogsResult{}, err
	}

	return agentclient.BundleLogsResult{
		LogsTarPath:  response.Value.LogsTarPath,
		SHA512Digest: response.Value.SHA512Digest,
	}, nil
}

func (c *AgentClient) RemoveFile(path string) error {
	var response SimpleTaskResponse
	err := c.AgentRequest.Send("remove_file", []interface{}{path}, &response)
	if err != nil {
		return err
	}

	return nil
}

func (c *AgentClient) CompilePackage(packageSource agentclient.BlobRef, compiledPackageDependencies []agentclient.BlobRef) (compiledPackageRef agentclient.BlobRef, err error) {
	dependencies := make(map[string]BlobRef, len(compiledPackageDependencies))
	for _, dependency := range compiledPackageDependencies {
		dependencies[dependency.Name] = BlobRef{
			Name:        dependency.Name,
			Version:     dependency.Version,
			SHA1:        dependency.SHA1,
			BlobstoreID: dependency.BlobstoreID,
		}
	}

	args := []interface{}{
		packageSource.BlobstoreID,
		packageSource.SHA1,
		packageSource.Name,
		packageSource.Version,
		dependencies,
	}

	responseRaw, err := c.SendAsyncTaskMessage("compile_package", args)
	responseValue, ok := responseRaw.(map[string]interface{})
	if !ok {
		c.logger.Warn(c.logTag, "Unable to parse compile_package response value: %#v", responseRaw)
	}
	if err != nil {
		return agentclient.BlobRef{}, bosherr.WrapError(err, "Sending 'compile_package' to the agent")
	}

	result, ok := responseValue["result"].(map[string]interface{})
	if !ok {
		return agentclient.BlobRef{}, bosherr.Errorf("Unable to parse 'compile_package' response from the agent: %#v", responseValue)
	}

	sha1, ok := result["sha1"].(string)
	if !ok {
		return agentclient.BlobRef{}, bosherr.Errorf("Unable to parse 'compile_package' response from the agent: %#v", responseValue)
	}

	blobstoreID, ok := result["blobstore_id"].(string)
	if !ok {
		return agentclient.BlobRef{}, bosherr.Errorf("Unable to parse 'compile_package' response from the agent: %#v", responseValue)
	}

	compiledPackageRef = agentclient.BlobRef{
		Name:        packageSource.Name,
		Version:     packageSource.Version,
		SHA1:        sha1,
		BlobstoreID: blobstoreID,
	}

	return compiledPackageRef, nil
}

func (c *AgentClient) DeleteARPEntries(ips []string) error {
	return c.AgentRequest.Send("delete_arp_entries", []interface{}{map[string][]string{"ips": ips}}, &TaskResponse{})
}

func (c *AgentClient) SyncDNS(blobID, sha1 string, version uint64) (string, error) {
	var response SyncDNSResponse
	err := c.AgentRequest.Send("sync_dns", []interface{}{blobID, sha1, version}, &response)
	if err != nil {
		return "", bosherr.WrapError(err, "Sending 'sync_dns' to the agent")
	}

	return response.Value, nil
}

func (c *AgentClient) SendAsyncTaskMessage(method string, arguments []interface{}) (value interface{}, err error) {
	var response TaskResponse
	err = c.AgentRequest.Send(method, arguments, &response)
	if err != nil {
		return value, bosherr.WrapErrorf(err, "Sending '%s' to the agent", method)
	}

	agentTaskID, err := response.TaskID()
	if err != nil {
		return value, bosherr.WrapError(err, "Getting agent task id")
	}

	sendErrors := 0
	getTaskRetryable := boshretry.NewRetryable(func() (bool, error) {
		var response TaskResponse
		err = c.AgentRequest.Send("get_task", []interface{}{agentTaskID}, &response)
		if err != nil {
			sendErrors++
			shouldRetry := sendErrors <= c.toleratedErrorCount
			err = bosherr.WrapError(err, "Sending 'get_task' to the agent")
			c.logger.Debug(c.logTag, "Error occurred sending get_task. Error retry %d of %d: %s", sendErrors, c.toleratedErrorCount, err.Error())
			return shouldRetry, err
		}
		sendErrors = 0

		c.logger.Debug(c.logTag, "get_task response value: %#v", response.Value)

		taskState, err := response.TaskState()
		if err != nil {
			return false, bosherr.WrapError(err, "Getting task state")
		}

		if taskState != "running" {
			value = response.Value
			return false, nil
		}

		return true, bosherr.Errorf("Task %s is still running", method)
	})

	getTaskRetryStrategy := boshretry.NewUnlimitedRetryStrategy(c.getTaskDelay, getTaskRetryable, c.logger)
	// cannot call getTaskRetryStrategy.Try in the return statement due to gccgo
	// execution order issues: https://code.google.com/p/go/issues/detail?id=8698&thanks=8698&ts=1410376474
	err = getTaskRetryStrategy.Try()
	return value, err
}

func (c *AgentClient) AddPersistentDisk(diskCID string, diskHints interface{}) error {
	_, err := c.SendAsyncTaskMessage("add_persistent_disk", []interface{}{diskCID, diskHints})
	return err
}

func (c *AgentClient) RemovePersistentDisk(diskCID string) error {
	_, err := c.SendAsyncTaskMessage("remove_persistent_disk", []interface{}{diskCID})
	return err
}
