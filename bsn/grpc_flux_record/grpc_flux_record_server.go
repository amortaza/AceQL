package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting Baba server...")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed %v", err)
	}

	s := grpc_flux_record.MyServer{}

	grpcServer := grpc.NewServer()

	grpc_flux_record.RegisterRecordServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed %v", err)
	}
}
