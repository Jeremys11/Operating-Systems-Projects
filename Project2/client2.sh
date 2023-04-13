go run tokenclient/tokenclient.go -create -id 1234 -host localhost -port 50051
go run tokenclient/tokenclient.go -read -id 1234 -host localhost -port 50051
go run tokenclient/tokenclient.go -write -id 1234 -name dac -low 0 -mid 10 -high 100 -host localhost -port 50051
go run tokenclient/tokenclient.go -read -id 1234 -host localhost -port 50051
go run tokenclient/tokenclient.go -drop -id 1234 -host localhost -port 50051