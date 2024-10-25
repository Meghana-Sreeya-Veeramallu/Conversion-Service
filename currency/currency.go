package currency

import (
	"encoding/json"
	"fmt"
	"os"
)

type CurrencyType struct {
	ConversionFactor float64 `json:"conversion_factor"`
}

var currencyMap = make(map[string]CurrencyType)

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

	currencyMap = make(map[string]CurrencyType)
	for currency, factor := range config.Currencies {
		currencyMap[currency] = CurrencyType{ConversionFactor: factor}
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
	c, exists := currencyMap[currency]
	if !exists {
		return CurrencyType{}, fmt.Errorf("invalid currency: %s", currency)
	}
	return c, nil
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
