package main

import (
	"fmt"
	"github.com/amortaza/aceql/bsn/grpc_record"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure()) // :9000

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := grpc_record.NewRecordServiceClient(conn) // Connect chat service

	runit(c)

	fmt.Println("bye")
}

func runit(client grpc_record.RecordServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.BabaSays(ctx)
	if err != nil {
		log.Fatalf("client.RouteChat failed: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			fmt.Println("client loop")
			in, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("got EOF")
				close(waitc)
				return
			}
			if err != nil {
				fmt.Println("failed: %v", err)
			}

			fmt.Println("Got Answer %s", in.Answer)
		}
	}()

	go func() {
		request := grpc_record.Request{Param: "p1", Operation: "operation cool cat"}
		if err := stream.Send(&request); err != nil {
			log.Fatalf("Server failed failed: %v", err)
		}

		time.Sleep(1000 * time.Millisecond)

		request = grpc_record.Request{Param: "p1b", Operation: "operation cool cat"}
		if err := stream.Send(&request); err != nil {
			log.Fatalf("Server failed failed: %v", err)
		}

		stream.CloseSend()
	}()

	//go func() {
	//	request := grpc_record.Request{Param: "p2", Operation: "operation cool cat"}
	//	if err := stream.Send(&request); err != nil {
	//		log.Fatalf("Server failed failed: %v", err)
	//	}
	//}()

	fmt.Println("waiting")
	<-waitc
	fmt.Println("done waiting, closing stream")
	stream.CloseSend()
}
