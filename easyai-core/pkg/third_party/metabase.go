package third_party

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strings"
)

// TODO  nginx ingress对应的公网IP无法在 aliyun-dev内访问
// 需要使用 ml-metabase.ml-system.cluster.local:80 作为 https://ml-metabase.ml.ssr.mihoyo.com的实际请求地址
// 不知道咋实现，可能需要通过改 http transport 的dial实现

func LoginMetabase(host, url, username, password string) (http.Header, []byte, error) {

	client := resty.New()
	headers := map[string]string{
		"Host":            host,
		"authority":       host,
		"accept":          "application/json",
		"accept-language": "zh-CN,zh;q=0.9",
		"cache-control":   "no-cache",
		"content-type":    "application/json",
	}
	payload := fmt.Sprintf(`{"username":"%s","password":"%s","remember":true}`, username, password)
	client.SetDebug(true)
	client.SetScheme("https")
	//client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
	//	r.RawRequest.Host = host
	//	return nil
	//})

	req, err := client.R().SetBody([]byte(payload)).SetHeaders(headers).SetSRV(&resty.SRVRecord{
		Service: "",
		Domain:  host,
	}).Post(url)

	if err != nil {
		return nil, nil, err
	}
	h := req.Header()
	b := req.Body()
	if !strings.Contains(string(b), "id:") {
		return nil, nil, fmt.Errorf("login metabase %s failed, maybe user<%s> not register, body: %s", url, username, b)
	}
	return h, b, nil
}

func _hostOfURL(url string) string {
	if strings.Contains(url, "://") {
		url = strings.Split(url, "://")[1]
	}
	return strings.Split(url, "/")[0]
}
