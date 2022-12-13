package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"seelie/internal/cli/run"
	"seelie/internal/cli/utils/pretty"
)

// NewDataCmd ...
func NewDataCmd() *cobra.Command {
	return dataCmd
}

func init() {
	// job get
	dataCmd.AddCommand(dataRegisterUserCmd)
}

// dataCmd represents the job command
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "nas storage provided by seelie platform",
	Long: `seelie data --help 

自助注册所有集群的nas storage用户`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage:\n  seelie data {subCommand} [args...]")
		_ = cmd.Help()
	},
}

// ===========================   dataRegisterUserCmd   ===============================
// dataRegisterUserCmd represents the data register user command
var dataRegisterUserCmd = &cobra.Command{
	Use:   "register-user",
	Short: "自助注册Seelie NAS文件浏览系统账号",
	Long: fmt.Sprintf(`Seelie机器学习平台支持多计算集群管理, NAS文件存储服务相互独立, 且%s跨集群访问。
由于网络隔离的原因，NAS文件浏览器服务的域名%s在办公网访问，因此注册账号也需要用户在办公网执行Seelie %s：
 %s
`, pretty.Red("无法"), pretty.Red("只能"), pretty.Green("client命令"), pretty.Blue("seelie data register-user")),
	Run: func(cmd *cobra.Command, args []string) {
		run.RegisterFBUser()
	},
}
