package main

type Cobranca struct {
	CobrancaID string `json:"cobranca_id"`
	Pagamento  []Pagamento `json:"pagamentos"`
}
