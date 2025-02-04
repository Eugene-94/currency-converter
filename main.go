package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Eugene-94/currency-converter/internal/api"
)

func main() {
	apiKey := os.Getenv("EXCHANGE_RATE_API_KEY")

	var amount = flag.Float64("amount", 1.0, "how much currency need to convert")
	var from = flag.String("from", "", "")
	var to = flag.String("to", "", "")
	flag.Parse()

	validateArgs(amount, from, to)

	params := api.Request{Amount: *amount, From: *from, To: *to}

	fetched, err := api.FetchRates(apiKey, params)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Результат: %.2f", convent(fetched.Rate, params.Amount))
}

func validateArgs(amount *float64, from *string, to *string) {
	if *amount <= 0 {
		log.Fatal("Ошибка: сумма должна быть больше 0")
	}

	if *from == "" || *to == "" {
		log.Fatal("Не указана валюта")
	}

	if len(*from) != 3 || len(*to) != 3 {
		log.Fatal("Код валюты должен состоять из 3 символов")
	}
}

func convent(rate, amount float64) float64 {
	return rate * amount
}
