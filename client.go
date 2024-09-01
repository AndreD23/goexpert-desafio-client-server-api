package main

import (
	"encoding/json"
	"github.com/AndreD23/goexpert-desafio-client-server-api/quotation"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:8080/cotacao"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var quotationResponse quotation.QuotationResponse
	err = json.Unmarshal(body, &quotationResponse)
	if err != nil {
		panic(err)
	}

	println(quotationResponse.USD_BRL.Bid)
}
