package currency

import (
	"math"
	"testing"
)

const tolerance = 1e-9

func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < tolerance
}

func TestToBase(t *testing.T) {
	tests := []struct {
		currency CurrencyType
		value    float64
		expected float64
	}{
		{USD, 100, 8400},
		{EUR, 100, 9100},
		{GBP, 100, 10900},
		{JPY, 100, 55},
		{INR, 100, 100},
	}

	for _, test := range tests {
		result := test.currency.ToBase(test.value)
		if !floatEqual(result, test.expected) {
			t.Errorf("ToBase(%v) = %v; want %v", test.value, result, test.expected)
		}
	}
}

func TestFromBase(t *testing.T) {
	tests := []struct {
		currency CurrencyType
		value    float64
		expected float64
	}{
		{USD, 840, 10},
		{EUR, 910, 10},
		{GBP, 1090, 10},
		{JPY, 5.5, 10},
		{INR, 10, 10},
	}

	for _, test := range tests {
		result := test.currency.FromBase(test.value)
		if !floatEqual(result, test.expected) {
			t.Errorf("FromBase(%v) = %v; want %v", test.value, result, test.expected)
		}
	}
}

func TestGetCurrencyType(t *testing.T) {
	tests := []struct {
		code     string
		expected CurrencyType
	}{
		{"USD", USD},
		{"EUR", EUR},
		{"GBP", GBP},
		{"JPY", JPY},
		{"INR", INR},
		{"XYZ", INR},
	}

	for _, test := range tests {
		result := GetCurrencyType(test.code)
		if result != test.expected {
			t.Errorf("GetCurrencyType(%v) = %v; want %v", test.code, result, test.expected)
		}
	}
}
