package schema

import (
	"time"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"
)

// User 用户对象
type User struct {
	ID            uint32        `json:"id,string"`                             // 唯一标识
	UserUID       string        `json:"user_uid" binding:"required"`           // 用户名
	RealName      string        `json:"real_name"`                             // 真实姓名
	Phone         string        `json:"phone"`                                 // 手机号
	Email         string        `json:"email" binding:"required"`              // 邮箱
	Department    string        `json:"department"`                            // 部门
	Status        int           `json:"status" binding:"required,max=2,min=1"` // 用户状态(1:启用 2:停用)
	Token         string        `json:"token"`                                 // 用户令牌
	SSHPrivateKey string        `json:"ssh_private_key"`                       // ssh登陆，私钥
	SSHPublicKey  string        `json:"ssh_public_key"`                        // ssh登陆，公钥
	Permissions   []*Permission `json:"permissions"`                           // 用户权限
	CreatedAt     time.Time     `json:"created_at"`                            // 创建时间
	ModifiedAt    time.Time     `json:"modified_at"`                           // 修改时间
}

func (a *User) String() string {
	return json.MarshalToString(a)
}
