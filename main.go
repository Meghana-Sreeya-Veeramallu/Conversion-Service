package main

import (
	"context"
	"fmt"
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

	if request.Amount <= 0 {

		return nil, fmt.Errorf("invalid amount: %v. Must be greater than 0", request.Amount)
	}

	fromCurrency := currency.GetCurrencyType(request.FromCurrency)
	toCurrency := currency.GetCurrencyType(request.ToCurrency)

	baseValue := fromCurrency.ToBase(request.Amount)
	convertedAmount := toCurrency.FromBase(baseValue)
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
