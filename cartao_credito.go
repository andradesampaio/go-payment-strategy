package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CartaoCredito struct {
	dados map[string]interface{}
}

func (p *CartaoCredito) Processar(pagamento Pagamento) ResultadoPagamento {
	resultado := ResultadoPagamento{
		PagamentoID:       uuid.NewString(),
		DataProcessamento: time.Now(),
		Detalhes:          make(map[string]string),
	}

	numeroCartao, okNum := pagamento.Detalhes["numero_cartao"]
	validade, okVal := pagamento.Detalhes["validade"]
	cvv, okCvv := pagamento.Detalhes["cvv"]

	if !okNum || numeroCartao == "" || !okVal || validade == "" || !okCvv || cvv == "" {
		resultado.Status = STATUS_RECUSADO
		resultado.Mensagem = "Dados do cartão de crédito incompletos"
		return resultado
	}

	resultado.Status = STATUS_APROVADO
	resultado.Mensagem = "Pagamento com cartão de crédito aprovado com sucesso"
	resultado.Detalhes["numero_cartao"] = numeroCartao[len(numeroCartao)-4:]
	resultado.Detalhes["numero_autorizacao"] = fmt.Sprintf("CC-%d", time.Now().Unix())

	return resultado
}
