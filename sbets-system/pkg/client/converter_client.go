package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ConverterClient struct {
	baseURL string
}

type ConvertRequest struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

type ConvertResponse struct {
	ConvertedAmount float64 `json:"convertedAmount"`
	OriginalAmount  float64 `json:"originalAmount"`
	FromCurrency    string  `json:"fromCurrency"`
	ToCurrency      string  `json:"toCurrency"`
	ExchangeRate    float64 `json:"exchangeRate"`
}

func NewConverterClient() *ConverterClient {
	baseURL := os.Getenv("CONVERTER_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8085"
	}
	return &ConverterClient{baseURL: baseURL}
}

func (c *ConverterClient) Convert(amount float64, from, to string) (*ConvertResponse, error) {
	req := ConvertRequest{
		Amount: amount,
		From:   from,
		To:     to,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.baseURL+"/convert", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to call converter service: %v", err)
	}
	// When Convert() finishes, close the response body.
	defer resp.Body.Close()

	print("RESPONSE FROM CONVERTER:  ", resp)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("converter service returned status %d", resp.StatusCode)
	}

	var result ConvertResponse
	// resp.Body is an open stream (important!)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
