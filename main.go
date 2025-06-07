package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	servico := NewServicoPagamento()

	http.HandleFunc("/cobranca", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
            return
        }

        var cobranca Cobranca
        if err := json.NewDecoder(r.Body).Decode(&cobranca); err != nil {
            http.Error(w, "JSON inválido", http.StatusBadRequest)
            return
        }

        var resultados []ResultadoPagamento

        resultado := servico.ProcessarPagamento(cobranca)
        resultados = append(resultados, resultado...)

        resposta := ResultadoCobranca{
            CobrancaID: cobranca.CobrancaID,
            Resultados: resultados,
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resposta)
    })

    fmt.Println("API rodando na URL: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}