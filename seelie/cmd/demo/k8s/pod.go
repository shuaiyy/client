package k8s

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"seelie/cmd/demo/job"
)

// GetPodResourceGPU ...
func GetPodResourceGPU() {
	config := "/Users/shuai.yang/.arena/prod-config"
	cfg, clt, _ := job.NewClientSet(config)
	_ = cfg
	pod, err := clt.CoreV1().Pods("default").Get(context.Background(), "train-114-chief-0", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	container := pod.Spec.Containers[0]
	// gpu
	if val, ok := container.Resources.Limits["nvidia.com/gpu"]; ok {
		gpu := float32(val.Value())
		fmt.Printf("gpu: %f\n", gpu)
		fmt.Printf("resource: %v", val)
	}
}
