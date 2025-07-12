package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

type exchageRatesMap = map[string]float64

func main() {
	exchangeRates := exchageRatesMap{
		"USDtoEUR": 0.87,
		"USDtoRUB": 78.29,
		"EURtoRUB": 91.06,
	}

	fmt.Println("*** Converter app ***")
	for {
		sourceCurrency, errorSourceCurrency := inputSourceCurrency()
		if errorSourceCurrency != nil {
			color.Red("Error! Enter the currency from the options above\n")
			continue
		}
		userInputData, errorUserInputData := getUserInput()
		if errorUserInputData != nil {
			color.Red("Error! The amount cannot be 0, less than 0, or contain letters\n")
			continue
		}
		targetCurrency, errorTargetCurrency := inputTargetCurrency(sourceCurrency)
		if errorTargetCurrency != nil {
			color.Red("Error! Enter the currency from the options above\n")
			continue
		}
		result := calculationData(&exchangeRates, userInputData, sourceCurrency, targetCurrency)
		resultOutput := fmt.Sprintf("Total: %.2f\n", result)
		color.Blue(resultOutput)
		repeatCalculation := repeatCalculation()
		if !repeatCalculation {
			break
		}
	}
}

func inputSourceCurrency() (string, error) {
	var sourceCurrency string
	fmt.Print("Enter the source currency (USD/EUR/RUB): ")
	fmt.Scan(&sourceCurrency)
	if sourceCurrency == "USD" || sourceCurrency == "EUR" || sourceCurrency == "RUB" {
		return sourceCurrency, nil
	}
	return "", errors.New("INVALID_DATA")
}

func getUserInput() (float64, error) {
	var userInputData float64
	fmt.Print("Enter the amount to convert: ")
	fmt.Scan(&userInputData)
	if userInputData <= 0.0 {
		return 0.0, errors.New("INVALID_DATA")
	} else {
		return userInputData, nil
	}
}

func inputTargetCurrency(sourceCurrency string) (string, error) {
	var targetCurrency string
	switch true {
	case sourceCurrency == "USD":
		fmt.Print("Enter the source currency (EUR/RUB): ")
		fmt.Scan(&targetCurrency)
	case sourceCurrency == "EUR":
		fmt.Print("Enter the source currency (USD/RUB): ")
		fmt.Scan(&targetCurrency)
	case sourceCurrency == "RUB":
		fmt.Print("Enter the source currency (USD/EUR): ")
		fmt.Scan(&targetCurrency)
	}
	if targetCurrency == "USD" || targetCurrency == "EUR" || targetCurrency == "RUB" {
		return targetCurrency, nil
	}
	return "", errors.New("INVALID_DATA")
}

func calculationData(exchangeRates *exchageRatesMap, userInputData float64, sourceCurrency string, targetCurrency string) float64 {
	var result float64
	switch {
	case sourceCurrency == "USD" && targetCurrency == "EUR":
		result = userInputData * (*exchangeRates)["USDtoEUR"]
	case sourceCurrency == "USD" && targetCurrency == "RUB":
		result = userInputData * (*exchangeRates)["USDtoRUB"]
	case sourceCurrency == "EUR" && targetCurrency == "USD":
		result = userInputData / (*exchangeRates)["USDtoEUR"]
	case sourceCurrency == "EUR" && targetCurrency == "RUB":
		result = userInputData * (*exchangeRates)["EURtoRUB"]
	case sourceCurrency == "RUB" && targetCurrency == "USD":
		result = userInputData / (*exchangeRates)["USDtoRUB"]
	case sourceCurrency == "RUB" && targetCurrency == "EUR":
		result = userInputData / (*exchangeRates)["EURtoRUB"]
	}
	return result
}

func repeatCalculation() bool {
	var userChoice string
	fmt.Print("Do you want to repeat the conversion? (Y/N): ")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	} else {
		return false
	}
}
