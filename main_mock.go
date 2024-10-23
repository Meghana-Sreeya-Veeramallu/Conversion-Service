package main

import (
	pb "ConversionService/proto"
	"context"
	"github.com/stretchr/testify/mock"
)

// MockCurrencyConverterServer is a mock implementation of the ConversionServiceServer interface
type MockCurrencyConverterServer struct {
	mock.Mock
}

// Convert mocks the Convert method of the ConversionServiceServer interface
func (m *MockCurrencyConverterServer) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.ConvertResponse), args.Error(1)
}

// mustEmbedUnimplementedConversionServiceServer is a placeholder for the embedded unimplemented server
func (m *MockCurrencyConverterServer) mustEmbedUnimplementedConversionServiceServer() {}
