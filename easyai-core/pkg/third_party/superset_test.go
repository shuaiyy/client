package third_party

import (
	"strings"
	"testing"
)

func TestLoginSuperset(t *testing.T) {
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
			name: "superset",
			args: args{
				host:     "47.102.127.184:80",
				url:      "https://superset-ml.ssr.mihoyo.com/login",
				username: "shuai.yang@mihoyo.com",
				password: "bb469afc",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := LoginSuperset(tt.args.host, tt.args.url, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginMetabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for k, v := range got {
				t.Log(k, v)
			}
			t.Log("result", string(got1))
			if cookie := got.Get("Set-Cookie"); !strings.Contains(cookie, "session") {
				t.Error("no valid set-cookie result:", cookie)
			}

		})
	}
}
