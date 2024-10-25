package mocks

import (
	"context"
	"fmt"
	"testing"

	pb "ConversionService/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestConvertWithMock(t *testing.T) {
	mockServer := &MockCurrencyConverterServer{}

	mockServer.On("Convert", mock.Anything, &pb.ConvertRequest{
		FromCurrency: "USD",
		ToCurrency:   "EUR",
		Amount:       100,
	}).Return(&pb.ConvertResponse{ConvertedAmount: 91.0}, nil)

	resp, err := mockServer.Convert(context.Background(), &pb.ConvertRequest{
		FromCurrency: "USD",
		ToCurrency:   "EUR",
		Amount:       100,
	})

	assert.NoError(t, err)
	assert.Equal(t, 91.0, resp.ConvertedAmount)

	mockServer.AssertExpectations(t)
}

func TestConvertInvalidCurrencyWithMock(t *testing.T) {
	mockServer := &MockCurrencyConverterServer{}

	mockServer.On("Convert", mock.Anything, &pb.ConvertRequest{
		FromCurrency: "RAN",
		ToCurrency:   "USD",
		Amount:       100,
	}).Return(&pb.ConvertResponse{}, fmt.Errorf("invalid currency: RAN"))

	resp, err := mockServer.Convert(context.Background(), &pb.ConvertRequest{
		FromCurrency: "RAN",
		ToCurrency:   "USD",
		Amount:       100,
	})

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid currency: RAN")
	assert.Equal(t, &pb.ConvertResponse{}, resp)

	mockServer.AssertExpectations(t)
}

func TestConvertInvalidAmountWithMock(t *testing.T) {
	mockServer := &MockCurrencyConverterServer{}

	mockServer.On("Convert", mock.Anything, &pb.ConvertRequest{
		FromCurrency: "USD",
		ToCurrency:   "USD",
		Amount:       -100,
	}).Return(&pb.ConvertResponse{}, fmt.Errorf("invalid amount: %v. Must be greater than 0", -100))

	req := &pb.ConvertRequest{FromCurrency: "USD", ToCurrency: "USD", Amount: -100}
	resp, err := mockServer.Convert(context.Background(), req)

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid amount: -100. Must be greater than 0")
	assert.Equal(t, &pb.ConvertResponse{}, resp)

	mockServer.AssertExpectations(t)
}
