package conv

import "testing"

func TestI(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want S
	}{
		{args: args{v: true}, want: S("true")},
		{args: args{v: false}, want: S("false")},
		{args: args{v: 10}, want: S("10")},
		{args: args{v: "123"}, want: S("123")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := I(tt.args.v); got != tt.want {
				t.Errorf("I() = %v, want %v", got, tt.want)
			}
		})
	}
}
