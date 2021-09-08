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
容器创建完毕，此时默认创建的私链已经启动。
![image](https://user-images.githubusercontent.com/86407596/132462791-0e4de6fe-78fc-4883-baca-2abc5341fd0d.png)
```
此时访问 http://localhost:8080/ 即可查看到私链数据。
```
如果你想创建自定义私链，或者连接主网/测试网，请查看**Usages**部分

## Usages
```
filetree 
├── NeoExplorer
├── neo3fura
├── neocli
├── LICENSE
├── README.md
└── deploy.sh
```

**Neo-cli 私链使用方法**

