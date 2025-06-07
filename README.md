# 💳 Pagamento simultâneo com Go, Strategy Pattern Goroutines e Channels

## 📘 Cenário

A sua aplicação deve gerenciar **múltiplas formas de pagamento simultâneas**, como:

- Cartão de crédito(pode existir mais de um)
- PIX
- (E no futuro: Bitcoins, pontos de milhas, etc)

Desafio é:

- Realizar o processamento de todas essas transações **simultaneamente**
- **Separar** a lógica para cada forma de pagamento
- Ter **um controle centralizado dos resultados de cada pagamento**
- Retornar o resultado de cada transação, permitindo implementar retry, logs e auditoria.

---

## 🎯 Solução Técnica

Este projeto utiliza:

- **✅ Strategy Pattern**  
  Simplifica a implementanção de novas formas de pagamento sem modificar a lógica central.
  Isso deixa a aplicação aberta para extensão e fechada para modificação conceito do (Open/Closed Principle ou OCP do SOLID).
  Por exemplo, seria fácil adicionar forma de pagamento Bitcoin ou pontos de milhas.

- **⚙️ Goroutines**  
  Cada forma de pagamento é processado em simultaneamente, aumentando a performance e eficiência.

- **📬 Channels + `sync.WaitGroup`**  
  Para sincronizar e capturar os resultados das execuções concorrentes com segurança e previsibilidade.

## 🚀 Como usar

Inicie a API:

```sh
go run .
```

## Exemplo de requisição

```sh
curl -X POST http://localhost:8080/cobranca \
  -H "Content-Type: application/json" \
  -d '{
    "cobranca_id": "cob0002",
    "pagamentos": [
      {
        "forma": "CARTAO_CREDITO",
        "valor": 100,
        "detalhes": {
          "numero_cartao": "1234567890123456",
          "validade": "12/27",
          "cvv": "123"
        }
      },
      {
        "forma": "PIX",
        "valor": 200,
        "detalhes": {
          "chave_pix": "pix@bacen.com"
        }
      },
      {
        "forma": "boleto",
        "valor": 100,
        "detalhes": {
          "codigo_barras": "3411234567890"
        }
      }
    ]
  }'
```