package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Quotation struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type QuotationResponse struct {
	USD_BRL Quotation `json:"USDBRL"`
}

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
	w.Write([]byte(fmt.Sprintf(`{"cotacao": %s}`, cotacao.USD_BRL.Bid)))
}

func BuscaCotacao() (*QuotationResponse, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quotationResponse QuotationResponse
	err = json.Unmarshal(body, &quotationResponse)
	if err != nil {
		return nil, err
	}

	return &quotationResponse, nil
}
