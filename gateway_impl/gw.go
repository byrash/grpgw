package gateway_impl

import (
    "context"
    "flag"
    gw "github.com/byrash/protos/gen/go/helloservice/v1"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "log"
    "net/http"
    "time"
)

var (
    // gRPC server endpoint
    grpcServerEndpoint = flag.String("grpc_impl-server-endpoint", "localhost:50051", "gRPC server endpoint")
    gwHttpServer       *http.Server
)

func StartGWServer() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()
    // Register gRPC server endpoint
    // Note: Make sure the gRPC server is running properly and accessible
    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
    err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
    if err != nil {
        log.Fatalf("Failed to start GW Server: %v", err)
    }
    // Start HTTP server (and proxy calls to gRPC server endpoint)
    log.Printf("Gateway server listening at %v", 8081)
    gwHttpServer = &http.Server{Addr: ":8081", Handler: mux}
    if err := gwHttpServer.ListenAndServe(); err != nil {
        //log.Fatalf("Failed to start GW Server: %v", err)
    }
}

func StopGWServer() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := gwHttpServer.Shutdown(ctx); err != nil {
        log.Fatalf("Failed to stop GW Server: %v", err)
    }
}
