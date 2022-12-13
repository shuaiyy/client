package third_party

import (
	"testing"
)

func TestNewFileBrowserClient(t *testing.T) {
	type args struct {
		url       string
		adminUser string
	}
	tests := []struct {
		name string
		args args
		want *FileBrowserClient
	}{
		{
			name: "xx",
			args: args{
				url:       "https://ml-dev.ssr.mihoyo.com/nfsdata",
				adminUser: "shuai.yang",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clt, err := NewFileBrowserClient(tt.args.url, tt.args.adminUser)
			if err != nil {
				t.Errorf("NewFileBrowserClient() = %v, err %v", clt, err)
			}
			users, err := clt.AllUsers()
			if err != nil {
				t.Errorf("AllUsers err %v", err)
			}
			t.Logf("users %+v", users)
			err = clt.AddUser("shuai.yang", "ss11", "ss11")
			if err != nil {
				t.Logf("AddUser err %v", err)
			}
			err = clt.AddUser("ss11", "ss11", "ss11")
			if err != nil {
				t.Errorf("AddUser err %v", err)
			}
		})
	}
}
