db.isMaster()
use pChainDb3
db.dropDatabase()

use pChainDb3
db.createUser({user: "pChain3",pwd: "123",roles: [{ role:"dbOwner",db: "pChainDb3" }]});
show
