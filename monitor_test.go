package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestRealizarRequisicao(t *testing.T) {
	servidorDeTeste := httptest.NewServer(http.HandlerFunc(func(resposta http.ResponseWriter, requisicao *http.Request) {
		resposta.WriteHeader(http.StatusOK)
	}))
	defer servidorDeTeste.Close()

	canalDeResultados := make(chan ResultadoVerificacao, 1)
	var grupoDeSincronizacao sync.WaitGroup

	grupoDeSincronizacao.Add(1)
	go realizarRequisicao(servidorDeTeste.URL, canalDeResultados, &grupoDeSincronizacao)

	grupoDeSincronizacao.Wait()
	close(canalDeResultados)

	resultado := <-canalDeResultados

	if resultado.Erro != nil {
		t.Errorf("Erro inesperado: %v", resultado.Erro)
	}

	if resultado.Status != http.StatusOK {
		t.Errorf("Esperado status %d, obtido %d", http.StatusOK, resultado.Status)
	}
}

func TestRealizarRequisicaoComErro(t *testing.T) {
	enderecoInvalido := "http://endereco.invalido.local"
	canalDeResultados := make(chan ResultadoVerificacao, 1)
	var grupoDeSincronizacao sync.WaitGroup

	grupoDeSincronizacao.Add(1)
	go realizarRequisicao(enderecoInvalido, canalDeResultados, &grupoDeSincronizacao)

	grupoDeSincronizacao.Wait()
	close(canalDeResultados)

	resultado := <-canalDeResultados

	if resultado.Erro == nil {
		t.Error("Esperava um erro para endereço inválido, mas o erro foi nil")
	}
}
