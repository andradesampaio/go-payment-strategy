package main

import "time"

type ResultadoPagamento struct {
	PagamentoID       string
	Status            StatusPagamento
	Mensagem          string
	DataProcessamento time.Time
	Detalhes          map[string]string
}
