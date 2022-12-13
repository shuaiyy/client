package schema

import "time"

// JupyterServer defines the jupyter server
type JupyterServer struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`

	// 集群 资源相关
	Cluster   string  `json:"cluster"`
	Namespace string  `json:"namespace"`
	CPU       float32 `json:"cpu"`
	GPU       float32 `json:"gpu"`
	Memory    float32 `json:"memory"`
	// 任务逻辑
	Image         string            `json:"image"`
	ImageType     JsImageType       `json:"image_type"`
	HostIpc       bool              `json:"host_ipc"`
	Envs          map[string]string `json:"envs"`
	RequiredHours uint8             `json:"required_hours"`
	MagicFlags    map[string]string `json:"magic_flags"`
	// 状态信息
	Status      JsStatus     `json:"status"`
	StartAt     time.Time    `json:"start_at"`
	Deadline    time.Time    `json:"deadline"`
	JobID       uint32       `json:"job_id"`
	PodName     string       `json:"pod_name"`
	PodIP       string       `json:"pod_ip"`
	NodeName    string       `json:"node_name"`
	NodeIP      string       `json:"node_ip"`
	LastRunInfo *LastRunInfo `json:"last_run_info"`

	IsDeleted  DeleteFlag `json:"is_deleted"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt time.Time  `json:"modified_at"`
}

// LastRunInfo ...
type LastRunInfo struct {
	StartAt   time.Time `json:"start_at"`
	Deadline  time.Time `json:"deadline"`
	JobID     uint32    `json:"job_id"`
	PodName   string    `json:"pod_name"`
	PodIP     string    `json:"pod_ip"`
	NodeName  string    `json:"node_name"`
	NodeIP    string    `json:"node_ip"`
	StoppedAt time.Time `json:"stopped_at"`
}
