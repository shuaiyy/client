package job

//import (
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/client-go/tools/clientcmd"
//
//	"seelie/internal/server/service/jobcontroller/ml"
//	"seelie/pkg/k8sx/cluster"
//)
//
//// NewCluster for k8s
//func NewCluster(kubeConfig string) cluster.Cluster {
//	if kubeConfig == "" {
//		kubeConfig = clientcmd.RecommendedHomeFile
//	}
//	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
//	if err != nil {
//		panic(err)
//	}
//	clt, err := cluster.New("local", config, runtime.NewScheme())
//	if err != nil {
//		panic(err)
//	}
//	if err = ml.RegisterK8sScheme(clt.GetScheme()); err != nil {
//		panic(err)
//	}
//	return clt
//}
