#!/bin/bash

# Regular Colors
Black='\033[0;30m'        # Black
Red='\033[0;31m'          # Red
Green='\033[0;32m'        # Green
Yellow='\033[0;33m'       # Yellow
Blue='\033[0;34m'         # Blue
Purple='\033[0;35m'       # Purple
Cyan='\033[0;36m'         # Cyan
White='\033[0;37m'        # White
Color_Off='\033[0m'       # no color

SCRIPT_DIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
echo $SCRIPT_DIR
cd $SCRIPT_DIR
echo -e "######## ${Green}Step 1${Color_Off} ########"
echo "准备安装seelie client命令行工具"

if [[ ! -f ./seelie ]]; then
  echo "install.sh所在文件夹下未发现seelie client binary, $SCRIPT_DIR"
  exit 1
fi

echo -e "执行 ${Green}sudo rm -rf /usr/local/bin/seelie && sudo cp ./seelie  /usr/local/bin/${Color_Off}"
echo -e "${Red}可能需要您为sudo提供密码${Color_Off}"
sudo rm -rf /usr/local/bin/seelie && sudo cp ./seelie /usr/local/bin/

echo -e "${Blue}已安装到 /usr/local/bin/seelie${Color_Off}"

echo -e "######## ${Green}Step 2${Color_Off} ########"
echo "配置seelie client"

echo -e "请输入您在seelie平台的${Green}用户token${Color_Off}，(token可在平台主页右上角获取，https://ml.ssr.mihoyo.com):"

read token

if [[ "$token" == "" ]]; then
  echo -e "${Red}invalid token${Color_Off}"
fi

echo -e "执行命令：${Green}seelie config init -H ml.ssr.mihoyo.com -P 443 -p https -t $token${Color_Off}"

seelie config init -H ml.ssr.mihoyo.com -P 443 -p https -t $token


echo -e "######## ${Green}Step 3${Color_Off} ########"
echo "README.md里有更多使用示例"
echo "查看客户端内置文档: "
echo "seelie --help"
echo "seelie submit tf --help"
seelie --help

echo -e "######## ${Green}Step 4${Color_Off} ########"
echo -e "${Red}2分钟${Color_Off}后自动退出，您也可以手动关闭本窗口"
sleep 120