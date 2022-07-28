package main

import (
	"fmt"
	"github.com/amortaza/aceql/bsn/grpc_flux_record"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting GRPC Flux Record server...")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed %v", err)
	}

	s := grpc_flux_record.FluxRecordServiceImp{}

	grpcServer := grpc.NewServer()

	grpc_flux_record.RegisterFluxRecordServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed %v", err)
	}
}
