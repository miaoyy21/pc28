package server

import (
	"context"
	"time"

	pb "pc28/proto"
)

type SearchService struct {
	pb.UnimplementedSearchServiceServer
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server " + time.Now().Format(time.RFC3339Nano)}, nil
}

func Run(targetGold, targetBetting string) {
	go gGold(targetGold)

	gBetting(targetBetting)
}
