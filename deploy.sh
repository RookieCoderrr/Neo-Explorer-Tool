#!/bin/sh

echo 'Start Neofura module=================='

cd neo3fura

sh start.sh TEST

echo 'Neofura Start Success'

cd ..

echo 'Start Neocli Module=================='

cd neocli

sh run.sh

echo 'Neocli Start Success'

cd ..

echo 'Start NeoExplorer Module=================='

cd NeoExplorer

sh run.sh

echo 'Neocli NeoExplorer Success'

cd ..

echo 'Create db and user for you private chian=================='

cd neo3fura

docker cp createNewDb.js mongodb1_test:/
docker cp createNewDb.sh mongodb1_test:/
docker exec -it mongodb1_test bash createNewDb.sh

echo 'Done'