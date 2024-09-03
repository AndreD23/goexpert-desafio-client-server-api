package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/AndreD23/goexpert-desafio-client-server-api/quotation"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	cotacao, err := BuscaCotacaoServer()
	if err != nil {
		log.Fatalf("Error fetching USD-BRL: %v", err)
	}

	SalvarArquivoCotacao(cotacao)
}

func BuscaCotacaoServer() (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	url := "http://localhost:8080/cotacao"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("Erro ao buscar cotação: timeout do contexto excedido")
		}
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var quotationResponse quotation.QuotationResponse
	err = json.Unmarshal(body, &quotationResponse)
	if err != nil {
		panic(err)
	}

	return quotationResponse.USD_BRL.Bid, nil
}

func SalvarArquivoCotacao(cotacao string) {
	// Salvar arquivo
	f, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	str := "Dólar: " + cotacao
	_, err = f.WriteString(str)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}
