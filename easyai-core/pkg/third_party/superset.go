package third_party

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

func main() {

	url := "https://superset-ml.ssr.mihoyo.com/login/"
	method := "POST"

	payload := strings.NewReader("username=shuai.yang%40mihoyo.com&password=bb469afc")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "superset-ml.ssr.mihoyo.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://superset-ml.ssr.mihoyo.com/login/")
	req.Header.Add("Cookie", "session=.eJyVUVtuwzAMu4u_uzqOYsvKVYoh8ENagqV1EKcfxbC7z8V2gBX6EiWKJPSlJtm5zmqUsFY-qWnJalTM3vsBOxvERbTkXIhIZui4FUrCBpNhBivW5CHFThybznGEAVP2IKGhfjAWewNR2BGBl4DgbN-BjzYF6sgiBwOGgARMJHDU20iOnTqpVHeZjvLJt-Yn9m2Ts1gSywkCI3FK6Cmw84yQBWOwOTbeWlJY-Znh1rotfPA0L_Uo-0ONFzUfxzZqXe8b75WPt-t6rnU_X5e5PMo5lavey8pVc14OPbQDLxH6fxHutU302kzpFwRe3f9L8H5ST73ft_bq-weSHaIO.YzQd_g.zGVBPmhCIMCZMHxmkXzkBJ7QS7k")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func LoginSuperset(host, url, username, password string) (http.Header, []byte, error) {

	client := resty.New()

	headers := map[string]string{
		"authority":     _hostOfURL(url),
		"accept":        "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"cache-control": "no-cache",
		"content-type":  "application/x-www-form-urlencoded",
		"pragma":        "no-cache",
		"referer":       url,
	}

	payload := map[string]string{
		"username": username,
		"password": password,
	}
	req, err := client.R().SetFormData(payload).SetHeaders(headers).Post(url)

	if err != nil {
		return nil, nil, err
	}
	return req.Header(), req.Body(), nil
}
