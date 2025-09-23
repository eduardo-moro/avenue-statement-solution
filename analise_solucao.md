# Análise do Projeto e Estrutura da Solução

Este documento descreve a arquitetura de comunicação esperada entre o projeto de validação (`statement`) e a sua solução, além de sugerir uma estrutura de pastas para o novo projeto.

## 1. Resumo da Comunicação (API)

A sua solução deve funcionar como um **servidor web** que responderá a requisições HTTP enviadas pelo script de validação. O validador espera que o seu servidor esteja rodando em `http://localhost:8080`.

### Endpoints a serem implementados:

#### 1.1. `POST /events`
- **Propósito:** Receber e processar uma nova transação (evento).
- **Método:** `POST`
- **Corpo da Requisição (Request Body):** Um objeto JSON com a estrutura do `TransactionEvent` (detalhado na Seção 2).
- **Resposta Esperada (Response):** `HTTP 200 OK` em caso de sucesso.

#### 1.2. `GET /statement/{userId}/{accountType}/{currencyType}/{period}`
- **Propósito:** Retornar o extrato de um usuário para uma conta e moeda específica.
- **Método:** `GET`
- **Parâmetros na URL:**
    - `userId`: ID do usuário (ex: `user-123`).
    - `accountType`: Tipo da conta (ex: `CONTA BRASILEIRA`).
    - `currencyType`: Moeda (ex: `BRL`).
    - `period`: Período do extrato (o teste usa `30d`).
- **Resposta Esperada (Response):** `HTTP 200 OK` com o extrato do usuário no corpo da resposta (o formato do extrato fica a seu critério, desde que seja consistente).

#### 1.3. `GET /health`
- **Propósito:** Verificar se a aplicação está no ar. Usado pelo `Makefile` (`make check-api`).
- **Método:** `GET`
- **Resposta Esperada (Response):** `HTTP 200 OK`.

---

## 2. Estrutura de Dados de Comunicação

O validador enviará para o endpoint `POST /events` um JSON que corresponde à seguinte estrutura Go. Sua API deve ser capaz de decodificar este objeto.

### `TransactionEvent` (Payload do `POST /events`)

```go
// extraído de script/model/model.go

type TransactionMetadata struct {
	Description string `json:"description"`
	Source      string `json:"source,omitempty"`
	Reference   string `json:"reference,omitempty"`
}

type TransactionEvent struct {
	ID          string              `json:"id"`
	UserID      string              `json:"user_id"`
	Account     string              `json:"account"`      // ex: "CONTA BRASILEIRA"
	Currency    string              `json:"currency"`     // ex: "BRL"
	Type        string              `json:"type"`         // ex: "PIX"
	Direction   string              `json:"direction"`    // ex: "CREDITO"
	Amount      float64             `json:"amount"`
	Balance     float64             `json:"balance"`
	Metadata    TransactionMetadata `json:"metadata"`
	ProcessedAt time.Time           `json:"processed_at"`
	CreatedAt   time.Time           `json:"created_at"`
}
```

### Exemplo de JSON (Enviado pelo Validador)

```json
{
  "id": "test_pix_001",
  "user_id": "user-123",
  "account": "CONTA BRASILEIRA",
  "currency": "BRL",
  "type": "PIX",
  "direction": "CREDITO",
  "amount": 600.0,
  "balance": 600.0,
  "metadata": {
    "description": "Transferência recebida",
    "source": "manual-test",
    "reference": "PIX_REF_001"
  },
  "processed_at": "2024-01-01T10:00:00Z",
  "created_at": "2024-01-01T10:00:00Z"
}
```

---

## 3. Estrutura de Pastas Sugerida para a Solução

A solução deve ser criada em um **diretório separado**, no mesmo nível da pasta `statement`. Isso mantém o validador e a sua implementação desacoplados.

```
C:/Users/eduar/source/avenue-golang/
├── statement/                  # <-- Projeto validador (existente)
│   ├── Makefile
│   ├── README.md
│   └── script/
│       └── ...
└── solution/                   # <-- SEU PROJETO DEVE SER CRIADO AQUI
    ├── go.mod                  # Arquivo de módulos do seu projeto
    ├── go.sum
    │
    ├── cmd/
    │   └── server/
    │       └── main.go         # Ponto de entrada do seu servidor HTTP
    │
    ├── internal/
    │   ├── api/                # Handlers da API (ex: para /events, /statement)
    │   │   └── handlers.go
    │   │   └── routes.go
    │   │
    │   ├── domain/             # Lógica de negócio e regras de domínio
    │   │   └── statement_service.go
    │   │
    │   └── repository/         # Camada de acesso a dados (banco de dados, cache)
    │       └── transaction_repo.go
    │
    └── pkg/
        └── model/              # Modelos de dados (pode reutilizar a estrutura do validador)
            └── transaction.go

```
