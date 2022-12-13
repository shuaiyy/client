package op

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type opResponse struct {
	RetCode int64  `json:"retcode"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

// User for opHome
type User struct {
	Code                int      `json:"code"`
	Username            string   `json:"username"`
	ChineseName         string   `json:"chinese_name"`
	Token               string   `json:"token"`
	MemberOf            []string `json:"member_of"`
	Email               string   `json:"email"`
	EmpType             string   `json:"emp_type"`
	EmpKind             string   `json:"emp_kind"`
	IsOutsourcing       bool     `json:"is_outsourcing"`
	SaveTime            string   `json:"save_time"`
	CsLoc               string   `json:"cs_loc"`
	EnglishAbbreviation string   `json:"english_abbreviation"`
	IsAdvisor           bool     `json:"is_advisor"`
}

// GetOpUser ,,,
func GetOpUser(username, token string) (*User, error) {

	var result opResponse

	url := "https://op-takumi.mihoyo.com/employee/op/checkLogin"
	resp, err := resty.New().NewRequest().SetBody(map[string]string{"username": username, "token": token}).
		SetHeader("Content-Type", "application/json").SetResult(&result).Post(url)
	if err != nil {
		return nil, err
	}
	if result.Data.Email == "" || result.Data.Username == "" {
		return nil, fmt.Errorf("fail to parse user info, retcode: %d, message: %s, raw body: %s",
			result.RetCode, result.Message, resp.Body())
	}
	return &result.Data, nil
}

const (
	cookieID    = "_id"
	cookieToken = "_t"
)

// GetOpUerFromRequest from gin context
func GetOpUerFromRequest(c *gin.Context) (*User, error) {
	uid, err := c.Cookie(cookieID)
	if err != nil {
		return nil, err
	}
	token, err := c.Cookie(cookieToken)
	if err != nil {
		return nil, err
	}
	return GetOpUser(uid, token)
}
