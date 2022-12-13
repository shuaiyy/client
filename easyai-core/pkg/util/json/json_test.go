package json

import "testing"

func TestMarshalToString(t *testing.T) {
	type args struct {
		s string
		V map[string]string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{s: "null"}},
		{name: "test2", args: args{s: "{}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Unmarshal([]byte(tt.args.s), &tt.args.V)
			t.Logf("%v", err)
			t.Logf("%+v", tt.args)
			t.Logf("V==nil %+v", tt.args.V == nil)
		})
	}
}
