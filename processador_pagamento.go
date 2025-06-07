package main

type ProcessadorPagamento interface {
	Processar(pagamento Pagamento) ResultadoPagamento
}
