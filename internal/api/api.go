package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Request struct {
	Amount float64
	From   string
	To     string
}

type Response struct {
	Rate float64 `json:"conversion_rate"`
}

func FetchRates(apiKey string, params Request) (Response, error) {
	url := buildRequest(apiKey, params)

	resp, err := http.Get(url)

	if err != nil {
		return Response{}, fmt.Errorf("network error: %v", err)
	}

	if resp.StatusCode == http.StatusForbidden {
		return Response{}, fmt.Errorf("not valid api key")
	}

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("request error: %v", err)
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("response parsing error: %v", err)
	}

	return result, err
}

func buildRequest(key string, params Request) string {
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%v/pair/%v/%v", key, params.From, params.To)
	return url
}
