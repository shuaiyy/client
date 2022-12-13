package main

/**
  一些功能验证的demo代码
*/
import (
	"fmt"
	"strings"

	"seelie/cmd/demo/k8s"
)

func main() {
	//gin.Cors()
	//_ = job.CreateTFJobSingle
	//job.CreateTFJobSingle()
	//k8s.UpdateNodeSpec()
	ListNode()
}

// ListNode ...
func ListNode() {

	ss := []string{
		"/Users/shuai.yang/.arena/kubeconfig",
		"/Users/shuai.yang/.arena/aws-dev-config",
		//"/Users/shuai.yang/.arena/prod-config",
		//"/Users/shuai.yang/.arena/aws-prod-config",
	}
	for _, s := range ss {
		k8s.ListK8sNode(s)
	}
}

// CommandArgs ...
func CommandArgs() {
	//strings.Fields 无法准确切分flag，需要自己实现
	ss := []string{
		`python "my aaa.txt" --key="aa cc" --key=""   --key3="xxx"`,
		`python "my aaa.txt" --key="aa cc" --key=" "   --key3="xxx"`,
		`python "my aaa.txt" --key="aa cc" --key3="xxx"`,
		fmt.Sprintf("%s %s", `bash -c`, `"sleep 604800"`),
	}
	for _, s := range ss {
		fmt.Println(strings.Join(strings.Fields(s), "* *"))
	}
}
