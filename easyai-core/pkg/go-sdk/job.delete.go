package sdk

import (
	"fmt"
	"net/url"
	"path/filepath"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"
)

// JobDeleteInput  获取
type JobDeleteInput struct {
	jobID uint32
}

// NewJobDeleteInput new
func NewJobDeleteInput() *JobDeleteInput {
	return &JobDeleteInput{}
}

// WithJobID set args
func (jg *JobDeleteInput) WithJobID(id uint32) *JobDeleteInput {
	jg.jobID = id
	return jg
}

// GetQueryParams set args
func (jg *JobDeleteInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

// GetPath return api path
func (jg *JobDeleteInput) GetPath() string {
	return fmt.Sprintf("/jobs/%d", jg.jobID)
}

// GetHeaders return http header
func (jg *JobDeleteInput) GetHeaders() Header {
	return map[string]string{}
}

// GetPayload return http body
func (jg *JobDeleteInput) GetPayload() interface{} {
	return nil
}

// Validate validate
func (jg *JobDeleteInput) Validate() error {
	if jg.jobID <= 0 {
		return fmt.Errorf("invalid JobDeleteInput: %+v", jg)
	}
	return nil
}

// DeleteJob by id
func (clt *Client) DeleteJob(input *JobDeleteInput) (string, error) {
	if err := input.Validate(); err != nil {
		return "", err
	}
	path := filepath.Join(clt.Config.APIPrefix, input.GetPath())
	headers := input.GetHeaders()
	headers[HTTPHeaderToken] = clt.Config.Token
	response, err := clt.Connect.SendRequest(path, "DELETE", input.GetPayload(), headers, input.GetQueryParams())
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
