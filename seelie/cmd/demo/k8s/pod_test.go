package k8s

import "testing"

func TestGetPodResourceGPU(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetPodResourceGPU()
		})
	}
}
