module github.com/galahade/bus_incomes

go 1.12

replace (
	golang.org/x/crypto v0.0.0-20190530122614-20be4c3c3ed5 => github.com/golang/crypto v0.0.0-20190530122614-20be4c3c3ed5
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sync v0.0.0-20190412183630-56d357773e84 => github.com/golang/sync v0.0.0-20190412183630-56d357773e84
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223
)

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/snappy v0.0.1 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/szuecs/gin-glog v1.1.1
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.2
	golang.org/x/crypto v0.0.0-20190530122614-20be4c3c3ed5 // indirect
	golang.org/x/net v0.0.0-20190522155817-f3200d17e092 // indirect
	golang.org/x/sys v0.0.0-20190531132440-69e3a3a65b5b // indirect
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/tools v0.0.0-20190530215528-75312fb06703 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
