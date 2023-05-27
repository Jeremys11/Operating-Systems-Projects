go run tokenclient/tokenclient.go -create -id 1
go run tokenclient/tokenclient.go -create -id 2
go run tokenclient/tokenclient.go -create -id 3
go run tokenclient/tokenclient.go -write -id 1 -name dac -low 0 -mid 10 -high 100
go run tokenclient/tokenclient.go -write -id 2 -name dac -low 0 -mid 10 -high 100
go run tokenclient/tokenclient.go -write -id 3 -name dac -low 0 -mid 10 -high 100
go run tokenclient/tokenclient.go -read -id 1
go run tokenclient/tokenclient.go -read -id 2
go run tokenclient/tokenclient.go -read -id 3