package sdk

import (
	"fmt"
	"path/filepath"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/api/schema"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"
)

// ClusterConfigsInput  获取
type ClusterConfigsInput struct{ baseInput }

// NewClusterConfigsInput new
func NewClusterConfigsInput() *ClusterConfigsInput {
	return &ClusterConfigsInput{}
}

// GetPath return api path
func (jg *ClusterConfigsInput) GetPath() string {
	return "/clusters/configs"
}

// GetClusterConfigs return all cluster configs
func (clt *Client) GetClusterConfigs(input *ClusterConfigsInput) (map[string]*schema.ClusterConfig, error) {
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
		Code    int64                            `json:"code"`
		Message string                           `json:"message"`
		Data    map[string]*schema.ClusterConfig `json:"data"`
	}{}
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf("got error code: %d, message: %s", resp.Code, resp.Message)
	}
	return resp.Data, nil
}
