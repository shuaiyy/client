# Seelie Client 安装与使用教程

> [Seelie机器学习平台](https://ml.ssr.mihoyo.com)

# 安装和配置token

1. 客户端下载: 平台主页右上角
2. token获取: 平台主页右上角
3. 安装并**配置客户端**

   ```bash
   # 解压
   tar -zxvf seelie-client-xxx.tar.gz
   # 安装
   sudo mv seelie /usr/local/bin
   
   ##  macos首次执行seelie命令，可能需要在 `系统偏好设置` 》`安全性和隐私` 里授权允许运行
   # 生成配置文件
   seelie config --help
   
   # token可以在平台主页右上角获取
   seelie config init -H ml.ssr.mihoyo.com -P 443 -p https -t 我的token
   
   # 配置自动补全，非必须 
   ## 对于使用bash shell的用户
   seelie completion bash --help
   ## 对于使用zsh shell的用户
   seelie completion zsh --help
   # 根据提示操作后，重新打开终端窗口即可
   ```

# 使用

1. 查看帮助文档
   ```bash
   # 每一级子命令都带帮助文档和example🌰
   seelie -h
   seelie submit -h 
   seelie submit tfjob -h
   ```

2. 注册NAS浏览器用户

   > seelie后端所在的网络环境无法访问计算集群所在的nas存储，因此需要用户手动执行注册命令

   ```shell
   seelie data register-user
   # 如下输出，提示注册成功或者已注册
   Using config file: /Users/shuai.yang/.seelie/config.yaml
   2022/12/06 11:56:25 [INFO] cluster: aws-dev, nas storage for data: https://ml-dev-aws.eks.hoyoverse.com/nfsdata
   2022/12/06 11:56:25 [WARN] user<shuai.yang> already exist for nas storage(data)
   2022/12/06 11:56:25 [INFO] cluster: aws-dev, nas storage for jupyter: https://ml-dev-aws.eks.hoyoverse.com/nfsjupyter
   2022/12/06 11:56:26 [WARN] user<shuai.yang> already exist for nas storage(jupyter)
   2022/12/06 11:56:26 [INFO] cluster: dev, nas storage for data: https://ml-dev.ssr.mihoyo.com/nfsdata
   2022/12/06 11:56:26 [WARN] user<shuai.yang> already exist for nas storage(data)
   2022/12/06 11:56:26 [INFO] cluster: dev, nas storage for jupyter: https://ml-dev.ssr.mihoyo.com/nfsjupyter
   2022/12/06 11:56:26 [WARN] user<shuai.yang> already exist for nas storage(jupyter)
   2022/12/06 11:56:26 [INFO] cluster: aliyun-sh, nas storage for data: https://ml-aliyun-sh.ssr.mihoyo.com/nfsdata
   2022/12/06 11:56:26 [WARN] user<shuai.yang> already exist for nas storage(data)
   2022/12/06 11:56:26 [INFO] cluster: aliyun-sh, nas storage for jupyter: https://ml-aliyun-sh.ssr.mihoyo.com/nfsjupyter
   2022/12/06 11:56:26 [WARN] user<shuai.yang> already exist for nas storage(jupyter)
   ```
   
   + `mihoyo.com` 域名下的nas浏览器，自动登录。只要浏览器登录过km或者`https://op.mihoyo.com/#/home`
   + `hoyoverse.com`域名没有sso，需要使用代理登录: `https://ml-dev-aws.eks.hoyoverse.com/seelie/login-assist?token=你的token`

3. 创建单机job
   > 单机job与算法框架无关,里面可以运行任意框架的job；这里我们使用tfjob即可(tensorflow framework)

   ```bash
   seelie submit tfjob --name "test job" --description "submit by seelie cli" \
   --cluster dev --namespace default --cpu 2 --memory 8 --gpu 0 --worker-count 1 \
   --image registry.cn-shanghai.aliyuncs.com/shuaiyy/mihoyo-ai:tf2.4.3-gpu-jupyter-lab \
   --entrypoint-type "bash -c" --entrypoint "sleep 10m; echo failed; exit 1" \
   -E enable_ema=1
   ```

4. 创建TF分布式训练
   > 2 ps + 4 worker;
   > 
   > data_dir=/data_dir/mnist_data  train_steps=30  batch_size=32

   ```shell
   seelie submit tfjob --name "dist-tf-2ps-4worker" --description "submit by seelie cli" \
    --cluster dev --namespace default --cpu 6 --memory 12 --gpu 0 --worker-count 4 \
    --image registry.cn-shanghai.aliyuncs.com/shuaiyy/2233:tf1.5-dist-mnist-demo-train_op1.4 \
    --entrypoint-type "python" --entrypoint "/workspace/dist_mnist.py" \
    -E enable_ema=1 -E env_aaa=test \
    -T data_dir=/data_dir/mnist_data -T train_steps=30 -T batch_size=32 \
    -M m_over_sale=3 -M m_enable_debug_toolbox=true -M m_instance_retain=true -M m_retain_time=1h \
    --ps-count 2 --ps-cpu 4 --ps-memory 20
   ```

5. 查看job
   `seelie job get --job_id=123`
6. stop job
   `seelie job stop --job_id=123`