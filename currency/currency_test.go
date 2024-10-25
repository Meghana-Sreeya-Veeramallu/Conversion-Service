package currency

import (
	"math"
	"testing"
)

func TestMain(m *testing.M) {
	if err := LoadCurrencies("currencies.json"); err != nil {
		panic(err)
	}

	m.Run()
}

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
		{currencyMap["USD"], 100, 8400},
		{currencyMap["EUR"], 100, 9100},
		{currencyMap["GBP"], 100, 10900},
		{currencyMap["JPY"], 100, 55},
		{currencyMap["INR"], 100, 100},
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
		{currencyMap["USD"], 840, 10},
		{currencyMap["EUR"], 910, 10},
		{currencyMap["GBP"], 1090, 10},
		{currencyMap["JPY"], 5.5, 10},
		{currencyMap["INR"], 10, 10},
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
		code      string
		expected  CurrencyType
		expectErr bool
	}{
		{"USD", currencyMap["USD"], false},
		{"EUR", currencyMap["EUR"], false},
		{"GBP", currencyMap["GBP"], false},
		{"JPY", currencyMap["JPY"], false},
		{"INR", currencyMap["INR"], false},
		{"XYZ", CurrencyType{}, true},
	}

	for _, test := range tests {
		result, err := GetCurrencyType(test.code)
		if test.expectErr {
			if err == nil {
				t.Errorf("GetCurrencyType(%v) = %v; want an error", test.code, result)
			}
		} else {
			if err != nil {
				t.Errorf("GetCurrencyType(%v) returned an error: %v; want %v", test.code, err, test.expected)
			} else if result != test.expected {
				t.Errorf("GetCurrencyType(%v) = %v; want %v", test.code, result, test.expected)
			}
		}
	}
}
