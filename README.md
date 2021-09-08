# Neo-Explorer-Tool
This is a project that integrates neo-cli, neo-modules-fura, neo3fura, and neo-explorer. You could configure and start a neo chain, view and query chain data through API service, mongoDb and blockchain browser.
* [neo-cli](https://github.com/neo-project/neo-node): A tool which could configure your node and connect to mainnet/testnet or start your own private chain.
* [neo-modules-fura](https://github.com/neo-ngd/neo-modules-fura): A plugin which could record chain data from node into mongodb. 
* [neo3fura](https://github.com/lutianzhou001/neo3fura): An api service for quickly getting data from the Neo.
* [neo-explorer](https://github.com/RookieCoderrr/Neo-Explorer-UI): A blockchain browser which could visulize your chain data.

## 目录

- [Requirements](#Requirements)
- [Tools recommanded](#Tools recommanded)
- [Get started](#Get started)
- [Usages](#Usages)
  - [Neo-cli usage and configration](#Neo-cli usage and configration)
  - [Mongodb usage](#Mongodb usage)
  - [API service usage](#API service usage)

### Requirements
Neo-Explorer-Tool requires [docker](https://www.docker.com/products/docker-desktop) to run, all of these modules are built in docker.

### Tools recommanded
* [docker-desktop](https://www.docker.com/products/docker-desktop)
The GUI tool for docker, you can operate the container on the UI interface such as enter the container command line, view the container status, restart and delete the container, etc.

* [Mongo Compass](https://www.mongodb.com/products/compass) 
The GUI for MongoDB, you can view the chain data recorded in default database.

* [Postman](https://www.postman.com/)
An API platform for building and using APIs, you can send http request to obtain your chain data.

### Getting started
Downdload project
```
git clone https://github.com/RookieCoderrr/Neo-Explorer-Tool.git
```
Start project 
```
cd Neo-Explorer-Tool
```
```
./delpy.sh 
```

![image](https://user-images.githubusercontent.com/86407596/132462791-0e4de6fe-78fc-4883-baca-2abc5341fd0d.png)

Docker images and respective containers built successfully.

Now, the default private chain created has been started, and you can view the private chain data in your browser http://localhost:8080/.

If you want to create your own private chain, or connect to the mainnet/testnet, please follow the [Usages](#Usages) part

### Usages

#### Neo-cli usage and configration**

Enter the neo-rc2-private container command line 
```
docker exec -it neo-rc2-private /bin/bash
```
Switch to the neo-cli command line
```
screen -r node 
```
Now you are in the neo-cli command line, a default wallet has been opened, which located at /neo-cli/defaultWallet.json with password 111111. Its rpc port is 20332, and the network id is 1234567890. you can operate cli command on your node such as “list asset" etc. For more info, please refer to [neo-cli command](https://docs.neo.org/docs/en-us/node/cli/cli.html).

Attention: If you intend to deploy a contract on the chain, please put your contract manifest file into neocli/SmartContract directory on you local host, the neo-rc2-private container will automatically synchronize the local SmartContract directory to its /SmartContract directory. 

neo-rc2-private容器的命令行
```
按住 control+A 再按 D，此时私链在后台运行
```
如果想再次进入neo-cli指令界面
```
screen -r 
```

How to configure your own chain ?

Please refer to [Build a private chain](https://docs.neo.org/docs/en-us/develop/network/private-chain/solo.html)

Attention: All of your modification and configurantion should be done in your local host (not in neo-rc2-private container which will cause data confusion)

After your configuration, excute
```
./delpy.sh 
```
to rebuild this project.

**Tips:**

Switch to docker command line from neo-cli command line, and neo-cli run in background.
```
Press control+A then D，
```
Switch to neo-cli command line from docker command line
```
screen -r
```
If you shut up neo-cli accidentally, run this command to reenter neo-cli 
```
dotnet neo-cli.dll
```


#### Mongodb usage**


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

#### API service usage*

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

