package main

type ResultadoCobranca struct {
	CobrancaID string               `json:"cobranca_id"`
	Resultados []ResultadoPagamento `json:"resultados"`
}
