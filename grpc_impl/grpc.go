package grpc_impl

import (
    "flag"
    "fmt"
    pb "github.com/byrash/protos/gen/go/helloservice/v1"
    "google.golang.org/grpc"
    "log"
    "net"
)

var (
    grpcPort   = flag.Int("port", 50051, "The GRPC server port")
    grpcServer *grpc.Server
)

func StartGrpcServer() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer = grpc.NewServer()
    pb.RegisterHelloServiceServer(grpcServer, &HelloServiceImpl{})
    log.Printf("GRPC server listening at %v", lis.Addr())
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func StopGrpcServer() {
    grpcServer.Stop()
}
