#!/bin/sh

docker build --no-cache -t neo-private .
docker rm  --force neo-rc2-private
docker run -dit --name=neo-rc2-private -v $PWD/SmartContract:/neo-cli/SmartContract --link mongodb1_test --net network -e INIT_ADDRESS=NhZTTsKwdNQXU9FwadjA8hnU156a8QB5kz  -p 20332-20334:20332-20334 neo-private