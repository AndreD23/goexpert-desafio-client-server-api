package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AndreD23/goexpert-desafio-client-server-api/quotation"
	"io"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", HelloServerHandler)
	http.HandleFunc("/cotacao", BuscaCotacaoHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloServerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, world!"}`))
}

func BuscaCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Busca cotacao
	cotacao, err := BuscaCotacao()
	if err != nil {
		fmt.Errorf("Error fetching USD-BRL: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)
}

func BuscaCotacao() (*quotation.QuotationResponse, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quotationResponse quotation.QuotationResponse
	err = json.Unmarshal(body, &quotationResponse)
	if err != nil {
		return nil, err
	}

	return &quotationResponse, nil
}
