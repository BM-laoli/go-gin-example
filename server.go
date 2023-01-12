package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/BM-laoli/go-gin-example/proto"
)

type SearchService struct {
	pb.SearchServiceServer
}

// 直接抓进来 实现？集成？还是用的 实现 来说比较合适吧

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
