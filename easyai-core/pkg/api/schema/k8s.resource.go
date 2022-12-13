package schema

// copied from pkg/k8sx/resourcecache/*

// ResourceSummary is the summary of cluster resource
type ResourceSummary struct {
	Nodes       []*NodeInfo `json:"nodes"`
	Allocatable Resource    `json:"allocatable"`
	Idle        Resource    `json:"idle"`
	Used        Resource    `json:"used"`
}

// NodeInfo is the info of k8s node
type NodeInfo struct {
	Name          string              `json:"name"`
	Hostname      string              `json:"hostname"`
	Ready         bool                `json:"ready"`
	Unschedulable bool                `json:"unschedulable"`
	Allocatable   Resource            `json:"allocatable"`
	Capacity      Resource            `json:"capacity"`
	Limit         Resource            `json:"limit"`
	Request       Resource            `json:"request"`
	Pods          map[string]*PodInfo `json:"pods"`
}

// PodInfo is the info of k8s pod in node
type PodInfo struct {
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	NodeName  string   `json:"nodeName"`
	Limit     Resource `json:"limit"`
	Request   Resource `json:"request"`
	Phase     string   `json:"phase"`
}

// Resource is k8s resource
type Resource struct {
	CPU float32 `json:"cpu"`
	Mem float32 `json:"memory"`
	GPU float32 `json:"gpu"`
}

// SeelieNodeType for seelie node marker
type SeelieNodeType string

// the same const also defined in `internal/server/service/jobcontroller/common/const.go`
const (
	SeelieNodeAll      SeelieNodeType = ""
	SeelieNodeTraining SeelieNodeType = "training"
	SeelieNodeServing  SeelieNodeType = "serving"
)

// WARN: Pool和Namespace的label和污点key是固定的，且二者始终一样！
// the same const also defined in `internal/server/service/jobcontroller/common/const.go`
const (
	SeelieNodeLabelKeyPool      = "seelie-node"
	SeelieTaintKeyPool          = "seelie-node"
	SeelieNodeLabelKeyNamespace = "allow-namespace"
	SeelieTaintKeyNamespace     = "allow-namespace"
	SeelieTaintKeyGPU           = "gpu-pod" // 米哈游运维容器发布平台的约定
)
