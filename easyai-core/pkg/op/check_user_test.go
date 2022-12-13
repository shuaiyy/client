package op

import (
	"testing"
)

func TestGetOpUser(t *testing.T) {
	type args struct {
		username string
		token    string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "shuai.yang",
			args: args{
				username: "shuai.yang",
				token:    "lfd7nN6l74iVI0FdKdkV",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOpUser(tt.args.username, tt.args.token)
			if err != nil {
				t.Errorf("GetOpUser() error = %v, ", err)
				return
			}
			if got.Email == "" {
				t.Errorf("GetOpUser() got = %v, email is empty", got)
			}
		})
	}
}
