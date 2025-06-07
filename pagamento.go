package main

type Pagamento struct {
    Forma   FormaPagamento 
    Valor   float64        
    Detalhes map[string]string
}