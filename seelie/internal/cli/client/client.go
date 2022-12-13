package client

import (
	"github.com/spf13/viper"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/go-sdk"
)

// Client for seelie service
type Client struct{ *sdk.Client }

// NewClient ...
func NewClient() (*Client, error) {

	/*
		viper.BindPFlag("host", initCmd.Flags().Lookup("host"))
			viper.BindPFlag("port", initCmd.Flags().Lookup("port"))
			viper.BindPFlag("protocol", initCmd.Flags().Lookup("protocol"))
			viper.BindPFlag("user-token", initCmd.Flags().Lookup("user-token"))
	*/
	var (
		protocol        = viper.GetString("protocol")
		host            = viper.GetString("host")
		port      int64 = viper.GetInt64("port")
		apiPrefix       = "/api/v1"
		token           = viper.GetString("user-token")
		debugHTTP       = viper.GetBool("debug")
	)

	clt, err := sdk.NewClient(host, port, protocol, apiPrefix, token, debugHTTP)
	if err != nil {
		return nil, err
	}
	return &Client{Client: clt}, nil
}
