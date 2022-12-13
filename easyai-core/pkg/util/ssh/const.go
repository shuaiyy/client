package ssh

// default ssh key-pair
const (
	DefaultPrivateKey = `-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----
`
	DefaultPublicKey     = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCaMzxKhnsj65HkadTTbop8BEYcI2VJFAhZOn34/nMo6jPFwpAHyd7Sfh38kNkQqh3mf6Qo6VJEb9dLjAs0QAGHx23CWJ0jP1sDIPF/aR+eulnJKLcAVPa8moKH0/HmONGiz3sfdfTrVokPtvjCmAsBBBxWK1VTgTXej8gat1MS32+PrsE2TuNwUh+PDkpPyKwLWbTvLSc0taFv5awOo2Lote9vjpTEMLgF6G+2MKp+JLnqz9o4zZkBcD6yld3I2aBQ+6sWQVkyMyUnpfRuWA1wtklyvnhf6K9oSREt6DaezBQCxjcgbEcur72XAcRf/sZUKOpghoVqepNtWNhRSgCr+LviBzy51UMqhSCz10O7yV829XPX41YUtibjZM0gnP6SCZdq36GoRTa4pzBItQq6f9qPwj7Sq4Bwf3jI2gVz6zpzkt+FOLQ9h2ZXdLR9ASB6BeoWf5pmf7QjK1nZ/tfBR00ZA7ydo2AG+Rs+lv7CicJTfwsntMwHcBO6iR3MOqk="
	DefaultHostPublicKey = "ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAzfg/8Z0VmJfK2n8ZNvzFZElaRzyEEgfq+nQjLPxryqdfEb0OKVIF5BbTvZAIJEa0P0CW02b/DCafHc+MqJzeU=\n"
)

const scriptTemplate = `#!/bin/bash

export insName="%s"
export insIP="%s"
export insPort="%s"
export socks5Addr="%s"
export socks5Auth="%s"
export sshPrivateKey="%s"
export hostPubKey="%s"

which ncat >/dev/null 2>&1 || { echo "ncat is not installed, please install it first and retry.";
echo "for macos: brew install nmap";
echo "for linux(Ubuntu/Debian): sudo apt-get install -y ncat";
echo "for linux(Centos/Redhat): sudo yum install -y nmap-ncat";
exit 1; }

read -e -p "请输入ssh连接名 (press enter for '${insName}'): " SSH_CONNECTION_NAME; \
if [[ $SSH_CONNECTION_NAME == "" ]]; then SSH_CONNECTION_NAME=${insName}; fi;


naming_re='^[a-z0-9_-]+$'

if ! [[ "$SSH_CONNECTION_NAME" =~ $naming_re ]]; then
    echo "[WARNING] 非法连接名， should only contain lowercase characters, numbers and - or _. Please run again and input a valid name.";
    exit 1;
fi

# Create ssh folder in case it does not exist
mkdir -p ~/.ssh && chmod 700 ~/.ssh
touch ~/.ssh/config
touch ~/.ssh/known_hosts

if ! grep -q "Host $SSH_CONNECTION_NAME" ~/.ssh/config; 
then 
    # Setup ssh config information
	echo "" >> ~/.ssh/config
    echo "
Host $SSH_CONNECTION_NAME
  Hostname ${insIP}
  Port ${insPort}
  User root
  ServerAliveCountMax 10
  ServerAliveInterval 60
  IdentityFile ~/.ssh/$SSH_CONNECTION_NAME
  IdentityFile ~/.ssh/id_rsa
  IdentityFile ~/.ssh/seelie_job_default_key
  ProxyCommand ncat --proxy-type socks5 --proxy ${socks5Addr} --proxy-auth ${socks5Auth} %%h %%p
" >> ~/.ssh/config;

    # Setup private key of runtime
    echo "${sshPrivateKey}" > ~/.ssh/$SSH_CONNECTION_NAME 
    chmod 700 ~/.ssh/$SSH_CONNECTION_NAME
else
    echo "[WARNING] A connection with the name $SSH_CONNECTION_NAME already exists in ~/.ssh/config. Try a different name."
    exit 1
fi;

if ! grep -q "\[${insIP}\]:${insPort}" ~/.ssh/known_hosts; 
then 
    echo "[${insIP}]:${insPort} ${hostPubKey}" >> ~/.ssh/known_hosts; 
else
    echo "[WARNING] DNS/IP and Port([${insIP}]:${insPort}) 已经在 ~/.ssh/known_hosts 中, 由于Seelie Job的任务IP是动态分配回收的，该记录可能是历史容器的身份信息，如果提示异常，请删除该记录后在尝试. 如果还不行，添加'StrictHostKeyChecking no' to the 'Host $SSH_CONNECTION_NAME' in the ~/.ssh/config";
fi;

# Test the connection
echo "Testing the SSH connection via 'ssh $SSH_CONNECTION_NAME'"
ssh -q $SSH_CONNECTION_NAME exit
if [ $? == 0 ]; then 
    echo "Connection successful!"

    # Setup jupyter remote kernel if remote_ikernel is installed
    if hash remote_ikernel 2>/dev/null; then
        while true; do
            read -p "remote_ikernel was detected on your machine. Do you like to setup a Python remote kernel for Jupyter (yes/no)? " yn
            case $yn in
                [Yy]* ) remote_ikernel manage --add --interface=ssh --kernel_cmd="ipython kernel -f {connection_file}" --name="Py 3.6" --host=$SSH_CONNECTION_NAME; break;;
                [Nn]* ) break;;
                * ) echo "Please answer yes or no.";;
            esac
        done
    fi;

    # Setup SFTP bookmarks in file explorer
    BOOKMARKS_FILE=~/.config/gtk-3.0/bookmarks
    if [ -f "$BOOKMARKS_FILE" ]; then
        # bookmark file detected
        while true; do
            read -p "Do you want to add this connection as mountable SFTP storage to the bookmarks of your file manager (yes/no)? " yn
            case $yn in
                [Yy]* ) printf "\nsftp://$SSH_CONNECTION_NAME/workspace/ $SSH_CONNECTION_NAME" >> $BOOKMARKS_FILE; break;;
                [Nn]* ) break;;
                * ) echo "Please answer yes or no.";;
            esac
        done
    fi

    # TODO Add additional features: e.g. autossh port forwarding for main workspace port, sshfs folder mounting

    # print out some user information
    echo "The ssh configuration is completed successfully. You can now securely connect via 'ssh $SSH_CONNECTION_NAME'.";
else
    echo "[WARNING] Connection test not successful! Please check the ssh setup manually within the ~/.ssh/config"
fi
`
