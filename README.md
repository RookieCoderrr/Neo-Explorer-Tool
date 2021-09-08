# Neo-Explorer-Tool
这是一个整合了neo-cli私链，neo-modules-fura链数据抓取工具,neo3fura API服务,以及neo-explorer区块链浏览器的工具，通过这个工具，你可以配置启动你自己的私链，通过MongodDb查看私链数据，通过API服务请求私链数据，使用浏览器可视化私链。

## Requirements
Neo-Explorer-Tool requires [docker](https://www.docker.com/products/docker-desktop) to run, all of these modules are built in docker.

## Tools recommanded
* [Mongo Compass](https://www.mongodb.com/products/compass) 
The GUI for MongoDB. Visually explore your data.

* [Postman](https://www.postman.com/)
Postman is an API platform for building and using APIs


## Getting Started
下载项目
```
git clone https://github.com/RookieCoderrr/Neo-Explorer-Tool.git
```
脚本文件启动项目
```
cd Neo-Explorer-Tool
```
```
./delpy.sh 创建docker镜像并启动容器
```
容器创建完毕

![image](https://user-images.githubusercontent.com/86407596/132462791-0e4de6fe-78fc-4883-baca-2abc5341fd0d.png)


此时默认创建的私链已经启动，访问 http://localhost:8080/ 即可查看到私链数据。


如果你想创建自定义私链，或者连接主网/测试网，请查看以下部分。

## Usages

**Neo-cli 私链使用方法**

首先进入neo-rc2-private容器的命令行
```
docker exec -it neo-rc2-private /bin/bash
```
进入neo-cli指令界面
```
screen -r node 
```
neo-cli命令参考

具体链操作请查看[neo-cli命令参考](https://docs.neo.org/docs/zh-cn/node/cli/cli.html).

ps:如果你想在链上部署合约，请把合约manifest文件放入到本地项目neocli/SmartContract目录下,容器会自动把本地SmartContract目录同步到容器/SmartContract目录下。

返回neo-rc2-private容器的命令行
```
按住 control+A 再按 D，此时私链在后台运行
```
如果想再次进入neo-cli指令界面
```
screen -r 
```
默认钱包配置
```
the wallet file located at /neo-cli/defaultWallet.json with password 111111. Rpc port is 20332, and the network id is 1234567890.
```
修改链配置方法

请参考[使用单节点搭建私有链](https://docs.neo.org/docs/zh-cn/develop/network/private-chain/solo.html)

修改好后重新创建neo-cli容器
```
neocli/sh run.sh
```

**Mongodb查看链数据**


* 在docker容器内查看数据


进入mongodb1_test容器的命令行

```
docker exec -it mongodb1_test /bin/bash
```

进入neo-cli指令界面

```
mongo
```

查看链默认数据库
```
show dbs
```
```
use pChainDb3
```

* 在主机电脑查看数据

数据库配置
```
Host:localhost
Port:27001
Username:pChain3
Password:123
Databse:pchainDb3
```
可以使用[Mongo命令行](https://docs.mongodb.com/manual/tutorial/getting-started/)，或者下载[Mongo Compass](https://www.mongodb.com/products/compass)查看。

**API服务访问链数据**

* 使用curl指令 进行Http请求

请求格式
```
curl --location --request POST 'https://testneofura.ngd.network:444' \ --header 'Content-Type: application/json' \ --data-raw '{ "jsonrpc": "2.0", "method": "GetBestBlockHash", "params": {}, "id": 1 }'
```
Refer to [Neofura-Doc](https://neo3phora-doc.readthedocs.io/en/latest/docs/APIs/index.html for more info

返回格式
```
{ "id": 1, "result": { "hash": "0x226844a595780dd2881bbfedbf4ffabe25fcc691969359aa1b7f87a715cdea75" }, "error": null }
```

* 下载[Postman](https://www.postman.com/) 发送请求

请求地址
```
http://127.0.0.1:1926
```

请求格式
```
{
  "jsonrpc": "2.0",
  "method": "GetBestBlockHash",
  "params": {},
  "id": 1
}
```
**浏览器访问链数据**
```
http://localhost:8080/
```

