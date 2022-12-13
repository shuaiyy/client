package sdk

import (
	"fmt"
	"net/url"
)

// Header for http request
type Header map[string]string

// Query  for http request
type Query struct {
	Prefix    *string
	StartKey  *string
	NextToken *string
	Limit     *int32
	Tags      map[string]string
}

// Client defines seelie client
type Client struct {
	Config  *Config
	Connect *Connection
}

// ServiceInput for http request
type ServiceInput interface {
	GetQueryParams() url.Values
	GetPath() string
	GetHeaders() Header
	GetPayload() interface{}
	Validate() error
}

// NewClient new seelie client
func NewClient(host string, port int64, protocol, apiPrefix, token string, debugHTTP ...bool) (*Client, error) {
	config := NewConfig()
	config.Endpoint = fmt.Sprintf("%s://%s:%d", protocol, host, port)
	if apiPrefix != "" {
		config.APIPrefix = apiPrefix
	}
	baseURL := config.Endpoint
	config.Token = token
	connect := NewConnection(baseURL, debugHTTP...)
	client := &Client{config, connect}
	return client, nil
}

type baseInput struct{}

func (b baseInput) GetQueryParams() url.Values {
	return url.Values{}
}

func (b baseInput) GetPath() string {
	return ""
}

func (b baseInput) GetHeaders() Header {
	return map[string]string{}
}

func (b baseInput) GetPayload() interface{} {
	return nil
}

func (b baseInput) Validate() error {
	return nil
}
