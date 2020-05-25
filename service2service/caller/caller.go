package main

import (
	"context"
	"fmt"
	"os"

	commonv1pb "github.com/dapr/go-sdk/dapr/proto/common/v1"
	pb "github.com/dapr/go-sdk/dapr/proto/dapr/v1"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"
)

func main() {
	// Get the Dapr port and create a connection
	daprPort := os.Getenv("DAPR_GRPC_PORT")
	daprAddress := fmt.Sprintf("localhost:%s", daprPort)
	fmt.Printf("Tryin to connect to dapr grpc with %s\n", daprAddress)
	conn, err := grpc.Dial(daprAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("after grpc.Dial(daprAddress, grpc.WithInsecure()) ")

	defer conn.Close()

	// Create the client
	client := pb.NewDaprClient(conn)
	fmt.Println("after pb.NewDaprClient(conn)")

	// Invoke a method called MyMethod on another Dapr enabled service with id client
	resp, err := client.InvokeService(context.Background(), &pb.InvokeServiceRequest{
		Id: "client",
		Message: &commonv1pb.InvokeRequest{
			Method:      "MyMethod",
			ContentType: "text/plain; charset=UTF-8",
			Data:        &any.Any{Value: []byte("Hello")},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("after client.InvokeService(context.Background(), ...)")

	if resp.GetContentType() != "text/plain; charset=UTF-8" {
		fmt.Printf("wrong content type: %s", resp.GetContentType())
	}

	fmt.Println(string(resp.GetData().GetValue()))
}
