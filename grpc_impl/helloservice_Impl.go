package grpc_impl

import (
    "context"
    pb "github.com/byrash/protos/gen/go/helloservice/v1"
    "log"
)

// HelloServiceImpl is used to implement helloworld.GreeterServer.
type HelloServiceImpl struct {
    pb.UnimplementedHelloServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *HelloServiceImpl) Greet(ctx context.Context, in *pb.GreetMsgReq) (*pb.GreetMsgResp, error) {
    log.Printf("GRPC Request Received: %v", in.GetName())
    return &pb.GreetMsgResp{Message: "Hello " + in.GetName()}, nil
}
