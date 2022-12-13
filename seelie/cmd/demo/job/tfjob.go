package job

//
//import (
//	"context"
//	"fmt"
//
//	commonv1 "github.com/kubeflow/common/pkg/apis/common/v1"
//	v1 "github.com/kubeflow/training-operator/pkg/apis/tensorflow/v1"
//	tfclient "github.com/kubeflow/training-operator/pkg/client/clientset/versioned/typed/tensorflow/v1"
//	corev1 "k8s.io/api/core/v1"
//	"k8s.io/apimachinery/pkg/api/resource"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/client-go/kubernetes"
//	"k8s.io/client-go/kubernetes/scheme"
//	"k8s.io/client-go/rest"
//	"k8s.io/client-go/tools/clientcmd"
//)
//
//func createJob(config *rest.Config, job *v1.TFJob) *v1.TFJob {
//
//	tfjob, err := tfclient.NewForConfigOrDie(config).TFJobs("default").Create(context.Background(), job, metav1.CreateOptions{
//		TypeMeta: metav1.TypeMeta{
//			Kind:       "TFJob",
//			APIVersion: "v1",
//		},
//		DryRun:       nil,
//		FieldManager: "",
//	})
//	if err != nil {
//		fmt.Printf("Failed to create job: %v\n", err)
//		return nil
//	}
//	fmt.Printf("Created job: %v\n", tfjob)
//	return job
//}
//
//// NewClientSet new
//func NewClientSet(config string) (*rest.Config, *kubernetes.Clientset, error) {
//	clientcmdConfig, err := clientcmd.BuildConfigFromFlags("", config)
//	if err != nil {
//		panic(err)
//
//	}
//
//	clientset, err := kubernetes.NewForConfig(clientcmdConfig)
//	if err != nil {
//		panic(err)
//	}
//	return clientcmdConfig, clientset, nil
//}
//
//var labelKeyPrefix = "easyai.mihoyo.com/"
//
//// CreateTFJobSingle create
//func CreateTFJobSingle() {
//	config := "/Users/bilibili/.arena/kubeconfig"
//	cfg, clt, _ := NewClientSet(config)
//
//	cleanPodRunning := commonv1.CleanPodPolicyRunning
//	image := "registry.cn-shanghai.aliyuncs.com/shuaiyy/mihoyo-ai:arena-tf1.5-py27-gpu-sample-code"
//	cmd := "python code/arena-tensorflow-sample-code/tfjob/docker/mnist/main.py --max_steps 5000"
//	var userRoot int64
//	chiefNum := int32(1)
//
//	var job = &v1.TFJob{
//		TypeMeta: metav1.TypeMeta{
//			Kind:       "TFJob",
//			APIVersion: "v1",
//		},
//		ObjectMeta: metav1.ObjectMeta{
//			Name:            "test-tf-job",
//			Namespace:       "default",
//			Labels:          map[string]string{labelKeyPrefix + "owner": "shuai.yang"},
//			Annotations:     map[string]string{labelKeyPrefix + "owner": "shuai.yang", "istio/inject": "false"},
//			OwnerReferences: nil,
//		},
//		Spec: v1.TFJobSpec{
//			RunPolicy: commonv1.RunPolicy{
//				CleanPodPolicy:          &cleanPodRunning,
//				TTLSecondsAfterFinished: nil,
//				ActiveDeadlineSeconds:   nil,
//				BackoffLimit:            nil,
//				SchedulingPolicy:        nil,
//			},
//			SuccessPolicy: nil,
//			TFReplicaSpecs: map[commonv1.ReplicaType]*commonv1.ReplicaSpec{
//				v1.TFReplicaTypeChief: {
//					Replicas: &chiefNum,
//					Template: corev1.PodTemplateSpec{
//						ObjectMeta: metav1.ObjectMeta{
//							Name: "test-tf-job-pod-chief-0",
//						},
//						Spec: corev1.PodSpec{
//							Volumes: []corev1.Volume{{
//								Name:         "train-data",
//								VolumeSource: corev1.VolumeSource{PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{ClaimName: "ml-data-shuai-yang"}},
//							}},
//							InitContainers: nil,
//							Containers: []corev1.Container{
//								{
//									Name:       v1.DefaultContainerName,
//									Image:      image,
//									Command:    []string{"bash", "-c"},
//									Args:       []string{cmd},
//									WorkingDir: "/root",
//									Env:        []corev1.EnvVar{{Name: "TEST_TMPDIR", Value: "code/arena-tensorflow-sample-code"}},
//									Resources: corev1.ResourceRequirements{
//										Requests: corev1.ResourceList{
//											"cpu":            resource.MustParse("1"),
//											"nvidia.com/gpu": resource.MustParse("1"),
//											"memory":         resource.MustParse("1Gi"),
//										},
//										Limits: corev1.ResourceList{
//											"cpu":            resource.MustParse("10"),
//											"nvidia.com/gpu": resource.MustParse("1"),
//											"memory":         resource.MustParse("10Gi"),
//										},
//									},
//									VolumeMounts: []corev1.VolumeMount{
//										{
//											Name:      "train-data",
//											ReadOnly:  true,
//											MountPath: "/workspace/data",
//										},
//									},
//									Ports: []corev1.ContainerPort{{ContainerPort: 20000, Name: v1.DefaultPortName}},
//									// VolumeDevices:            nil,
//									// LivenessProbe:            nil,
//									// ReadinessProbe:           nil,
//									// StartupProbe:             nil,
//									// Lifecycle:                nil,
//									// TerminationMessagePath:   "",
//									// TerminationMessagePolicy: "",
//									ImagePullPolicy: corev1.PullAlways,
//									SecurityContext: &corev1.SecurityContext{
//										RunAsUser:  &userRoot,
//										RunAsGroup: &userRoot,
//									},
//									Stdin:     false,
//									StdinOnce: false,
//									TTY:       false,
//								},
//							},
//							EphemeralContainers:           nil,
//							RestartPolicy:                 "",
//							TerminationGracePeriodSeconds: nil,
//							ActiveDeadlineSeconds:         nil,
//							DNSPolicy:                     "",
//							NodeSelector:                  nil,
//							ServiceAccountName:            "",
//							HostNetwork:                   false,
//							HostPID:                       false,
//							HostIPC:                       false,
//							ShareProcessNamespace:         nil,
//							SecurityContext:               nil,
//							ImagePullSecrets:              nil,
//							Hostname:                      "",
//							Subdomain:                     "",
//							Affinity:                      nil,
//							SchedulerName:                 "",
//							Tolerations: []corev1.Toleration{
//								{Key: "gpu-pod", Operator: corev1.TolerationOpExists},
//							},
//							HostAliases:               nil,
//							PriorityClassName:         "",
//							Priority:                  nil,
//							DNSConfig:                 nil,
//							ReadinessGates:            nil,
//							RuntimeClassName:          nil,
//							EnableServiceLinks:        nil,
//							PreemptionPolicy:          nil,
//							Overhead:                  nil,
//							TopologySpreadConstraints: nil,
//							SetHostnameAsFQDN:         nil,
//						},
//					},
//					RestartPolicy: commonv1.RestartPolicyNever,
//				},
//			},
//			EnableDynamicWorker: false,
//		},
//	}
//	_ = clt.RESTClient()
//	createJob(cfg, job)
//}
//
//// NewRestClient new
//func NewRestClient() *rest.RESTClient {
//	// RESTClient
//	// config
//	config, err := clientcmd.BuildConfigFromFlags("", "/Users/bilibili/.arena/kubeconfig")
//	if err != nil {
//		panic(err)
//	}
//
//	// 不同的gvk 要设置对应的schema
//	config.GroupVersion = &corev1.SchemeGroupVersion
//	config.NegotiatedSerializer = scheme.Codecs
//	config.APIPath = "/api"
//	config.APIPath = "/apis"
//	config.APIPath = ""
//	// 默认的限流是 5 qps
//	config.QPS = 5.0
//
//	// client
//	restClient, err := rest.RESTClientFor(config)
//	if err != nil {
//		panic(err)
//	}
//	return restClient
//}
