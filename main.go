package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type request struct {
	amount float64
	from   string
	to     string
}

type response struct {
	Rate float64 `json:"conversion_rate"`
}

func main() {
	apiKey := os.Getenv("EXCHANGE_RATE_API_KEY")

	var amount = flag.Float64("amount", 1.0, "how much currency need to convert")
	var from = flag.String("from", "", "")
	var to = flag.String("to", "", "")
	flag.Parse()

	validateArgs(amount, from, to)

	params := request{amount: *amount, from: *from, to: *to}
	url := buildRequest(apiKey, params)

	resp, _ := http.Get(url)

	body, _ := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	var result response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	fmt.Printf("Результат: %.2f", result.Rate*params.amount)
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

func buildRequest(key string, params request) string {
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%v/pair/%v/%v", key, params.from, params.to)
	return url
}
