package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/AndreD23/goexpert-desafio-client-server-api/quotation"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	url := "http://localhost:8080/cotacao"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("Erro ao buscar cotação: timeout do contexto excedido")
		}
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
