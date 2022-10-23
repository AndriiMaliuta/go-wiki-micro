


Stack:

- Fiber
- Pgx (v4)
- 

https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get -u google.golang.org/protobuf/proto
#go get -u github.com/golang/protobuf/proto
```

```
protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
```

```
protoc --proto_path=src \
  --go_opt=Mprotos/buzz.proto=example.com/project/protos/fizz \
  --go_opt=Mprotos/bar.proto=example.com/project/protos/foo \
  protos/buzz.proto protos/bar.proto
```
