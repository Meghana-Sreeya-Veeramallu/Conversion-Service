package currency

import "fmt"

type CurrencyType struct {
	ConversionFactor float64
}

var (
	USD = CurrencyType{ConversionFactor: 84.0}
	EUR = CurrencyType{ConversionFactor: 91.0}
	GBP = CurrencyType{ConversionFactor: 109.0}
	JPY = CurrencyType{ConversionFactor: 0.55}
	INR = CurrencyType{ConversionFactor: 1.0}
)

func (c CurrencyType) ToBase(value float64) float64 {
	return value * c.ConversionFactor
}

func (c CurrencyType) FromBase(value float64) float64 {
	return value / c.ConversionFactor
}

func GetCurrencyType(currency string) (CurrencyType, error) {
	switch currency {
	case "USD":
		return USD, nil
	case "EUR":
		return EUR, nil
	case "GBP":
		return GBP, nil
	case "JPY":
		return JPY, nil
	case "INR":
		return INR, nil
	default:
		return CurrencyType{}, fmt.Errorf("invalid currency: %s", currency)
	}
}
