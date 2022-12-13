package schema

// LogFile ...
// 1. we may output a few other log files except stdout.log & stderr.log
// 2. we may save logs to nas or oss, or just only in pod container.
type LogFile struct {
	Source        LogStorageSource `json:"source" form:"source"` // k8s nas oss
	FileName      LogFileName      `json:"file_name" form:"file_name"`
	Path          string           `json:"path,omitempty" form:"path"`
	Cluster       string           `json:"cluster,omitempty" form:"cluster"`
	Namespace     string           `json:"namespace,omitempty" form:"namespace"`
	Instance      string           `json:"instance,omitempty" form:"instance"`
	MainContainer string           `json:"main_container,omitempty" form:"main_container"`
	JobID         uint32           `json:"job_id,omitempty" form:"job_id"`
}

// LogStorageSource ...
type LogStorageSource string

// for log storage backend
const (
	LogStorageK8s LogStorageSource = "k8s"
	LogStorageNFS LogStorageSource = "nfs"
	LogStorageOSS LogStorageSource = "s3"
)

// LogFileName for file name
type LogFileName string

// for log file name
const (
	LogFileStdout = "stdout.log"
	LogFileStderr = "stderr.log"
	LogFileStdAll = "std_all.log" // k8s's pod/log only support mixed log stream
)
