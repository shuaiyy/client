package schema

import "time"

// Cluster cluster
type Cluster struct {
	ID             uint32       `json:"id"`               // 自增id
	UID            string       `json:"uid"`              // cluster uid
	Name           string       `json:"name"`             // cluster name, name == uid
	Comment        string       `json:"comment"`          // cluster comment
	KubeConfigPath string       `json:"kube_config_path"` // kubeConfigPath
	KubeConfigText string       `json:"kube_config_text"` // kubeConfigText
	Status         int          `json:"status"`           // cluster status, 1: active, 2: disable
	CreatedAt      time.Time    `json:"created_at"`       // create at
	ModifiedAt     time.Time    `json:"modified_at"`      // modified at
	Namespaces     []*Namespace `json:"namespaces"`       // namespaces
}

// Namespace ns
type Namespace struct {
	ID         uint32    `json:"id"`          // 自增id
	UID        string    `json:"uid"`         // namespace uid
	ClusterID  uint32    `json:"cluster_id"`  // cluster id, foreign key
	ClusterUID string    `json:"cluster_uid"` // cluster uid
	Name       string    `json:"name"`        // namespace name, name == uid
	Comment    string    `json:"comment"`     // namespace comment
	Type       string    `json:"type"`        // namespace type (training || serving)
	Status     int       `json:"status"`      // namespace status, 1: active, 2: disable
	CreateAt   time.Time `json:"create_at"`   // create at
	ModifiedAt time.Time `json:"modified_at"` // modified at
}

// ClusterConfig 集群相关配置
type ClusterConfig struct {
	Disabled    bool                       `json:"disabled"`
	Name        string                     `json:"name"`
	FbURL       string                     `json:"fb_url"`    // file browsers' url
	FbPrefix    map[string]string          `json:"fb_prefix"` // prefix for file browser: job/data/jupyter
	ProxyURL    string                     `json:"proxy_url"` // proxy url for job http service
	ProxyPrefix map[string]string          `json:"proxy_prefix"`
	TrainingNS  map[string]NamespaceConfig `json:"training_namespaces"` // ns for train
	ServingNS   map[string]NamespaceConfig `json:"serving_namespaces"`  // ns for serving
}

// NamespaceConfig ...
type NamespaceConfig struct {
	Name            string   `json:"name"`
	Disabled        bool     `json:"disabled"`
	ShareNamespaces []string `json:"share_namespaces"` // 可以共享其他业务空间的资源池
	Type            string   `json:"type"`             // 类型
	Comment         string   `json:"comment"`
	Exclusive       bool     `json:"exclusive"` // 本业务空间的资源池独占，禁止其他业务方分享
}
