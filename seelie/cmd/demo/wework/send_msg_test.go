package wework

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	type args struct {
		user string
		msg  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sss",
			args: args{
				user: "xiaohan.wang",
				msg:  "你好呀(*´▽｀)ノノ",
			},
			wantErr: false,
		},
		{
			name: "sss1",
			args: args{
				user: "yang.fei",
				msg:  "给费扬大佬递茶☕️☕️☕️(*´▽｀)ノノ",
			},
			wantErr: false,
		},
		{
			name: "sss1",
			args: args{
				user: "shuai.yang",
				msg:  "给费扬大佬递茶☕️☕️☕️(*´▽｀)ノノ",
			},
			wantErr: false,
		},
		{
			name: "sss2",
			args: args{
				user: "dong.jiang",
				msg:  "给姜总递茶☕️☕️☕️(*´▽｀)ノノ",
			},
			wantErr: false,
		},
		{
			name: "sss3",
			args: args{
				user: "tuo.liu",
				msg:  "给刘总递茶☕️☕️☕️(*´▽｀)ノノ",
			},
			wantErr: false,
		},
		{
			name: "sss",
			args: args{
				user: "yong.zhang02",
				msg:  "hello world",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Send(tt.args.user, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var (
	_jobHelpDoc = `<font color="info">Job相关操作指令</font>

> 用户必须拥有相应权限(管理员)，才能操作别人的Job
>
> **/jobGet {id}**  查看Job详细信息
> **/jobStop {id}**  停止正在运行的Job
> **/jobList {job_status}** 获取指定状态的Job
> -> **/jobList Running** 获取正在运行的Job
> -> **/jobList Failed** 获取失败的Job
`
	_jupyterHelpDoc = `<font color="info">JupyterServer相关操作指令</font>

> 用户必须拥有相应权限(管理员)，才能操作别人的Job
>
> **/jupyterGet {id}**  查看JupyterServer详细信息
> **/jupyterStop {id}**  停止正在运行的JupyterServer
> **/jupyterList {status}** 获取指定状态的JupyterServer
> -> **/jupyterList Running** 获取正在运行的JupyterServer
> **/jupyterRenewal {id} {required_hours}** 延长JupyterServer的运行时间
`

	_helpDoc = fmt.Sprintf(`<font color="info">Seelie ChatOps机器人使用帮助</font>

> 通过向聊天机器人发送指令，实现平台常用操作

+ 使用场景：
  + 在非办公网/无VPN的情况下，无法访问seelie平台
  + 电脑不在身边或未开机
+ 支持的功能（新需求可以@帅羊羊）
  + Job相关操作
  + JupyterServer相关操作

%s

%s`, _jobHelpDoc, _jupyterHelpDoc)
)

func TestSendMarkdown(t *testing.T) {
	type args struct {
		user string
		msg  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				user: "shuai.yang",
				msg:  _helpDoc,
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				user: "shuai.yang",
				msg:  _jupyterHelpDoc,
			},
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				user: "shuai.yang",
				msg:  _jobHelpDoc,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMarkdown(tt.args.user, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendMarkdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
