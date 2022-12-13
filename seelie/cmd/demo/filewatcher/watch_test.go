package filewatcher

import "testing"

func TestWatch(t *testing.T) {
	type args struct {
		configfile string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				configfile: "/Users/shuai.yang/tmp/config.link",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Watch(tt.args.configfile)
		})
	}
}

/*
watch link , 会监听真实的地址
received event: "/Users/shuai.yang/tmp/dev2.yaml": CHMOD
write: "/Users/shuai.yang/tmp/dev2.yaml": WRITE
received event: "/Users/shuai.yang/tmp/dev2.yaml": WRITE
received event: "/Users/shuai.yang/tmp/dev2.yaml": CHMOD
received event: "/Users/shuai.yang/tmp/dev2.yaml": RENAME
# rename之后，watcher失效
*/
