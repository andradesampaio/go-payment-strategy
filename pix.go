package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Pix struct {
	dados map[string]interface{}
}

func (p *Pix) Processar(pagamento Pagamento) ResultadoPagamento {
	resultado := ResultadoPagamento{
		PagamentoID:       uuid.NewString(),
		DataProcessamento: time.Now(),
		Detalhes:          make(map[string]string),
	}

	chavePix, ok := pagamento.Detalhes["chave_pix"]
	if !ok || chavePix == "" {
		resultado.Status = STATUS_RECUSADO
		resultado.Mensagem = "Chave Pix não informada ou inválida"
		return resultado
	}

	resultado.Status = STATUS_APROVADO
	resultado.Mensagem = "Pagamento Pix aprovado com sucesso"
	resultado.Detalhes["chave_pix"] = chavePix
	resultado.Detalhes["numero_comprovante"] = fmt.Sprintf("PIX-%d", time.Now().Unix())

	return resultado
}
