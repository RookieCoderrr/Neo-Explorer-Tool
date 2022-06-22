#!/bin/sh

echo 'Shut down existed neo-cli================'

docker stop neo-rc2-private

echo 'Start Neofura module=================='

cd neo3fura

sh start.sh TEST

cd ..

echo 'Start NeoExplorer Module=================='

cd NeoExplorer

sh run.sh

cd ..

echo 'Creating db and user for you private chain=================='

cd neo3fura

docker cp createNewOnlineDb.js mongodb1_test:/
docker cp createNewOnlineDb.sh mongodb1_test:/

docker cp createNewLocalDb.js mongo_test:/
docker cp createNewLocalDb.sh mongo_test:/

sleep 20s

docker exec -it mongodb1_test bash createNewOnlineDb.sh

docker exec -it mongo_test bash createNewLocalDb.sh

cd ..

echo 'Start Neocli Module=================='

cd neocli

sh run.sh

cd ..
