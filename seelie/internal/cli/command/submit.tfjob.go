package command

import (
	"fmt"

	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"

	"seelie/internal/cli/run"
	"seelie/internal/cli/utils/stdlogger"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/api/schema"
)

// submitTFJobCmd represents submit tfjob command todo 提供栗子 && 文案优化
var submitTFJobCmd = &cobra.Command{
	Use:     "tfjob",
	Aliases: []string{"tfJob", "tf", "tfjob"},
	Short:   "submit a tensorflow job",
	Long: ` 提交 seelie tensorflow job
seelie submit tfjob --help # todo 提供栗子 && 文案优化`,
	Example: `# 1. 单机job
seelie submit tfjob --name "test job" --description "submit by seelie cli" \
 --cluster dev --namespace default --cpu 2 --memory 8 --gpu 0 --worker-count 1 \
 --image registry.cn-shanghai.aliyuncs.com/shuaiyy/mihoyo-ai:tf2.4.3-gpu-jupyter-lab \
 --entrypoint-type "bash -c" --entrypoint "sleep 10m; echo failed; exit 1" \
 -E enable_ema=1

# 2. 分布式训练 2 PS + 4 Worker
seelie submit tfjob --name "dist-tf-2ps-4worker" --description "submit by seelie cli" \
 --cluster dev --namespace default --cpu 6 --memory 12 --gpu 0 --worker-count 4 \
 --image registry.cn-shanghai.aliyuncs.com/shuaiyy/2233:tf1.5-dist-mnist-demo-train_op1.4 \
 --entrypoint-type "python" --entrypoint "/workspace/dist_mnist.py" \
 -E enable_ema=1 -E env_aaa=test \
 -T data_dir=/data_dir/mnist_data -T train_steps=30 -T batch_size=32 \
 -M m_over_sale=3 -M m_enable_debug_toolbox=true -M m_instance_retain=true -M m_retain_time=1h \
 --ps-count 2 --ps-cpu 4 --ps-memory 20`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := submitTfJobArgs.Validator(); err != nil {
			stdlogger.Error("%v", err)
			_ = cmd.Help()
			return
		}

		if err := run.SubmitTFJob(submitTfJobArgs.Common, submitTfJobArgs.TFJob); err != nil {
			stdlogger.Error("fail to submit job: %+v", err)
			stdlogger.Info("for help: ")
			fmt.Println("  seelie submit tfjob --help")

		}
	},
}

var submitTfJobArgs submitTfJobArg

type submitTfJobArg struct {
	Common schema.CommonConfig
	TFJob  schema.TFJobConfig
}

func (stf *submitTfJobArg) Validator() error {
	c := stf.Common
	if c.Name == "" || c.Cluster == "" || c.Namespace == "" || c.CPU <= 0 || c.Memory <= 0 || c.WorkerCount <= 0 ||
		c.EntrypointType == "" {
		return fmt.Errorf("job sumit args非法，缺少必要参数或必要参数的值非法")
	}
	return nil
}

func bindTfJobArgs(cmd *cobra.Command, tf *schema.TFJobConfig, fss *cliflag.NamedFlagSets) {
	// 注意 命名和短命名不要和common绑定冲突
	// worker
	workerFS := fss.FlagSet("tf分布式: worker")
	workerFS.Int32VarP(&tf.Port, "port", "", 0, optionArg+"tf服务监听port, 默认值2222")
	workerFS.StringVarP(&tf.WorkerImage, "worker-image", "", "", optionArg+"指定worker的镜像，默认值使用`image`参数")
	workerFS.Int32VarP(&tf.WorkerPort, "worker-port", "", 0, optionArg+"指定worker的服务端口, 默认值使用`port`参数")
	workerFS.Float32VarP(&tf.WorkerCPU, "worker-cpu", "", 0, optionArg+"指定worker的cpu最大值, 默认值使用`cpu`参数")
	workerFS.Float32VarP(&tf.WorkerMemory, "worker-memory", "", 0, optionArg+"指定worker的memory最大值, 默认值使用`memory`参数")
	workerFS.Float32VarP(&tf.WorkerGPU, "worker-gpu", "", 0, optionArg+"指定worker的gpu, 默认值使用`gpu`参数")
	workerFS.Int32VarP(&tf.WorkerCount, "worker-num", "", 0, optionArg+"指定worker的数量, 默认值使用`worker-count`参数")
	// ps
	psFS := fss.FlagSet("tf分布式: ps")
	psFS.StringVarP(&tf.PsImage, "ps-image", "", "", optionArg+"指定ps的镜像, 默认值使用`image`参数")
	psFS.Int32VarP(&tf.PsPort, "ps-port", "", 0, optionArg+"指定ps的服务端口, 默认值使用`port`参数")
	psFS.Float32VarP(&tf.PsCPU, "ps-cpu", "", 0, optionArg+"指定ps的cpu最大值, 默认值使用`cpu`参数")
	psFS.Float32VarP(&tf.PsMemory, "ps-memory", "", 0, optionArg+"指定ps的memory最大值, 默认值使用`memory`参数")
	psFS.Float32VarP(&tf.PsGPU, "ps-gpu", "", 0, optionArg+"指定ps的gpu, 默认值为0，通常PS实例不需要gpu")
	psFS.Int32VarP(&tf.PsCount, "ps-count", "", 0, optionArg+"指定ps的数量,值为0时，表示非ps-worker类型的分布式任务")
	// chief
	chiefFS := fss.FlagSet("tf分布式: chief")
	chiefFS.BoolVarP(&tf.UseChief, "use-chief", "", false, optionArg+"启用chief节点， 默认不使用chief")
	chiefFS.StringVarP(&tf.ChiefImage, "chief-image", "", "", optionArg+"指定chief的镜像, 默认值使用`image`参数")
	chiefFS.Int32VarP(&tf.ChiefPort, "chief-port", "", 0, optionArg+"指定chief的服务端口, 默认值使用`port`参数")
	chiefFS.Float32VarP(&tf.ChiefCPU, "chief-cpu", "", 0, optionArg+"指定chief的cpu最大值, 默认值使用`cpu`参数")
	chiefFS.Float32VarP(&tf.ChiefMemory, "chief-memory", "", 0, optionArg+"指定chief的memory最大值, 默认值使用`memory`参数")
	chiefFS.Float32VarP(&tf.ChiefGPU, "chief-gpu", "", 0, optionArg+"指定chief的gpu, 默认值使用`gpu`参数")
	chiefFS.Int32VarP(&tf.ChiefCount, "chief-count", "", 1, optionArg+"指定chief的数量，通常只需要1个")
	// evaluator
	evaluatorFS := fss.FlagSet("tf分布式: evaluator")
	evaluatorFS.BoolVarP(&tf.UseEvaluator, "use-evaluator", "", false, optionArg+"启用evaluator节点， 默认不使用evaluator")
	evaluatorFS.StringVarP(&tf.EvaluatorImage, "evaluator-image", "", "", optionArg+"指定evaluator的镜像, 默认值使用`image`参数")
	evaluatorFS.Int32VarP(&tf.EvaluatorPort, "evaluator-port", "", 0, optionArg+"指定evaluator的服务端口, 默认值使用`port`参数")
	evaluatorFS.Float32VarP(&tf.EvaluatorCPU, "evaluator-cpu", "", 0, optionArg+"指定evaluator的cpu最大值, 默认值使用`cpu`参数")
	evaluatorFS.Float32VarP(&tf.EvaluatorMemory, "evaluator-memory", "", 0, optionArg+"指定evaluator的memory最大值, 默认值使用`memory`参数")
	evaluatorFS.Float32VarP(&tf.EvaluatorGPU, "evaluator-gpu", "", 0, optionArg+"指定evaluator的gpu, 通常不需要gpu")
	// master
	masterFS := fss.FlagSet("tf分布式: master")
	masterFS.StringVarP(&tf.MasterImage, "master-image", "", "", optionArg+"指定master的镜像, 默认值使用`image`参数")
	masterFS.Int32VarP(&tf.MasterPort, "master-port", "", 0, optionArg+"指定master的服务端口, 默认值使用`port`参数")
	masterFS.Float32VarP(&tf.MasterCPU, "master-cpu", "", 0, optionArg+"指定master的cpu最大值, 默认值使用`cpu`参数")
	masterFS.Float32VarP(&tf.MasterMemory, "master-memory", "", 0, optionArg+"指定master的memory最大值, 默认值使用`memory`参数")
	masterFS.Float32VarP(&tf.MasterGPU, "master-gpu", "", 0, optionArg+"指定master的gpu, 默认值使用`gpu`参数")
	masterFS.Int32VarP(&tf.MasterCount, "master-count", "", 0, optionArg+"指定master的数量,值为0时，表示非master-worker类型的分布式任务")
	cmd.Flags().AddFlagSet(workerFS)
	cmd.Flags().AddFlagSet(psFS)
	cmd.Flags().AddFlagSet(chiefFS)
	cmd.Flags().AddFlagSet(evaluatorFS)
	cmd.Flags().AddFlagSet(masterFS)
}

func init() {
	tfJobFss := &cliflag.NamedFlagSets{}
	bindJobCommonArgs(submitTFJobCmd, &submitTfJobArgs.Common, tfJobFss)
	bindTfJobArgs(submitTFJobCmd, &submitTfJobArgs.TFJob, tfJobFss)
	updateCmdPrintFuncWithNamedFlagSet(submitTFJobCmd, tfJobFss)
}
