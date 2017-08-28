const db = new Mongo().getDB('invoicer')

print('dropping users')
let result = db.users.drop()
print(result)

print('creating index')
result = db.users.createIndex({ email: 1 }, { unique: true })
printjson(result)

print('inserting fred')
const insertresult = db.users.insertOne(
  {
    "firstname": "Fred",
    "lastname": "Flintstone",
    "email": "fred@bedrock.gov",
    "isAdmin": true,
    "password": "$2a$10$fjLgV3E0xjS54.AdDNEX4.ZeYfD1oqhzkJrVuNi82YVPGOa9gLGtu"
  }
)
printjson(insertresult)
