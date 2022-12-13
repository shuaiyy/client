// Package cli ...
/**
Copyright © 2022 shuai.yang
seelie: the CLI tool for mihoyo machine learning platform Seelie
*/
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/auth"
	"seelie/internal/cli/command"
	"seelie/internal/cli/utils/stdlogger"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "seelie",
	Short: "A CLI tool for seelie(mihoyo machine learning platform)",
	Long: `seelie is a CLI client for Seelie platform:

you can use seelie:
1. submit training jobs, and get/list/stop/delete jobs.
2. login a remote container of a job.
3. download files produced by a job.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	// 全局flag
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.seelie/config.yaml)")
	// 添加子命令
	rootCmd.AddCommand(command.NewSubmitCmd(), command.NewJobCmd(), command.NewConfigCmd(), command.NewUpdateCmd(), command.NewDataCmd())
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		token := viper.GetString("user-token")
		uid, err := parseUserID(token)
		if err != nil {
			fmt.Println("parse user token failed, please check token", uid)
			return err
		}
		viper.Set("user-uid", uid)
		return nil
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else if v := os.Getenv("SeelieConfig"); v != "" {
		viper.SetConfigFile(v)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".seelie"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	if viper.GetBool("debug") {
		stdlogger.LogLevel = stdlogger.DebugLevel
	} else {
		stdlogger.LogLevel = stdlogger.InfoLevel
	}
}

func parseUserID(token string) (string, error) {
	/* configs online
	[jwt_auth]
	# 是否启用
	enable = true
	# 签名方式(支持：HS512/HS384/HS512)
	signing_method = "HS512"
	# 签名key
	signing_key = "seelie"
	# 过期时间（单位秒） 2年
	expired = 126144000
	# 存储(支持：file/redis)
	store = "file"
	# 文件路径
	file_path = "data/jwt_auth.db"
	# redis 数据库(如果存储方式是redis，则指定存储的数据库)
	redis_db = 10
	# 存储到 redis 数据库中的键名前缀
	redis_prefix = "auth_"
	*/
	signingKey := "seelie"
	var opts []auth.Option
	opts = append(opts, auth.SetExpired(3600*24*365*2))
	opts = append(opts, auth.SetSigningKey([]byte(signingKey)))
	opts = append(opts, auth.SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(signingKey), nil
	}))
	var method jwt.SigningMethod = jwt.SigningMethodHS512
	opts = append(opts, auth.SetSigningMethod(method))

	auther := auth.New(nil, opts...)
	uid, _, _, _, err := auther.ParseUserInfo(context.Background(), token, true)
	return uid, err
}
