package main

type FormaPagamento string

const (
	PIX            FormaPagamento = "PIX"
	CARTAO_CREDITO FormaPagamento = "CARTAO_CREDITO"
	CRIPTOMOEDA    FormaPagamento = "CRIPTOMOEDA"
)
