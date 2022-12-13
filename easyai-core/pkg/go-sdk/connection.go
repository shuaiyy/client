package sdk

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

// default parameter
const (
	RequestTimeout     = 60
	defaultIdleTimeout = 50 * time.Second
)

// Connection with fc
type Connection struct {
	Timeout uint // 超时时间，默认60秒
	Debug   bool // for http debug
	Client  *resty.Client
}

// NewConnection get default connection
func NewConnection(baseURL string, enableDebug ...bool) *Connection {

	connection := &Connection{
		Timeout: RequestTimeout,
		Debug:   len(enableDebug) > 0 && enableDebug[0],
	}
	clt := resty.New()
	clt.SetTransport(&http.Transport{
		IdleConnTimeout: defaultIdleTimeout,
	})
	clt.SetTimeout(time.Duration(connection.Timeout) * time.Second)
	clt.SetDebug(connection.Debug)
	clt.SetBaseURL(baseURL)
	connection.Client = clt
	return connection
}

// PrepareRequest prepare http request
func (conn *Connection) PrepareRequest(postBody interface{},
	headerParams map[string]string,
	queryParams url.Values) *resty.Request {

	request := conn.Client.R()
	if postBody != nil {
		request.SetBody(postBody)
	}

	// add header parameter, if any
	if len(headerParams) > 0 {
		request.SetHeaders(headerParams)
	}

	// add query parameter, if any
	if len(queryParams) > 0 {
		request.SetQueryParamsFromValues(queryParams)
	}
	return request
}

// SendRequest send http request
func (conn *Connection) SendRequest(path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values) (*resty.Response, error) {

	request := conn.PrepareRequest(postBody, headerParams, queryParams)

	switch strings.ToUpper(method) {
	case http.MethodGet:
		response, err := request.Get(path)
		return response, err
	case http.MethodPost:
		response, err := request.Post(path)
		return response, err
	case http.MethodPut:
		response, err := request.Put(path)
		return response, err
	case http.MethodDelete:
		response, err := request.Delete(path)
		return response, err
	}
	return nil, fmt.Errorf("invalid method %v", method)
}

// ParameterToString serialize parameters
func ParameterToString(obj interface{}, collectionFormat string) string {
	if reflect.TypeOf(obj).String() == "[]string" {
		switch collectionFormat {
		case "pipes":
			return strings.Join(obj.([]string), "|")
		case "ssv":
			return strings.Join(obj.([]string), " ")
		case "tsv":
			return strings.Join(obj.([]string), "\t")
		case "csv":
			return strings.Join(obj.([]string), ",")
		}
	}
	return fmt.Sprintf("%v", obj)
}
