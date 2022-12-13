package sdk

import (
	"fmt"
	"net/url"
	"path/filepath"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"
)

// JobStopInput  stop job
type JobStopInput struct {
	jobID uint32
}

// NewJobStopInput new
func NewJobStopInput() *JobStopInput {
	return &JobStopInput{}
}

// WithJobID set args
func (js *JobStopInput) WithJobID(id uint32) *JobStopInput {
	js.jobID = id
	return js
}

// GetQueryParams set args
func (js *JobStopInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

// GetPath return api path
func (js *JobStopInput) GetPath() string {
	return fmt.Sprintf("/jobs/%d/stop", js.jobID)
}

// GetHeaders return http header
func (js *JobStopInput) GetHeaders() Header {
	return map[string]string{}
}

// GetPayload return http body
func (js *JobStopInput) GetPayload() interface{} {
	return nil
}

// Validate validate
func (js *JobStopInput) Validate() error {
	if js.jobID <= 0 {
		return fmt.Errorf("invalid JobStopInput: %+v", js)
	}
	return nil
}

// StopJob by id
func (clt *Client) StopJob(input *JobStopInput) (string, error) {
	if err := input.Validate(); err != nil {
		return "", err
	}
	path := filepath.Join(clt.Config.APIPrefix, input.GetPath())
	headers := input.GetHeaders()
	headers[HTTPHeaderToken] = clt.Config.Token
	response, err := clt.Connect.SendRequest(path, "POST", input.GetPayload(), headers, input.GetQueryParams())
	if err != nil {
		return "", err
	}
	var resp = struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	}{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		return "", err
	}
	if resp.Code != 200 {
		return "", fmt.Errorf("got error code: %d, message: %s", resp.Code, resp.Message)
	}
	return resp.Message, nil
}
