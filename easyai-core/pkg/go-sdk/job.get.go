package sdk

import (
	"fmt"
	"net/url"
	"path/filepath"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/api/schema"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"
)

// JobGetInput  获取
type JobGetInput struct {
	jobID uint32
}

// NewJobGetInput new
func NewJobGetInput() *JobGetInput {
	return &JobGetInput{}
}

// WithJobID set args
func (jg *JobGetInput) WithJobID(id uint32) *JobGetInput {
	jg.jobID = id
	return jg
}

// GetQueryParams set args
func (jg *JobGetInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

// GetPath return api path
func (jg *JobGetInput) GetPath() string {
	return fmt.Sprintf("/jobs/%d", jg.jobID)
}

// GetHeaders return http header
func (jg *JobGetInput) GetHeaders() Header {
	return map[string]string{}
}

// GetPayload return http body
func (jg *JobGetInput) GetPayload() interface{} {
	return nil
}

// Validate validate
func (jg *JobGetInput) Validate() error {
	if jg.jobID <= 0 {
		return fmt.Errorf("invalid JobGetInput: %+v", jg)
	}
	return nil
}

// GetJob return a job
func (clt *Client) GetJob(input *JobGetInput) (*schema.Job, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	path := filepath.Join(clt.Config.APIPrefix, input.GetPath())
	headers := input.GetHeaders()
	headers[HTTPHeaderToken] = clt.Config.Token
	response, err := clt.Connect.SendRequest(path, "GET", input.GetPayload(), headers, input.GetQueryParams())
	if err != nil {
		return nil, err
	}
	var resp = struct {
		Code    int64      `json:"code"`
		Message string     `json:"message"`
		Data    schema.Job `json:"data"`
	}{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf("got error code: %d, message: %s", resp.Code, resp.Message)
	}
	return &resp.Data, nil
}
