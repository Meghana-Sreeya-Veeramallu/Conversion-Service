package main

import (
	"context"
	"google.golang.org/grpc/status"
	"log"
	"math"
	"testing"

	pb "ConversionService/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
)

const tolerance = 1e-2

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < tolerance
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterConversionServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestConvert(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewConversionServiceClient(conn)

	tests := []struct {
		fromCurrency string
		toCurrency   string
		amount       float64
		expected     float64
	}{
		{"USD", "USD", 100, 100},
		{"INR", "USD", 8400, 100},
		{"USD", "INR", 100, 8400},
		{"INR", "INR", 100, 100},
		{"EUR", "USD", 100, 108.33},
		{"GBP", "USD", 100, 129.76},
		{"JPY", "INR", 10000, 5500.00},
		{"USD", "GBP", 130, 100.18},
		{"INR", "EUR", 8500, 93.40},
		{"GBP", "JPY", 100, 19818.18},
	}

	for _, test := range tests {
		req := &pb.ConvertRequest{
			FromCurrency: test.fromCurrency,
			ToCurrency:   test.toCurrency,
			Amount:       test.amount,
		}
		resp, err := client.Convert(ctx, req)
		if err != nil {
			t.Fatalf("Convert(%v) failed: %v", req, err)
		}
		if !floatEqual(resp.ConvertedAmount, test.expected) {
			t.Errorf("Convert(%v) = %v; want %v", req, resp.ConvertedAmount, test.expected)
		}
	}
}

func TestConvertInvalidCurrencies(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewConversionServiceClient(conn)

	tests := []struct {
		fromCurrency  string
		toCurrency    string
		amount        float64
		expectedError string
	}{
		{"RAN", "INR", 100, "invalid currency: RAN"},
		{"INR", "RAN", 100, "invalid currency: RAN"},
		{"RAN", "RAN", 100, "invalid currency: RAN"},
		{"USD", "RAN", 100, "invalid currency: RAN"},
		{"RAN", "USD", 8400, "invalid currency: RAN"},
		{"EUR", "RAN", 100, "invalid currency: RAN"},
		{"RAN", "EUR", 100, "invalid currency: RAN"},
	}

	for _, test := range tests {
		req := &pb.ConvertRequest{
			FromCurrency: test.fromCurrency,
			ToCurrency:   test.toCurrency,
			Amount:       test.amount,
		}
		_, err := client.Convert(ctx, req)
		if err == nil {
			t.Errorf("Convert(%v) = nil; want error", req)
		} else {
			if status, ok := status.FromError(err); ok {
				if status.Message() != test.expectedError {
					t.Errorf("Convert(%v) = unexpected error: %v; want %v", req, status.Message(), test.expectedError)
				}
			} else {
				t.Errorf("Convert(%v) = unexpected error: %v; want %v", req, err, test.expectedError)
			}
		}
	}
}

func TestConvertInvalidAmountException(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewConversionServiceClient(conn)

	tests := []struct {
		fromCurrency string
		toCurrency   string
		amount       float64
	}{
		{"USD", "USD", 0},
		{"USD", "USD", -100},
	}

	for _, test := range tests {
		req := &pb.ConvertRequest{
			FromCurrency: test.fromCurrency,
			ToCurrency:   test.toCurrency,
			Amount:       test.amount,
		}
		_, err := client.Convert(ctx, req)
		if err == nil {
			t.Errorf("Expected error for Convert(%v), but got none", req)
		}
	}
}
