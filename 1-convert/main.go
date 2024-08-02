package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const UsdToEur = 0.92305
const UsdToRub = 0.011346
const EurToRub = UsdToEur * UsdToRub

var currencies = [3]string{"USD", "EUR", "RUB"}

var currencyRate = map[string]map[string]float64{
	"USD": {"EUR": UsdToEur, "RUB": 1 / UsdToRub},
	"EUR": {"USD": 1 / UsdToEur, "RUB": 1 / EurToRub},
	"RUB": {"USD": UsdToRub, "EUR": EurToRub}}

func checkCurrency(c string) error {
	for _, currency := range currencies {
		if currency == c {
			return nil
		}
	}
	return errors.New("ОШИБКА! Валюта не найдена!")
}

func checkValue(v float64) error {
	if v <= 0 {
		return errors.New("ОШИБКА! Значение должно быть больше нуля!")
	}
	return nil
}

func getAvalibleCurrencies(currencyExclude string) string {
	var line string
	for _, currency := range currencies {
		if currency != currencyExclude {
			line = line + currency + ","
		}
	}
	return strings.TrimSuffix(line, ",")
}

func requestCurrency(message string) string {
	var currency string
	for {
		fmt.Printf("%s", message)
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)
		if err := checkCurrency(currency); err != nil {
			fmt.Println(err.Error())
		} else {
			return currency
		}
	}
}

func requestValue(message string) float64 {
	var requestValue string
	for {
		fmt.Printf("%s", message)
		fmt.Scan(&requestValue)

		value, err := strconv.ParseFloat(requestValue, 64)
		if err != nil {
			fmt.Println("ОШИБКА! Введенное значение не число!")
			continue
		}

		if err := checkValue(value); err != nil {
			fmt.Println(err.Error())
		} else {
			return value
		}
	}
}

func calc(value float64, sourceCurrency, targetCurrency string) float64 {
	return currencyRate[sourceCurrency][targetCurrency] * value
}

func main() {
	sourceCurrency := requestCurrency("Введите исходную валюту [" + getAvalibleCurrencies("") + "]: ")
	targetCurrency := requestCurrency("Введите целевую валюту [" + getAvalibleCurrencies(sourceCurrency) + "]: ")
	value := requestValue("Ведите количество единиц валюты: ")

	fmt.Println(calc(value, sourceCurrency, targetCurrency))
}
