package third_party

import (
	"strings"
	"testing"
)

func TestLoginMetabase(t *testing.T) {
	type args struct {
		host     string
		url      string
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "metabase",
			args: args{
				host:     "metabase-ml.ssr.mihoyo.com",
				url:      "https://47.102.127.184:443/api/session",
				username: "shuai.yang@mihoyo.com",
				password: "bb469afc",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := LoginMetabase(tt.args.host, tt.args.url, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginMetabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for k, v := range got {
				t.Log(k, v)
			}
			t.Log("result", string(got1))
			if !strings.Contains(string(got1), "id:") {
				t.Error("login result:", string(got1))
			}

		})
	}
}
