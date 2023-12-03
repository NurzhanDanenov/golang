package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "lecture18stream/protofiles/data_streaming"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Println("Error connecting to gRPC server: ", err.Error())
	}

	defer conn.Close()

	client := pb.NewStreamingServiceClient(conn)

	req := pb.DataRequest{Id: "123"}
	stream, err := client.GetDataStreaming(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return
		} else if err == nil {
			valStr := fmt.Sprintf("Response\n Part: %d \n Val: %s", resp.Part, resp.Buffer)
			log.Println(valStr)
		}

		if err != nil {
			panic(err)
		}

	}
}
