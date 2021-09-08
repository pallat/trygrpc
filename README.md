# gRPC in Go

https://developers.google.com/protocol-buffers/docs/gotutorial

### helloworld.proto

```proto3
syntax = "proto3";
package helloworld;

option go_package = "github.com/protocolbuffers/protobuf/examples/go/helloworldpb";

// The greeting service definition.
service Greeter {
    // Sends a greeting
    rpc SayHello (HelloRequest) returns (HelloReply) {}
    // Sends another greeting
    rpc SayHelloAgain (HelloRequest) returns (HelloReply) {}
  }
  
  // The request message containing the user's name.
  message HelloRequest {
    string name = 1;
  }
  
  // The response message containing the greetings
  message HelloReply {
    string message = 1;
  }
```

## install

### protoc

https://grpc.io/docs/protoc-installation/

```sh
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```

### protoc-gen-go

this not generate for gRPC

> go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

### protoc-gen-go-grpc

> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## compile proto

> protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

### deprecated

> protoc -I=. --go_out=plugins=grpc:. ./addressbook.proto

### currently use

> protoc --go_out=. --go-grpc_out=. ./helloworld.proto
