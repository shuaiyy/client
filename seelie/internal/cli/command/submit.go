package command

import (
	"fmt"

	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/api/schema"
)

// NewSubmitCmd ...
func NewSubmitCmd() *cobra.Command {
	return submitCmd
}

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "submit a training job",
	Long: `seelie submit --help 

提交 seelie job， 不同框架类型的job请查看子命令，比如：

seelie submit tfJob --help # 提交tensorflow job
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage:\n  seelie submit {subCommand} [args...]")
		_ = cmd.Help()
	},
}

func init() {
	submitCmd.AddCommand(submitTFJobCmd)
}

func bindJobCommonArgs(cmd *cobra.Command, common *schema.CommonConfig, fss *cliflag.NamedFlagSets) {
	// job属性
	jobAttrFS := fss.FlagSet("job属性")
	jobAttrFS.StringVarP(&common.Name, "name", "n", "", mustArg+"job name")
	jobAttrFS.StringVarP(&common.Description, "description", "d", "", optionArg+"job描述")
	// 计算资源
	resourceFS := fss.FlagSet("resource计算资源")
	resourceFS.StringVarP(&common.Cluster, "cluster", "C", "", mustArg+"计算集群")
	resourceFS.StringVarP(&common.Namespace, "namespace", "N", "", mustArg+"业务空间")
	resourceFS.Float32VarP(&common.CPU, "cpu", "c", 0.1, mustArg+"worker实例的cpu最大可用值, 数值0.1表示1/10的cpu时间片")
	resourceFS.Float32VarP(&common.Memory, "memory", "m", 1.0, mustArg+"worker实例的内存最大可用值, GB")
	resourceFS.Float32VarP(&common.GPU, "gpu", "g", 0.0, optionArg+"worker实例的GPU, 单位卡，目前暂未实现虚拟化，不支持0.x粒度的卡申请")
	resourceFS.Int32VarP(&common.WorkerCount, "worker-count", "w", 1, mustArg+"worker实例的数量")
	// 任务运行逻辑
	taskFS := fss.FlagSet("task任务逻辑")
	taskFS.StringVarP(&common.Image, "image", "i", "", mustArg+"job任务镜像")
	taskFS.StringVarP(&common.ImagePullPolicy, "image-pull-policy", "",
		"IfNotPresent", optionArg+"镜像下载策略, IfNotPresent：节点上不存在才下载，Always: 从镜像仓库拉取最新镜像")
	taskFS.BoolVarP(&common.IsNonRoot, "non-root", "", false, optionArg+"使用非root用户运行进程，即使用镜像内预设的linux group:user")
	taskFS.Int32VarP(&common.MaxRetry, "max-retry", "", 0, optionArg+"最大重试次数，fail-over功能暂未实现")
	taskFS.StringVarP(&common.Workspace, "workspace", "", "/workspace", optionArg+"任务进程的工作目录，如无必要请勿更改")
	taskFS.StringVarP(&common.EntrypointType, "entrypoint-type", "e", "bash", mustArg+"任务进程执行方式，如 bash|python|sh 或任意可执行文件路径，一种特殊的情况是: `bash|sh|python -c`")
	taskFS.StringVarP(&common.Entrypoint, "entrypoint", "p", "", mustArg+"被执行的文件路径，当entrypoint-type为 `xxx -c`时，此处应填被执行的脚本内容")
	taskFS.StringToStringVarP(&common.Envs, "env", "E", nil, optionArg+"环境变量, 比如: -E env_k1=v1 -E k2=v2")
	taskFS.StringSliceVarP(&common.TrainArgs, "train-arg", "T", nil, optionArg+"训练超参数，追加到entrypoint后,比如: -T lr=0.01 -T sync")
	// 高级调度特性
	advanceFS := fss.FlagSet("seelie平台高级功能")
	advanceFS.StringToStringVarP(&common.NodeSelectors, "node-selector", "s", nil,
		optionArg+"调度参数(慎重使用),选择特定的节点运行job,多key之间为And关系, 比如: -s k1=v1 -s k2=v2")
	advanceFS.StringSliceVarP(&common.Tolerations, "toleration", "t", nil,
		optionArg+"调度参数(慎重使用),某些节点禁止任何job调度, 除非容忍所有的排异条件，相同key为Or关系，不同key为And关系，比如: -t k1=v1 -t k1=v2 -t k=v")
	advanceFS.StringToStringVarP(&common.MagicFlags, "magic-flag", "M", nil,
		optionArg+"魔法开关，开启seelie高级特性, 比如: -M m_instance_retain=true 为容器失败时保留现场不销毁")
	// add to command
	cmd.Flags().AddFlagSet(jobAttrFS)
	cmd.Flags().AddFlagSet(resourceFS)
	cmd.Flags().AddFlagSet(taskFS)
	cmd.Flags().AddFlagSet(advanceFS)

	//cmd.Flags().StringVarP(&common.Name, "name", "n", "", "job name, 必填")
	//cmd.Flags().StringVarP(&common.Cluster, "cluster", "C", "", "计算集群, 必填")
	//cmd.Flags().StringVarP(&common.Namespace, "namespace", "N", "", "业务空间, 必填")
	//cmd.Flags().StringVarP(&common.Description, "description", "d", "", "job描述, 选填")
	//cmd.Flags().StringToStringVarP(&common.NodeSelectors, "node-selector", "s", nil,
	//	"高级调度参数(非必填，慎重使用),选择特定的节点运行job,多条件为And关系, 比如: -s k1=v1 -s k2=v2")
	//cmd.Flags().StringSliceVarP(&common.Tolerations, "toleration", "t", nil,
	//	"高级调度参数(非必填，慎重使用),某些节点禁止任何job调度, 除非容忍所有的排异条件，相同key为Or关系，不同key为And关系，比如: -t k1=v1 -t k1=v2 -t k=v")
	//cmd.Flags().StringVarP(&common.Image, "image", "i", "", "job任务镜像, 必填")
	//cmd.Flags().StringVarP(&common.ImagePullPolicy, "image-pull-policy", "",
	//	"IfNotPresent", "镜像下载策略, 选填，IfNotPresent：节点上不存在才下载，Always: 从镜像仓库拉取最新镜像")
	//cmd.Flags().Float32VarP(&common.CPU, "cpu", "c", 0.1, "worker实例的cpu最大可用值, 数值0.1表示1/10的cpu时间片")
	//cmd.Flags().Float32VarP(&common.Memory, "memory", "m", 1.0, "worker实例的内存最大可用值, GB")
	//cmd.Flags().Float32VarP(&common.GPU, "gpu", "g", 0.0, "worker实例的GPU, 单位卡，目前暂未实现虚拟化，不支持0.x粒度的卡申请")
	//cmd.Flags().BoolVarP(&common.IsNonRoot, "non-root", "", false, "使用非root用户运行进程，即使用镜像内预设的linux group:user")
	//cmd.Flags().Int32VarP(&common.WorkerCount, "worker-count", "w", 1, "worker实例的数量")
	//cmd.Flags().Int32VarP(&common.MaxRetry, "max-retry", "", 0, "最大重试次数，fail-over功能暂未实现")
	//cmd.Flags().StringVarP(&common.Workspace, "workspace", "", "/workspace", "任务进程的工作目录，如无必要请勿更改")
	//cmd.Flags().StringVarP(&common.EntrypointType, "entrypoint-type", "e", "bash", "任务进程执行方式，如 bash|python|sh 或任意可执行文件路径，一种特殊的情况是: `bash|sh|python -c`")
	//cmd.Flags().StringVarP(&common.Entrypoint, "entrypoint", "p", "", "被执行的文件路径，非必填，当entrypoint-type为 `xxx -c`时，此处应填被执行的脚本内容")
	//cmd.Flags().StringToStringVarP(&common.Envs, "env", "E", nil, "环境变量, 选填， 比如: -E env_k1=v1 -E k2=v2")
	//cmd.Flags().StringSliceVarP(&common.TrainArgs, "train-arg", "T", nil, "训练超参数，追加到entrypoint后, 选填， 比如: -T lr=0.01 -T sync")
	//cmd.Flags().StringToStringVarP(&common.MagicFlags, "magic-flag", "M", nil,
	//	"魔法开关，开启seelie高级特性, 选填， 比如: -M m_instance_retain=true -M k2=v2")
}
