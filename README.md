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
```
具体链操作请查看[neo-cli命令参考](https://docs.neo.org/docs/zh-cn/node/cli/cli.html).
```
返回neo-rc2-private容器的命令行
```
按住 control+A 再按 D，此时私链在后台运行
```
如果想再次进入neo-cli指令界面
```
screen -r 
```

