package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/go-sdk"

	"seelie/internal/cli/run"
	"seelie/internal/cli/utils/stdlogger"
)

// NewJobCmd ...
func NewJobCmd() *cobra.Command {
	return jobCmd
}

func init() {
	// job get
	jobCmd.AddCommand(jobGetCmd)
	jobGetCmd.Flags().Uint32VarP(&jobGetArgs.Jid, "job_id", "j", 0, "job id")
	// job stop
	jobCmd.AddCommand(jobStopCmd)
	jobStopCmd.Flags().Uint32VarP(&jobStopArgs.Jid, "job_id", "j", 0, "job id")
	// job delete
	jobCmd.AddCommand(jobDeleteCmd)
	jobDeleteCmd.Flags().Uint32VarP(&jobDeleteArgs.Jid, "job_id", "j", 0, "job id")
	// job list
	jobCmd.AddCommand(jobListCmd)
	jobListCmd.Flags().Uint32VarP(&jobListArgs.Jid, "job_id", "j", 0, "job id")
	jobListCmd.Flags().StringVarP(&jobListArgs.Owner, "owner", "o", "", "owner")
	jobListCmd.Flags().StringVarP(&jobListArgs.Cluster, "cluster", "c", "", "cluster")
	jobListCmd.Flags().StringVarP(&jobListArgs.Namespace, "namespace", "n", "", "namespace")
	jobListCmd.Flags().StringVarP(&jobListArgs.Name, "name", "N", "", "name中包含关键字")
	jobListCmd.Flags().StringVarP(&jobListArgs.Description, "description", "d", "", "description中包含关键字")
	jobListCmd.Flags().StringVarP(&jobListArgs.EntrypointType, "entrypoint-type", "e", "", "entrypoint-type")
	jobListCmd.Flags().StringVarP(&jobListArgs.Entrypoint, "entrypoint", "E", "", "entrypoint中包含关键字")
	jobListCmd.Flags().StringVarP(&jobListArgs.Status, "status", "s", "", "job状态")
	jobListCmd.Flags().StringVarP(&jobListArgs.FrameWork, "framework", "f", "", "framework类型")
}

var (
	jobGetArgs    jobGetArg
	jobStopArgs   jobGetArg
	jobDeleteArgs jobGetArg
	jobListArgs   jobListArg
)

type jobGetArg struct {
	Jid uint32
}

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "ml training job",
	Long: `seelie job --help 

查询、停止、删除 seelie job`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage:\n  seelie job {subCommand} [args...]")
		_ = cmd.Help()
	},
}

// ===========================   jobGetCmd   ===============================
// jobGetCmd represents the get command
var jobGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get job detail by job id",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := run.GetJobByID(jobGetArgs.Jid); err != nil {
			stdlogger.Error("fail to get job: %+v", err)
		}
	},
}

// ===========================   jobListCmd   ===============================

// jobListArg 查询job
type jobListArg struct {
	Jid            uint32
	Owner          string
	Cluster        string
	Namespace      string
	Name           string
	Description    string
	Entrypoint     string
	EntrypointType string
	Status         string
	FrameWork      string
}

func (j *jobListArg) sdkPayload() *sdk.JobListInput {

	res := sdk.NewJobListInput().WithJobID(j.Jid).WithOwner(j.Owner).WithNamespace(j.Namespace).WithCluster(j.Cluster).
		WithName(j.Name).WithDescription(j.Description).WithEntrypoint(j.Entrypoint).WithEntrypointType(j.EntrypointType).
		WithStatus(j.Status).WithFramework(j.FrameWork)
	res.WithLimit(5)
	return res
}

// jobGetCmd represents the list command
var jobListCmd = &cobra.Command{
	Use:   "list",
	Short: "list jobs",
	Long:  `列出所有的job（有权限查看的）`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := run.ListJobs(jobListArgs.sdkPayload()); err != nil {
			stdlogger.Error("fail to stop job: %+v", err)
		}
	},
}

// ===========================   jobStopCmd   ===============================
// jobStopCmd represents the stop command
var jobStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop a job",
	Long:  `停止正在运行的job`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := run.StopJobByID(jobStopArgs.Jid); err != nil {
			stdlogger.Error("fail to stop job: %+v", err)
		}
	},
}

// ===========================   jobDeleteCmd   ===============================
// jobDeleteCmd represents the delete command
var jobDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a job",
	Long:  `删除job，正在运行中的job无法删除， 需要先stop`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := run.DeleteJobByID(jobDeleteArgs.Jid); err != nil {
			stdlogger.Error("fail to delete job: %+v", err)
		}
	},
}
