# Neo-Explorer-Tool
This is a project that integrates neo-cli, neo-modules-fura, neo3fura, and neo-explorer. You could configure and start a neo chain, view and query chain data through API service, mongoDb and blockchain browser.
* [neo-cli](https://github.com/neo-project/neo-node): A tool which could configure your node and connect to mainnet/testnet or start your own private chain.
* [neo-modules-fura](https://github.com/neo-ngd/neo-modules-fura): A plugin which could record chain data from node into mongodb. 
* [neo3fura](https://github.com/lutianzhou001/neo3fura): An api service for quickly getting data from the Neo.
* [neo-explorer](https://github.com/RookieCoderrr/Neo-Explorer-UI): A blockchain browser which could visulize your chain data.

## Table of Contents

*[Requirements](#Requirements)
* [Tools recommanded](#Tools recommanded)
* [Get started](#Get started)
* [Usages](#Usages)
  * [Neo-cli](#Neo-cli)
  * [Mongodb ](#Mongodb)
  * [API service ](#API service )
  * [Blockchain browser](#Blockchain browser)

### Requirements
Neo-Explorer-Tool requires [docker](https://www.docker.com/products/docker-desktop) to run, all of these modules are built in docker.

### Tools recommanded
* [docker-desktop](https://www.docker.com/products/docker-desktop),
The GUI tool for docker, you can operate the container on the UI interface such as enter the container command line, view the container status, restart and delete the container, etc.

* [Mongo Compass](https://www.mongodb.com/products/compass) ,
The GUI for MongoDB, you can view the chain data recorded in default database.

* [Postman](https://www.postman.com/),
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

Docker images and respective containers are built successfully as shown above. Now, a default private chain is running in background in the container neo-cr2-private, and you can view the private chain data in your browser http://localhost:8080/. 

If you want to create your own private chain, or connect to the mainnet/testnet, please follow the [Usages](#Usages) part

### Usages

#### Neo-cli 

* Usage

Enter the neo-rc2-private container command line 
```
docker exec -it neo-rc2-private /bin/bash
```
Switch to the neo-cli command line
```
screen -r node 
```
Now you are in the neo-cli command line, a default wallet has been opened.

It locates at **Neo-Explorer-Tool/neo-cli/defaultWallet.json** with password **111111**. 

Its rpc port is **20332**, and the network id is **1234567890**. 

you can excute cli command on your node such as “list asset" etc. For more info, please refer to [neo-cli command](https://docs.neo.org/docs/en-us/node/cli/cli.html).

Attention: If you intend to deploy a contract on the chain, please put your contract manifest file into Neo-Explorer-Tool/neocli/SmartContract directory on you local host, the neo-rc2-private container will automatically synchronize the local SmartContract directory to its /SmartContract directory. 

* Configuration

How to configure your own chain ?

Please refer to [Build a private chain](https://docs.neo.org/docs/en-us/develop/network/private-chain/solo.html)

Attention: All of your modification and configurantion should be done in your local host (not in neo-rc2-private container which will cause data confusion)

After your configuration, excute
```
./delpy.sh 
```
to rebuild this project.

* Tips

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


#### Mongodb 


* connect to mongodb in the container 


enter into docker command line in mongodb1_test container

```
docker exec -it mongodb1_test /bin/bash
```

then enter into mongo command line

```
mongo
```

enter into database 
```
show dbs
```
```
use pChainDb3
```

For more info about mongo command line, please refer to [mongo manual](https://docs.mongodb.com/v4.4/mongo/)

* connect to mongodb in your host 

database configuration
```
Host:localhost
Port:27001
Username:pChain3
Password:123
Databse:pchainDb3
```
You could use [Mongo command](https://docs.mongodb.com/manual/tutorial/getting-started/) ，or [Mongo Compass](https://www.mongodb.com/products/compass) to check chain data .

#### API service

* use curl command to send http request

request format
```
curl --location --request POST 'http://127.0.0.1:1926' \ --header 'Content-Type: application/json' \ --data-raw '{ "jsonrpc": "2.0", "method": "GetBestBlockHash", "params": {}, "id": 1 }'
```
Refer to [Neofura-Doc](https://neo3phora-doc.readthedocs.io/en/latest/docs/APIs/index.html for more info.

response format
```
{ "id": 1, "result": { "hash": "0x226844a595780dd2881bbfedbf4ffabe25fcc691969359aa1b7f87a715cdea75" }, "error": null }
```
The url and port exposed is http://127.0.0.1:1926

#### Blockchain browser

The url and port exposed is http://localhost:8080/

