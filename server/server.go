package server

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "pc28/proto"
)

type SearchService struct {
	pb.UnimplementedSearchServiceServer
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server " + time.Now().Format(time.RFC3339Nano)}, nil
}

func Run() error {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		return err
	}

	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}
