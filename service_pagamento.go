package main

import (
	"sync"
	"time"
)

type ServicoPagamento struct {
	processadores map[FormaPagamento]ProcessadorPagamento
}

func NewServicoPagamento() *ServicoPagamento {
	return &ServicoPagamento{
		processadores: map[FormaPagamento]ProcessadorPagamento{
			PIX:            &Pix{},
			CARTAO_CREDITO: &CartaoCredito{},
		},
	}
}

func (s *ServicoPagamento) ProcessarPagamento(cobranca Cobranca) (ResultadoPagamentos []ResultadoPagamento) {
	var wg sync.WaitGroup
	var resultChan = make(chan ResultadoPagamento, len(cobranca.Pagamento))
	for _, pagamento := range cobranca.Pagamento {
		wg.Add(1)
		go func(pgto Pagamento) {
			defer wg.Done()
			processador, ok := s.processadores[pgto.Forma]
			if !ok {
				resultChan <- ResultadoPagamento{
					PagamentoID:       "",
					Status:            STATUS_RECUSADO,
					Mensagem:          "Forma de pagamento não suportada",
					Detalhes:          nil,
					DataProcessamento: time.Now(),
				}
				return
			}
			resultado := processador.Processar(pgto)
			resultChan <- resultado
		}(pagamento)
	}

	wg.Wait()
	close(resultChan)
	var resultados []ResultadoPagamento

	for resultado := range resultChan {
		resultados = append(resultados, resultado)
	}

	return resultados
}
