
# task 1:  seelie image build

```shell
# 入口命令
seelie image build --project_dir=./ --base_image=xxx.com/aaa/bbb:tttt --tag=test123 --repo_type=training --backend=local --push

# 结果
## 1. 构建出镜像：  {RegistryHost}/{Repo}/{username}-training:test123
## 2. 将镜像推送到镜像仓库

```

## 构建逻辑

1. context 为 project_dir，根目录下存在`ml.flag`标志文件
2. 如果context路径下存在Dockerfile，则使用Dockerfile构建；否则，构建逻辑为：
   ```shell
   # 伪代码
   Copy . project_dir /workspace/code
   Workdir /workspace
   Run cd  /workspace/code; 
       if exist ./build_script.sh, exec build_script.sh; 
       if exist ./requirements.txt, exec pip install -r requirements.txt;
   ```
3. 构建之前可能需要先检查 base镜像是否存在，以及先拉远程镜像
4. base image合法前缀检查：
   1. 平台的镜像仓库
   2. nvcr 英伟达官方镜像 nvcr.io/nvidia
   3. docker官方镜像: tensorflow/tensorflow, pytorch/pytorch

## backend： local vs remote

1. local且用户电脑上安装了docker：使用[client](https://github.com/moby/moby/tree/master/client)连接docker server，执行build，push逻辑
2. remote： 使用[podman](https://github.com/containers/podman)连接远程的podman server，执行build&push
3. 有多组remote server(不同国际区域)，需要使用对应的sock5代理服务器连接

# task2: commit container image for k8s Pod

使用场景： 1. seelie云上交互式镜像构建 2. 用户开发环境的镜像快照

## 定位pod的容器

```shell
kubectl get pod  jupyter-minimal -o jsonpath='{.status.containerStatuses[0].containerID}'
   containerd://2532ea87fe612bcfe17a9c37721de9acf839149ee92e4e3c97f18e3791370270
kubectl get pod  jupyter-minimal -o jsonpath='{.spec.nodeName}'
  cn-shanghai.10.14.128.24
```
## commit逻辑

1. 在目标机器上创建一个pod，挂载物理机上的容器运行时的sock文件
2. 执行committer程序, 需要编程实现
   1. 连接到容器运行时 docker或containerd
   2. 找到目标容器
   3. 在目标容器里执行clean逻辑(docker exec)，清理临时文件： apt cahce， pip cache，以及 /tmp目录下最近48小时内产生的文件
   4. 执行docker commit
   5. 执行docker push
   6. 执行docker image rm

## 测试环境

使用minikube搭建2套集群，并记录文档

1. k8s1.21.10， docker运行时
2. k8s1.21.10，containerd运行时

```
brew install minikube

minikube start --driver=docker --image-mirror-country='cn' --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers' --registry-mirror='https://registry.docker-cn.com' --base-image="kicbase/stable:v0.0.36" --kubernetes-version=v1.21.10 --container-runtime=docker -p k8s-docker

kubectl version
kubectl get nodes
minikube ssh -p k8s-docker

```

## 在pod容器中访问容器运行时
挂载sock file： /var/run/docker.sock  /var/run/containerd/ /run/containerd/

containerd client: 参考[nerdctl](https://github.com/containerd/nerdctl)

在pod中连接物理机容器运行时：可以参考openkruise实现，
+ [openkruise init criruntime client](https://github.com/openkruise/kruise/blob/af3e254fa09fbdb96084549cb84da2d33c31fa95/pkg/daemon/criruntime/imageruntime/containerd.go)
+ openkruise使用cri接口，实现镜像预热和容器原地重启，cri接口规范里没有container commit，因此我们需要原生的容器运行时client

# task3： k8s云构建技术调研

## 背景
不依赖docker，在k8s中构建用户镜像
context：挂载nas存储
构建逻辑：同`seelie image build`，但是不依赖容器运行时
构建工具，tool that can build container images in Kubernetes clusters:

- [Kaniko](https://github.com/GoogleContainerTools/kaniko)
- [Cloud Native Buildpacks](https://buildpacks.io/)
- [BuildKit](https://github.com/moby/buildkit)
- [Buildah](https://buildah.io/)

有一个整合以上工具的build项目： [shipwright-io/build](https://github.com/shipwright-io/build)

## 落地形式

实现一个Pod定义，仅改变部分参数即切换不同工具，实现镜像构建和推送

输入: 
+ base image
+ build context path: 挂载远程存储到pod里
+ build tool
+ target image
+ hub secrets

输出：
+ 构建镜像成功，并推送到镜像仓库

# task4： seelie job submit command

- seelie client已经支持通过命令行提交job
- 现在，给定一个job id，实现通过seelie client，逆向出提交命令

```shell
# ./seelie/tarball/README.md
seelie submit tfjob --name "test job" --description "submit by seelie cli" \
   --cluster dev --namespace default --cpu 2 --memory 8 --gpu 0 --worker-count 1 \
   --image registry.cn-shanghai.aliyuncs.com/shuaiyy/mihoyo-ai:tf2.4.3-gpu-jupyter-lab \
   --entrypoint-type "bash -c" --entrypoint "sleep 10m; echo failed; exit 1" \
   -E enable_ema=1
   
# 待实现

seelie job get --job_id 123 -o text|json|yaml|cmd

## 其中cmd方式为 提交命令

```