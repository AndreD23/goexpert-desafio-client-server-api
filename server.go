package main

import (
	"fmt"
	"net/http"
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
	w.Write([]byte(fmt.Sprintf(`{"cotacao": %f}`, cotacao)))
}

func BuscaCotacao() (float32, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)

	return 5.5, nil
}
