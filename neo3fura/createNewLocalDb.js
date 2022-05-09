
use job
db.dropDatabase()

use job
db.createUser({user: "pChain3",pwd: "123",roles: [{ role:"dbOwner",db: "job" }]});
show
