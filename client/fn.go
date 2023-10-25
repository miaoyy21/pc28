package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "pc28/proto"
)

func Search() (*pb.SearchResponse, error) {
	// Create a client connection to the given target with a credentials which disables transport security
	conn, err := grpc.Dial(":9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}
