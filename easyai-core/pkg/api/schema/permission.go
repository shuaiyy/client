package schema

import "time"

// Permission 用户权限
type Permission struct {
	ID           uint32    `json:"id,string"`            // 唯一标识
	UserID       uint32    `json:"user_id,string"`       // 用户id
	UserUID      string    `json:"user_uid,string"`      // 用户UID
	ClusterUID   string    `json:"cluster_uid,string"`   // k8s cluster UID
	NamespaceUID string    `json:"namespace_uid,string"` // k8s namespace UID
	RoleUID      string    `json:"role_uid,string"`      // 角色UID
	Comment      string    `json:"comment"`              // 权限备注说明
	CreatedAt    time.Time `json:"created_at"`           // 创建时间
	ModifiedAt   time.Time `json:"modified_at"`          // 修改时间
}
