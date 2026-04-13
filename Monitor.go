package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type ResultadoVerificacao struct {
	Endereco string
	Status   int
	Erro     error
	Latencia time.Duration
}

func realizarRequisicao(endereco string, canal chan<- ResultadoVerificacao, grupo *sync.WaitGroup) {
	defer grupo.Done()

	if !strings.HasPrefix(endereco, "http") {
		endereco = "https://" + endereco
	}

	cliente := http.Client{
		Timeout: 10 * time.Second,
	}

	inicio := time.Now()
	resposta, erro := cliente.Get(endereco)
	duracao := time.Since(inicio)

	if erro != nil {
		canal <- ResultadoVerificacao{
			Endereco: endereco,
			Erro:     erro,
			Latencia: duracao,
		}
		return
	}
	defer resposta.Body.Close()

	canal <- ResultadoVerificacao{
		Endereco: endereco,
		Status:   resposta.StatusCode,
		Latencia: duracao,
	}
}

func main() {
	entradaUrls := flag.String("urls", "", "Lista de URLs")
	flag.Parse()

	if *entradaUrls == "" {
		fmt.Println("Uso: go run monitor.go -urls=google.com")
		os.Exit(1)
	}

	listaDeUrls := strings.Split(*entradaUrls, ",")
	canalResultados := make(chan ResultadoVerificacao, len(listaDeUrls))
	var grupoSincronizacao sync.WaitGroup

	for _, url := range listaDeUrls {
		grupoSincronizacao.Add(1)
		go realizarRequisicao(strings.TrimSpace(url), canalResultados, &grupoSincronizacao)
	}

	go func() {
		grupoSincronizacao.Wait()
		close(canalResultados)
	}()

	for item := range canalResultados {
		fmt.Printf("Resultado: %s | Status: %d\n", item.Endereco, item.Status)
	}
}
