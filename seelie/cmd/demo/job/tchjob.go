package job

//
//import (
//	"context"
//
//	commonv1 "github.com/kubeflow/common/pkg/apis/common/v1"
//	tchv1 "github.com/kubeflow/training-operator/pkg/apis/pytorch/v1"
//	corev1 "k8s.io/api/core/v1"
//	"k8s.io/apimachinery/pkg/api/resource"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//)
//
//// CreateTchJobSingle create
//func CreateTchJobSingle() {
//	config := "/Users/bilibili/.arena/kubeconfig"
//	cls := NewCluster(config)
//	clt := cls.GetClient()
//	cleanPodRunning := commonv1.CleanPodPolicyRunning
//	image := "registry.cn-shanghai.aliyuncs.com/shuaiyy/mihoyo-ai:arena-tf1.5-py27-gpu-sample-code"
//	cmd := "sleep 2h"
//	var userRoot int64
//	chiefNum := int32(1)
//
//	var job = &tchv1.PyTorchJob{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:            "test-pytorch-job",
//			Namespace:       "default",
//			Labels:          map[string]string{labelKeyPrefix + "owner": "shuai.yang"},
//			Annotations:     map[string]string{labelKeyPrefix + "owner": "shuai.yang", "istio/inject": "false"},
//			OwnerReferences: nil,
//		},
//		Spec: tchv1.PyTorchJobSpec{
//			RunPolicy: commonv1.RunPolicy{
//				CleanPodPolicy:          &cleanPodRunning,
//				TTLSecondsAfterFinished: nil,
//				ActiveDeadlineSeconds:   nil,
//				BackoffLimit:            nil,
//				SchedulingPolicy:        nil,
//			},
//			PyTorchReplicaSpecs: map[commonv1.ReplicaType]*commonv1.ReplicaSpec{
//				tchv1.PyTorchReplicaTypeMaster: {
//					Replicas: &chiefNum,
//					Template: corev1.PodTemplateSpec{
//						Spec: corev1.PodSpec{
//							Volumes:        nil,
//							InitContainers: nil,
//							Containers: []corev1.Container{
//								{
//									Name:       tchv1.DefaultContainerName,
//									Image:      image,
//									Command:    []string{"bash", "-c"},
//									Args:       []string{cmd},
//									WorkingDir: "/root",
//									Env:        []corev1.EnvVar{{Name: "TEST_TMPDIR", Value: "code/arena-tensorflow-sample-code"}},
//									Resources: corev1.ResourceRequirements{
//										Requests: corev1.ResourceList{
//											"cpu": resource.MustParse("1"),
//											//"nvidia.com/gpu": resource.MustParse("1"),
//											"memory": resource.MustParse("1Gi"),
//										},
//										Limits: corev1.ResourceList{
//											"cpu": resource.MustParse("10"),
//											//"nvidia.com/gpu": resource.MustParse("1"),
//											"memory": resource.MustParse("10Gi"),
//										},
//									},
//									// VolumeMounts: []corev1.VolumeMount{
//									//	{
//									//		Name:      "train-data",
//									//		ReadOnly:  true,
//									//		MountPath: "/workspace/data",
//									//	},
//									//},
//									Ports: []corev1.ContainerPort{
//										{ContainerPort: 20000, Name: tchv1.DefaultPortName},
//										{ContainerPort: 20001, Name: "seelie-debug"},
//									},
//
//									// VolumeDevices:            nil,
//									// LivenessProbe:            nil,
//									// ReadinessProbe:           nil,
//									// StartupProbe:             nil,
//									// Lifecycle:                nil,
//									// TerminationMessagePath:   "",
//									// TerminationMessagePolicy: "",
//									ImagePullPolicy: corev1.PullIfNotPresent,
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
//			ElasticPolicy: nil,
//		},
//	}
//	if err := clt.Create(context.TODO(), job); err != nil {
//		panic(err)
//	}
//}
