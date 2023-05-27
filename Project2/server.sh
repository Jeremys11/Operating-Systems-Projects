#Writer Servers
go run tokenserver/tokenserver.go -port 50001
go run tokenserver/tokenserver.go -port 50002
#Read Servers
go run tokenserver/tokenserver.go -port 60001
go run tokenserver/tokenserver.go -port 60002
go run tokenserver/tokenserver.go -port 60003