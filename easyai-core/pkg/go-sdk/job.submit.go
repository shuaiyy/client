package sdk

import (
	"fmt"
	"net/url"
	"path/filepath"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/api/schema"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"
)

// JobSubmitInput  stop job
type JobSubmitInput struct {
	job schema.Job
}

// NewJobSubmitInput new
func NewJobSubmitInput() *JobSubmitInput {
	return &JobSubmitInput{}
}

// WithName set args
func (js *JobSubmitInput) WithName(n string) *JobSubmitInput {
	js.job.Name = n
	return js
}

// WithNamespace set args
func (js *JobSubmitInput) WithNamespace(v string) *JobSubmitInput {
	js.job.Namespace = v
	return js
}

// WithCluster set args
func (js *JobSubmitInput) WithCluster(v string) *JobSubmitInput {
	js.job.Cluster = v
	return js
}

// WithJob set args
func (js *JobSubmitInput) WithJob(j schema.Job) *JobSubmitInput {
	js.job = j
	return js
}

// GetQueryParams set args
func (js *JobSubmitInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

// GetPath return api path
func (js *JobSubmitInput) GetPath() string {
	return "/jobs"
}

// GetHeaders return http header
func (js *JobSubmitInput) GetHeaders() Header {
	return map[string]string{"Content-Type": "application/json"}
}

// GetPayload return http body
func (js *JobSubmitInput) GetPayload() interface{} {
	return js.job
}

// Validate validate
func (js *JobSubmitInput) Validate() error {
	if js.job.Image == "" {
		return fmt.Errorf("job.image cannot be empty")
	}
	if js.job.Name == "" {
		return fmt.Errorf("job.name cannot be empty")
	}
	if js.job.Namespace == "" {
		return fmt.Errorf("job.namespace cannot be empty")
	}
	if js.job.Cluster == "" {
		return fmt.Errorf("job.cluster cannot be empty")
	}
	if js.job.CPU <= 0 {
		return fmt.Errorf("job.cpu must bigger than zero, > 0")
	}
	if js.job.Memory <= 0 {
		return fmt.Errorf("job.memory must bigger than zero, > 0")
	}
	if js.job.GPU < 0 {
		return fmt.Errorf("job.gpu must no less than zero, >= 0")
	}
	if js.job.WorkerCount < 0 {
		return fmt.Errorf("job.worker_count must bigger than 0, > 0")
	}
	if js.job.EntrypointType == "" {
		return fmt.Errorf("job.entypoint_type cannot be empty")
	}
	return nil
}

// SubmitJob by
func (clt *Client) SubmitJob(input *JobSubmitInput) (uint32, error) {
	if err := input.Validate(); err != nil {
		return 0, fmt.Errorf("invalid input args, %+v", err)
	}
	path := filepath.Join(clt.Config.APIPrefix, input.GetPath())
	headers := input.GetHeaders()
	headers[HTTPHeaderToken] = clt.Config.Token
	response, err := clt.Connect.SendRequest(path, "POST", input.GetPayload(), headers, input.GetQueryParams())
	if err != nil {
		return 0, err
	}
	var resp = struct {
		Code    int64           `json:"code"`
		Message string          `json:"message"`
		Data    schema.IDResult `json:"data"`
	}{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		return 0, err
	}
	if resp.Code != 200 {
		return 0, fmt.Errorf("got error code: %d, message: %s", resp.Code, resp.Message)
	}
	return resp.Data.ID, nil
}
