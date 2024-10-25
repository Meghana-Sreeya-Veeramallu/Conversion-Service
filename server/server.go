package main

import (
	"context"
	"log"
	"net"

	"ConversionService/currency"
	pb "ConversionService/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedConversionServiceServer
}

func (s *server) Convert(_ context.Context, request *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	convertedAmount, err := currency.ConvertCurrency(request.FromCurrency, request.ToCurrency, request.Amount)
	if err != nil {
		return nil, err
	}

	return &pb.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterConversionServiceServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
