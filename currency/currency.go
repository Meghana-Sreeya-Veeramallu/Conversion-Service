package currency

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

func GetCurrencyType(currency string) CurrencyType {
	switch currency {
	case "USD":
		return USD
	case "EUR":
		return EUR
	case "GBP":
		return GBP
	case "JPY":
		return JPY
	case "INR":
		return INR
	default:
		return INR
	}
}
