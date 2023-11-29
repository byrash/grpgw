# Generate Proto & GRPC Server
```shell
protoc -I. -I./googleapis --go_out=. \
    --go-grpc_out=. \
    ./protos/hello.proto
```
# Annotations
Leveraging google api based annotations with inline approach in proto.
GRPC Gateway annotations [Examples] (https://github.com/grpc-ecosystem/grpc-gateway/blob/main/examples/internal/proto/examplepb/a_bit_of_everything.proto)

# Generate Gateway
```shell
protoc -I. -I./googleapis --grpc_impl-gateway_out=. \
    ./protos/hello.proto
```

# Download submodules
```shell
git submodule sync
git submodule update --init --recursive --remote
```

# Generate Swagger Spec
```shell
protoc -I. -I./googleapis --openapiv2_out=. \
    ./protos/hello.proto
```

# Run
```shell
go run main.go
```

# Test
```shell
curl --location 'http://localhost:8081/v1/greet' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Shivaji"
}'
```

# Logs
look at `go run main.go` terminal output.