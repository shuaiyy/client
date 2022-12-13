package k8s

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"seelie/cmd/demo/job"
)

// UpdateNodeSpec ...
func UpdateNodeSpec() {
	config := "/Users/shuai.yang/.kube/config"
	cfg, clt, _ := job.NewClientSet(config)
	_ = cfg
	nodeList, err := clt.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := range nodeList.Items {
		fmt.Printf("node: %s, taint: %v\n", node.Name, node.Spec.Taints)
		for _, c := range node.Status.Conditions {
			fmt.Printf("condition: %s(%v), status: %s, message: %s, reason: %s\n", c.Type, c.LastTransitionTime, c.Status, c.Message, c.Reason)
		}
	}
	node0 := nodeList.Items[0].DeepCopy()
	node0.Spec.Taints = append(node0.Spec.Taints, v1.Taint{
		Key:    "test",
		Value:  "test",
		Effect: v1.TaintEffectNoExecute,
	})
	node1, err := clt.CoreV1().Nodes().Update(context.Background(), node0, metav1.UpdateOptions{})
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		return
	}
	fmt.Println("success", node1)
}

// ListK8sNode ....
func ListK8sNode(config string) {
	cfg, clt, _ := job.NewClientSet(config)
	_ = cfg
	nodeList, err := clt.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	nocare := map[string]bool{
		"zest_zone":        true,
		"env":              true,
		"ingress-pod":      true,
		"system-component": true,
		"pool":             true,
	}
	fmt.Println(config)
	nodes := map[string]string{}
	taint := map[string]string{}
	for _, node := range nodeList.Items {
		fmt.Printf("node: %s, taint: %v, unscheduler: %v, gpu: %v\n", node.Name, node.Spec.Taints, node.Spec.Unschedulable,
			node.Status.Capacity["nvidia.com/gpu"],
		)
		if len(node.Spec.Taints) <= 0 {
			taint[node.Name] = ""
		}
		care := true
		for _, t := range node.Spec.Taints {
			taint[node.Name] = fmt.Sprintf("%s %s=<%s>%s, ", taint[node.Name], t.Key, string(t.Effect), t.Value)
			if t.Effect == "NoSchedule" && nocare[t.Key] {
				care = false
			}
		}
		if care {
			nodes[node.Name] = taint[node.Name]
		}
	}
	fmt.Println(config)
	fmt.Println("===== node taint =====")
	for n, v := range taint {
		fmt.Println(n, v)
	}
	fmt.Println("===== node to look =====")
	for n, v := range nodes {
		fmt.Println(n, v)
	}
}
