# kDeck-Server

## Build Proto

`protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false .\proto\data.proto`
