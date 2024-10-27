package currency

import (
	"encoding/json"
	"fmt"
	"os"
)

type CurrencyType struct {
	ConversionFactor float64 `json:"conversion_factor"`
}

var (
	USD CurrencyType
	EUR CurrencyType
	GBP CurrencyType
	JPY CurrencyType
	INR CurrencyType
)

func LoadCurrencies(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var config struct {
		Currencies map[string]float64 `json:"currencies"`
	}
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return err
	}

	for currency, factor := range config.Currencies {
		switch currency {
		case "USD":
			USD = CurrencyType{ConversionFactor: factor}
		case "EUR":
			EUR = CurrencyType{ConversionFactor: factor}
		case "GBP":
			GBP = CurrencyType{ConversionFactor: factor}
		case "JPY":
			JPY = CurrencyType{ConversionFactor: factor}
		case "INR":
			INR = CurrencyType{ConversionFactor: factor}
		}
	}

	return nil
}

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

func ConvertCurrency(fromCurrency string, toCurrency string, amount float64) (float64, error) {
	if err := LoadCurrencies("../currency/currencies.json"); err != nil {
		return 0, fmt.Errorf("failed to load currencies: %v", err)
	}

	if amount <= 0 {
		return 0, fmt.Errorf("invalid amount: %v. Must be greater than 0", amount)
	}

	from, err := GetCurrencyType(fromCurrency)
	if err != nil {
		return 0, err
	}

	to, err := GetCurrencyType(toCurrency)
	if err != nil {
		return 0, err
	}

	baseValue := from.ToBase(amount)
	convertedAmount := to.FromBase(baseValue)
	return convertedAmount, nil
}
