## Mongodb

### Mongodb Shell

### Install mongo on ubuntu
https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/

### Run mongodb in mac
```
 mongod --config /usr/local/etc/mongod.conf
```

### Run mongodb service by brew
```
brew services start mongodb-community@4.0
```

### Run mongo Shell
```
mongo
```

#### Get a collection's documents
```
use bus
db.getCollection("stats").find()


db.incomes.findOne( { line: {line_no: 20 }, income: {cash: NumberLong(100000)}} )
db.incomes.findOne( { line: {line_no: 20 }, income: {cash: { $eq: 100000}}} )

db.incomes.findOne( { line: {line_no: 20 }, income: {income_month: new ISODate("2000-01-01T00:00:00Z")}} )

db.incomes.deleteOne( { status: "D" } )

```

## Golang

### Init project with go modular 
```
go mod init github.com/galahade/bus_incomes
```

### Test a Method
```
cd ~/git/go/src/github.com/galahade/bus_incomes/domain
go test -run TestMongoDBInsertMany
```

### Install code to go/lib
```
cd ~/git/go/src/github.com/galahade/bus_incomes
go install -i
```


### Run app
```
bus_incomes -log_dir=log -alsologtostderr

# On server

bus_incomes -log_dir=log &
```
