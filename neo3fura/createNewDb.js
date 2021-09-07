use pChainDb
db.createUser({user: "pChain",pwd: "123",roles: [{ role:"dbOwner",db: "pChainDb" }]});