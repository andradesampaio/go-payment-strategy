package main

type StatusPagamento string

const (
    STATUS_APROVADO   StatusPagamento = "APROVADO"
    STATUS_RECUSADO   StatusPagamento = "RECUSADO"
    STATUS_PENDENTE   StatusPagamento = "PENDENTE"
    STATUS_PROCESSANDO StatusPagamento = "PROCESSANDO"
)