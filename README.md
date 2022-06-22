# Neo-Explorer-Tool
This is a project that integrates neo-cli, neo-modules-fura, neo3fura, and neo-explorer. You could configure and start a neo chain, view and query chain data through API service, mongoDb and blockchain browser.
* [neo-cli](https://github.com/neo-project/neo-node)
  * Author: [Neo Project](https://github.com/neo-project)
  * Description: A tool which could configure your node and connect to mainnet/testnet or start your own private chain.
* [neo-modules-fura](https://github.com/neo-ngd/neo-modules-fura)
  *  Author: [vikkkko](https://github.com/vikkkko)
  *  Description: A plugin which could record chain data from node into mongodb. 
* [neo3fura](https://github.com/lutianzhou001/neo3fura)
  *  Author: [lutianzhou001](https://github.com/lutianzhou001)
  *  Description:An api service for quickly getting data from the Neo.
* [neo-explorer](https://github.com/RookieCoderrr/Neo-Explorer-UI)
  * Author: [dremofly](https://github.com/dremofly), [WSbaikaishui](https://github.com/WSbaikaishui), [zifanwangsteven](https://github.com/zifanwangsteven), [RookieCoderrr](https://github.com/RookieCoderrr)
  * Description: A blockchain browser which could visulize your chain data.

## Table of Contents

* [Requirements](#Requirements)
* [Tools recommanded](#tools-recommanded)
* [Getting started](#getting-started)
* [Usages](#Usages)
  * [Neo-cli](#Neo-cli-usage)
  * [Mongodb ](#Mongodb)
  * [API service ](#aPI-service )
  * [Blockchain browser](#blockchain-browser)

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
git clone -b neo-tool-jiuquan https://github.com/RookieCoderrr/Neo-Explorer-Tool.git
```
Start project 
```
cd Neo-Explorer-Tool
```
```
./delpy.sh 
```

![image](https://user-images.githubusercontent.com/86407596/132462791-0e4de6fe-78fc-4883-baca-2abc5341fd0d.png)

Docker images and respective containers are built successfully as shown above. There will be 8 docker containers.

* neo-rc2-private is a container for running neo nodes

* vueApp is a container of the browser front-end pages

* service_http_test is the container for the browser API interface

* mongodb1_test, mongodb2_test, mongodb3_test are mongo replica sets containers, of which Mongodb1_test is the primary database of the sets. And mongodb1_test is the only entry to view the node data

* mongodb_test is a mongo container that stores hot data on the chain

* cache_test is the container for the redis cache on the chain

Now, a default private chain is running in background in the container neo-cr2-private, and you can view the private chain data in your browser http://localhost:8080. 

If you want to create your own private chain, or connect to the mainnet/testnet, please follow the [Usages](#Usages) part

### Usages

#### Neo-cli Usage

Enter the neo-rc2-private container command line 
```
docker exec -it neo-rc2-private /bin/bash
```
Neo-cli is running in Screen, use screen command to enter neo-cli.
```
screen -ls
screen -r 名字.node
```
Now you are in the neo-cli command line, a default wallet has been opened.

It locates at Neo-Explorer-Tool/neo-cli/defaultWallet.json with password **111111**. 

Its rpc port is **10332**, and the network id is **4411722**. 

You can excute cli command on your node such as “list asset" etc. 

For more info, please refer to [neo-cli command](https://docs.neo.org/docs/en-us/node/cli/cli.html).

> **Notes:**
>If you intend to deploy a contract on the chain, please put your contract manifest file into Neo-Explorer-Tool/neocli/SmartContract directory on you local host,
the neo-rc2-private container will automatically synchronize the local SmartContract directory to its /SmartContract directory. 


#### Neo-cli Configuration

How to configure your own chain ?  Please refer to [Build a private chain](https://docs.neo.org/docs/en-us/develop/network/private-chain/solo.html).

> **Notes**:
>All of your modification and configurantion should be done in your local host (not in neo-rc2-private container which will cause data confusion)

After your configuration, excute this command to rebuild the project.
```
./delpy.sh 
```

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

For more info about mongo command line, please refer to  [Mongo command](https://docs.mongodb.com/manual/tutorial/getting-started/).

If you prefer GUI interface, [Mongo Compass](https://www.mongodb.com/products/compass) is a good choice, connect to db with these configuration below.

database configuration
```
Host:localhost
Port:27001
Username:pChain3
Password:123
Databse:pchainDb3
```


#### API service

use curl command to send http request

request format
```
curl --location --request POST 'http://127.0.0.1:1926' \ --header 'Content-Type: application/json' \ --data-raw '{ "jsonrpc": "2.0", "method": "GetBestBlockHash", "params": {}, "id": 1 }'
```

response format
```
{ "id": 1, "result": { "hash": "0x226844a595780dd2881bbfedbf4ffabe25fcc691969359aa1b7f87a715cdea75" }, "error": null }
```

Refer to [Neofura-Doc](https://neo3phora-doc.readthedocs.io/en/latest/docs/APIs/index.html) for more API service info.

The url and port exposed for API service is http://127.0.0.1:1926

#### Blockchain browser

The url and port exposed for Blocakchain browser is http://localhost:8080/

