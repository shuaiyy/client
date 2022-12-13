package command

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"seelie/internal/cli/utils/stdlogger"
)

// NewConfigCmd ...
func NewConfigCmd() *cobra.Command {
	return configCmd
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "get/init/update seelie CLI config file",
	Long: `初始化配置文件：seelie config init -H 127.0.0.1 -P 10088 -p http -t my_token -d false
更新配置文件：seelie config init -H 127.0.0.1 -P 10088 -p http -t my_token -d false --overwrite true
查看当前配置文件内容：seelie config cat`,
	// Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("config called")
	//},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configCatCmd)

	// flags for configInitCmd
	configInitCmd.Flags().StringP("host", "H", "127.0.0.1", "the seelie platform server host")
	configInitCmd.Flags().StringP("protocol", "p", "http", "the seelie platform server protocal")
	configInitCmd.Flags().IntP("port", "P", 80, "the seelie platform server port")
	configInitCmd.Flags().StringP("user-token", "t", "", "the seelie platform user token")
	configInitCmd.Flags().BoolP("debug", "d", false, "client enable http debug")
	configInitCmd.Flags().BoolP("overwrite", "o", false, "overwrite config file if exist")

	_ = viper.BindPFlag("host", configInitCmd.Flags().Lookup("host"))
	_ = viper.BindPFlag("port", configInitCmd.Flags().Lookup("port"))
	_ = viper.BindPFlag("protocol", configInitCmd.Flags().Lookup("protocol"))
	_ = viper.BindPFlag("user-token", configInitCmd.Flags().Lookup("user-token"))
	_ = viper.BindPFlag("debug", configInitCmd.Flags().Lookup("debug"))

}

// configInitCmd represents the init command
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "create/update seelie client config",
	Long: `for example:
seelie config init -H 127.0.0.1 -P 10088 -p http -t my_token -d false -o true

用户token可以从平台主页右上角获取，请保护好个人token，勿随意分享`,
	Run: func(cmd *cobra.Command, args []string) {
		if v, _ := cmd.Flags().GetString("user-token"); v == "" {
			_ = cmd.Help()
			return
		}
		cobra.CheckErr(makeConfigDir())
		if ok, _ := cmd.Flags().GetBool("overwrite"); ok {
			cobra.CheckErr(viper.WriteConfig())
			stdlogger.Info("success to update config\n cat %s", viper.ConfigFileUsed())
			return
		}
		cobra.CheckErr(viper.SafeWriteConfig())
		stdlogger.Info("success to update config, cat %s", viper.ConfigFileUsed())
	},
}

func makeConfigDir() error {
	file := viper.ConfigFileUsed()
	if file == "" {
		usr, err := user.Current()
		if err != nil {
			return err
		}
		file = filepath.Join(usr.HomeDir, ".seelie", "config.yaml")
	}
	if _, err := os.Stat(file); err == nil {
		return nil
	}
	fmt.Println(file)
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); err == nil {
		return nil
	}
	return os.MkdirAll(dir, 0777)
}

// configCatCmd
var configCatCmd = &cobra.Command{
	Use:     "cat",
	Short:   "cat config file",
	Long:    "查看当前使用的配置文件内容",
	Example: "seelie config cat",
	Run: func(cmd *cobra.Command, args []string) {
		file := viper.ConfigFileUsed()
		readFile, err := os.ReadFile(file)
		if err != nil {
			stdlogger.Error("fail to read config file[%s], error: %+v", file, err)
			return
		}
		stdlogger.Info("=== config file content start ===")
		fmt.Println(string(readFile))
		stdlogger.Info("=== config file content end ===")
	},
}
