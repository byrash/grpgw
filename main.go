package main

import (
    "flag"
    "github.com/byrash/gateway_impl"
    "github.com/byrash/grpc_impl"
    "os"
    "os/signal"
)

func main() {
    flag.Parse()

    go grpc_impl.StartGrpcServer()
    go gateway_impl.StartGWServer()

    // Setting up signal capturing
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    // Waiting for SIGINT (kill -2)
    <-stop

    gateway_impl.StopGWServer()
    grpc_impl.StopGrpcServer()
}
