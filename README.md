# goexpert-desafio-client-server-api

## Descrição
Este projeto é uma API cliente-servidor que busca a cotação do dólar (USD-BRL) e a salva em um banco de dados SQLite e em um arquivo de texto.

## Instalação

1. Clone o repositório
```bash
git clone https://github.com/AndreD23/goexpert-desafio-client-server-api.git
cd goexpert-desafio-client-server-api
```

2. Instale os módulos Go:
```bash
go mod tidy
```

## Execute o servidor

Entre na pasta do projeto e execute o servidor:
```bash
go run ./server.go
```

O servidor estará disponível em `http://localhost:8080`

## Execute o cliente

Abra um novo terminal, acesse a pasta do projeto e execute o cliente:
```bash
go run ./client.go
```

O cliente fará uma requisição ao servidor para buscar a cotação do dólar e salvará o resultado em um arquivo cotacao.txt.

## Verificação dos Dados

1. Abra o banco de dados SQLite
```bash
sqlite3 quotations.db
```

2. Liste as tabelas
```bash
.tables
```

3. Verifique os dados na tabela de cotações:
```bash
SELECT * FROM quotations;
```

4. Verifique o arquivo cotacao.txt
```bash
cat cotacao.txt
```

O arquivo conterá a cotação do dólar no formato:
```
Dólar: <valor>
```