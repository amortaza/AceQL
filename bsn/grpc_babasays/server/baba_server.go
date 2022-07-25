package main

import (
	"fmt"
	"github.com/amortaza/aceql/bsn/grpc_record"
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

	s := grpc_record.MyServer{}

	grpcServer := grpc.NewServer()

	grpc_record.RegisterRecordServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed %v", err)
	}
}
