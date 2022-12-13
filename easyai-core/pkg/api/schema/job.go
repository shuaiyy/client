package schema

import (
	"fmt"
	"sort"
	"time"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/json"
)

// Job defines the job
type Job struct {
	CommonConfig `json:",inline" yaml:",inline"`

	ID              uint32       `json:"id" yaml:"id"`
	UkUUID          string       `json:"uk_uuid" yaml:"uk_uuid"`
	State           JobState     `json:"state" yaml:"state"`
	Status          JobStatus    `json:"status" yaml:"status"`
	Message         string       `json:"message" yaml:"message"`
	Reason          string       `json:"reason" yaml:"reason"`
	Framework       JobFramework `json:"framework" yaml:"framework" binding:"required" example:"tensorflow"`
	FrameworkConfig string       `json:"framework_config" yaml:"framework_config"`
	Owner           string       `json:"owner" yaml:"owner"`
	Result          string       `json:"result" yaml:"result"`
	IsDeleted       DeleteFlag   `json:"is_deleted" yaml:"is_deleted"`
	Duration        int32        `json:"duration" yaml:"duration"`
	StartTime       time.Time    `json:"start_time" yaml:"start_time"`
	SchedulerTime   time.Time    `json:"scheduler_time" yaml:"scheduler_time"`
	CreatedAt       time.Time    `json:"created_at" yaml:"created_at"`
	ModifiedAt      time.Time    `json:"modified_at" yaml:"modified_at"`

	Tasks         []*Task `json:"tasks" yaml:"tasks"`
	FrontendField `json:",inline"`
}

// CommonConfig ...
type CommonConfig struct {
	Name            string            `json:"name" yaml:"name" binding:"required" example:"a demo 作业"`
	Namespace       string            `json:"namespace" yaml:"namespace" example:"default"`
	Cluster         string            `json:"cluster" yaml:"cluster" example:"dev"`
	Description     string            `json:"description" yaml:"description" example:"desc作业描述"`
	NodeSelectors   map[string]string `json:"node_selectors" yaml:"node_selectors" example:"key:value"`
	Tolerations     []string          `json:"tolerations" yaml:"tolerations" example:"key=value"`
	Image           string            `json:"image" yaml:"image" binding:"required" example:"registry.cn-shanghai.aliyuncs.com/shuaiyy/mihoyo-ai:tf2.4.3-gpu-jupyter-lab"`
	ImagePullPolicy string            `json:"image_pull_policy" yaml:"image_pull_policy" example:"IfNotPresent"`
	Envs            map[string]string `json:"envs" yaml:"envs" example:"key1:value1"`
	// resources
	CPU      float32 `json:"cpu" yaml:"cpu" binding:"required" example:"2.5"`
	Memory   float32 `json:"memory" yaml:"memory" binding:"required" example:"8.0"`
	GPU      float32 `json:"gpu" yaml:"gpu" example:"0"`
	Hardware string  `json:"hardware" yaml:"hardware"` // cpu or gpu
	// code logic
	Workspace      string            `json:"workspace" yaml:"workspace" example:"/workspace"`
	EntrypointType string            `json:"entrypoint_type" yaml:"entrypoint_type" binding:"required" example:"bash -c"`
	Entrypoint     string            `json:"entrypoint" yaml:"entrypoint" binding:"required" example:"echo hello world"`
	WorkerCount    int32             `json:"worker_count" yaml:"worker_count" binding:"required" example:"1"`
	MaxRetry       int32             `json:"max_retry" yaml:"max_retry" example:"0"`
	VolumeMounts   map[string]string `json:"volume_mounts" yaml:"volume_mounts"`
	HostpathMounts map[string]string `json:"hostpath_mounts" yaml:"hostpath_mounts"`
	IsNonRoot      bool              `json:"is_non_root" yaml:"is_non_root" example:"false"`

	// labels && annotations
	Labels      map[string]string `json:"labels" yaml:"labels"`
	Annotations map[string]string `json:"annotations" yaml:"annotations"`
	MagicFlags  map[string]string `json:"magic_flags" yaml:"magic_flags"`
	TrainArgs   []string          `json:"train_args" yaml:"train_args" example:"lr=0.1,async"`
}

// TFJobConfig tf job config, framework=="tf"
type TFJobConfig struct {
	RoleNodeSelectors map[string]map[string]string `json:"role_node_selectors"`
	CleanPolicy       string                       `json:"clean_policy"`

	Port int32 `json:"port"`
	// worker config
	WorkerImage  string  `json:"worker_image"`
	WorkerPort   int32   `json:"worker_port"`
	WorkerCPU    float32 `json:"worker_cpu"`
	WorkerMemory float32 `json:"worker_memory"`
	WorkerGPU    float32 `json:"worker_gpu"`
	WorkerCount  int32   `json:"worker_count"`
	// ps config
	PsImage  string  `json:"ps_image"`
	PsPort   int32   `json:"ps_port"`
	PsCPU    float32 `json:"ps_cpu"`
	PsMemory float32 `json:"ps_memory"`
	PsGPU    float32 `json:"ps_gpu"`
	PsCount  int32   `json:"ps_count"`
	// chief config
	UseChief    bool    `json:"use_chief"`
	ChiefImage  string  `json:"chief_image"`
	ChiefPort   int32   `json:"chief_port"`
	ChiefCPU    float32 `json:"chief_cpu"`
	ChiefMemory float32 `json:"chief_memory"`
	ChiefGPU    float32 `json:"chief_gpu"`
	ChiefCount  int32   `json:"chief_count"`
	// evaluator config
	UseEvaluator    bool    `json:"use_evaluator"`
	EvaluatorImage  string  `json:"evaluator_image"`
	EvaluatorPort   int32   `json:"evaluator_port"`
	EvaluatorCPU    float32 `json:"evaluator_cpu"`
	EvaluatorMemory float32 `json:"evaluator_memory"`
	EvaluatorGPU    float32 `json:"evaluator_gpu"`
	// master config todo for now, I don't know which distribute case should use master
	MasterImage  string  `json:"master_image"`
	MasterPort   int32   `json:"master_port"`
	MasterCPU    float32 `json:"master_cpu"`
	MasterMemory float32 `json:"master_memory"`
	MasterGPU    float32 `json:"master_gpu"`
	MasterCount  int32   `json:"master_count"`
}

// PytorchJobConfig pytorch job config, framework=="pytorch"
type PytorchJobConfig struct {
	// todo implement me!
}

// NotebookJobConfig Notebook job config, framework=="pytorch"
type NotebookJobConfig struct {
	// todo implement me!
	JupyterServerID uint32 `json:"jupyter_server_id"`
}

// Task task
type Task struct {
	ID         uint32    `json:"id"`
	JobID      uint32    `json:"job_id"`
	Role       string    `json:"role"`
	Index      int8      `json:"index"`
	Command    string    `json:"command"`
	Args       string    `json:"args"`
	OverSale   int8      `json:"over_sale"`
	CPU        float32   `json:"cpu"`
	Memory     float32   `json:"memory"`
	GPU        float32   `json:"gpu"`
	Status     JobStatus `json:"status"`
	Reason     string    `json:"reason"`
	Message    string    `json:"message"`
	ExitCode   string    `json:"exit_code"`
	PodName    string    `json:"pod_name"`
	PodIP      string    `json:"pod_ip"`
	NodeName   string    `json:"node_name"`
	NodeIP     string    `json:"node_ip"`
	Devices    string    `json:"devices"`
	StartTime  time.Time `json:"start_time"`
	Port       int32     `json:"port"`
	NeedCheck  bool      `json:"need_check"`
	RetryCount int8      `json:"retry_count"`
	MaxRetry   int8      `json:"max_retry"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

// RangeTasks ...
func (j *Job) RangeTasks(f func(t *Task) bool) {
	for _, item := range j.Tasks {
		f(item)
	}
}

// WrapFrontendFiled for fe filed to present, depreciated
func (j *Job) WrapFrontendFiled() *Job {
	f := FrontendField{
		FeTasks:    map[string][]*Task{},
		FeResource: map[string]string{},
	}
	for _, t := range j.Tasks {
		f.FeTasks[t.Role] = append(f.FeTasks[t.Role], t)
	}
	for role, tg := range f.FeTasks {
		f.FeResource[role] = fmt.Sprintf("%d * (CPU:%f Memory:%fGB GPU: %f)", len(tg), tg[0].CPU, tg[0].Memory, tg[0].GPU)
	}
	j.FrontendField = f
	return j
}

// TFJobConfig tf config
func (j *Job) TFJobConfig() (TFJobConfig, error) {
	var config TFJobConfig

	if j.Framework != FrameworkTensorflow {
		return config, fmt.Errorf("get frmaework config for tfjob failed, job framework is %s, != %s", j.Framework, FrameworkTensorflow)
	}

	err := json.Unmarshal([]byte(j.FrameworkConfig), &config)
	return config, err
}

// NotebookJobConfig todo Notebook Job Config
func (j *Job) NotebookJobConfig() (NotebookJobConfig, error) {
	var config NotebookJobConfig

	if j.Framework != FrameworkNotebook {
		return config, fmt.Errorf("get frmaework config for notebook job failed, job framework is %s, != %s", j.Framework, FrameworkNotebook)
	}

	err := json.Unmarshal([]byte(j.FrameworkConfig), &config)
	return config, err
}

// PytorchJobConfig todo Pytorch Job Config
func (j *Job) PytorchJobConfig() (PytorchJobConfig, error) {
	var config PytorchJobConfig

	if j.Framework != FrameworkPytorch {
		return config, fmt.Errorf("get frmaework config for pytorch job failed, job framework is %s, != %s", j.Framework, FrameworkPytorch)
	}

	err := json.Unmarshal([]byte(j.FrameworkConfig), &config)
	return config, err
}

// MagicFlag 魔法参数
type MagicFlag struct {
	Key    string `json:"key"`
	DValue string `json:"default_value"`
	Alias  string `json:"alias"`
	Desc   string `json:"description"`
}

// FrontendField for web view
type FrontendField struct {
	FeTasks    map[string][]*Task `json:"fe_tasks"`
	FeResource map[string]string  `json:"fe_resource"`
}

// SortTasks sort tasks
func (ff FrontendField) SortTasks() {
	for _, tasks := range ff.FeTasks {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Index < tasks[j].Index
		})
	}
}
