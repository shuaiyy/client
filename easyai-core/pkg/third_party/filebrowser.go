package third_party

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/conv"
	"strings"
)

// FileBrowserClient is a client for the FileBrowser API
// only works for seelie customed file browser: nginx will put cookie to proxy header
type FileBrowserClient struct {
	url        string
	adminUser  string
	adminToken string
	client     *resty.Client
}

// NewFileBrowserClient creates a new FileBrowserClient
func NewFileBrowserClient(url string, adminUser string) (*FileBrowserClient, error) {
	clt := resty.New()
	url = strings.TrimSuffix(url, "/")
	loginUrl := fmt.Sprintf("%s/api/login", url)
	payload := map[string]string{"username": "", "password": "", "recaptcha": ""}
	resp, err := clt.R().SetCookie(&http.Cookie{Name: "_id", Value: adminUser}).SetBody(payload).Post(loginUrl)
	if err != nil {
		return nil, err
	}
	token := string(resp.Body())
	return &FileBrowserClient{url: url, adminToken: token, adminUser: adminUser, client: clt}, nil
}

// AddUser adds a user
func (fb *FileBrowserClient) AddUser(username, passwd, scope string) error {
	users, err := fb.AllUsers()
	if err != nil {
		return fmt.Errorf("list users error: %+v", err)
	}
	if conv.StrSliceContains(users, username) {
		return fmt.Errorf("user %s already exists", username)
	}
	payload := map[string]interface{}{"what": "user", "which": nil, "data": map[string]interface{}{"scope": scope, "locale": "zh-cn", "viewMode": "mosaic", "singleClick": false,
		"sorting": map[string]interface{}{"by": "", "asc": false}, "perm": map[string]bool{"admin": false, "execute": true, "create": true, "rename": true, "modify": true, "delete": true, "share": true, "download": true}, "commands": nil, "hideDotfiles": false, "dateFormat": false, "username": username, "rules": nil, "lockPassword": false, "id": 0, "password": passwd}}
	url := fmt.Sprintf("%s/api/users", fb.url)
	resp, err := fb.client.R().SetHeader("x-auth", fb.adminToken).SetBody(payload).Post(url)
	if err != nil {
		fmt.Printf("%v\n", resp)
		fmt.Printf("%v\n", resp.Body())
		return err
	}
	return nil
}

// AllUsers returns all users
func (fb *FileBrowserClient) AllUsers() ([]string, error) {
	var users []string
	var fbUsers []*FbUser
	url := fmt.Sprintf("%s/api/users", fb.url)
	resp, err := fb.client.R().SetHeader("x-auth", fb.adminToken).SetResult(&fbUsers).Get(url)
	if err != nil {
		fmt.Printf("%+v\n", resp)
		fmt.Printf("%s\n", resp.Body())
		return nil, err
	}
	for _, u := range fbUsers {
		users = append(users, u.Username)
	}
	return users, nil
}

// FbUser is a user in filebrowser
type FbUser struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Scope        string `json:"scope"`
	Locale       string `json:"locale"`
	LockPassword bool   `json:"lockPassword"`
	ViewMode     string `json:"viewMode"`
	SingleClick  bool   `json:"singleClick"`
	Perm         struct {
		Admin    bool `json:"admin"`
		Execute  bool `json:"execute"`
		Create   bool `json:"create"`
		Rename   bool `json:"rename"`
		Modify   bool `json:"modify"`
		Delete   bool `json:"delete"`
		Share    bool `json:"share"`
		Download bool `json:"download"`
	} `json:"perm"`
	Commands []interface{} `json:"commands"`
	Sorting  struct {
		By  string `json:"by"`
		Asc bool   `json:"asc"`
	} `json:"sorting"`
	Rules        []interface{} `json:"rules"`
	HideDotfiles bool          `json:"hideDotfiles"`
	DateFormat   bool          `json:"dateFormat"`
}
