package run

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"platgit.mihoyo.com/easyai/easyai-core/pkg/go-sdk"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/third_party"
	"platgit.mihoyo.com/easyai/easyai-core/pkg/util/conv"
	"seelie/internal/cli/client"
	"seelie/internal/cli/utils/pretty"
	"seelie/internal/cli/utils/stdlogger"
)

// RegisterFBUser ...
func RegisterFBUser() {
	seelieClt, err := client.NewClient()
	if err != nil {
		stdlogger.Fatal("fail to init seelie client: %v", err)
	}

	clusterConfigs, err := seelieClt.GetClusterConfigs(sdk.NewClusterConfigsInput())
	if err != nil {
		stdlogger.Fatal("fail to get cluster configs: %v", err)
	}
	uid := viper.GetString("user-uid")
	adminUser := "shuai.yang"
	for clsName, clsCfg := range clusterConfigs {
		for _, fbName := range []string{"data", "jupyter"} {
			if clsCfg.FbPrefix != nil && clsCfg.FbPrefix[fbName] != "" {
				subPath := strings.Split(strings.TrimPrefix(clsCfg.FbPrefix[fbName], "/"), "/")[0]
				fbURL := fmt.Sprintf("%s/%s", strings.TrimSuffix(clsCfg.FbURL, "/"), subPath)
				stdlogger.Info("cluster: %s, nas storage for %s: %s", clsName, fbName, fbURL)
				fbClt, err := third_party.NewFileBrowserClient(fbURL, adminUser)
				if err != nil {
					stdlogger.Fatal("fail to init filebrowser client for nas storage(%s): %v", fbName, err)
				}
				allUsers, err := fbClt.AllUsers()
				if err != nil {
					stdlogger.Warn("fail to list users for nas storage(%s): %v", fbName, err)
					continue
				}
				if conv.StrSliceContains(allUsers, uid) {
					stdlogger.Warn("user<%s> already exist for nas storage(%s)", uid, fbName)
					continue
				}
				if err = fbClt.AddUser(uid, uid, uid); err != nil {
					stdlogger.Error("fail to create user(%s) for nas storage(%s): %v", uid, fbName, err)
					continue
				}
				stdlogger.Info("%s create user(%s) for nas storage(%s)", pretty.Green("success"), uid, fbName)
			}
		}
	}
}
