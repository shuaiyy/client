package conv

import "testing"

func TestFormatFloat(t *testing.T) {
	type args struct {
		v    float32
		prec int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				v:    0.00001,
				prec: 2,
			},
			want: "0",
		},
		{
			name: "",
			args: args{
				v:    110.00001,
				prec: 2,
			},
			want: "110",
		},
		{
			name: "",
			args: args{
				v:    110.10001,
				prec: 2,
			},
			want: "110.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatFloat(tt.args.v, tt.args.prec); got != tt.want {
				t.Errorf("FormatFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
