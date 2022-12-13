package k8s

import "testing"

func TestLabelValue(t *testing.T) {
	tests := []struct {
		name string
		v    string
	}{
		{
			name: "1",
			v:    ".aa+",
		},
		{
			name: "2",
			v:    "a=b",
		},
		{
			name: "2",
			v:    "a.com/b123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LabelValue(tt.v); err != nil {
				t.Logf("error: %v", err)
			}
		})
	}
}

func TestValidLabels(t *testing.T) {
	tests := []struct {
		name   string
		labels map[string]string
	}{
		{
			name: "aa",
			labels: map[string]string{
				"mihoyo.com/123-abc":  "aaa",
				"mih-oyo.com/123-abc": "aaa",
				"mih-oyo.com/123.abc": "aaa",
			},
		},
		{
			name: "bb",
			labels: map[string]string{
				"mihoyo.com/123*abc":  "aaa",
				"mih-oyo.com/123/abc": "aaa",
				"mih-oyo.com/123-abc": "aaa+bbb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidLabels(tt.labels); err != nil {
				t.Logf("ValidLabels() error: %v", err)
			}
		})
	}
}
