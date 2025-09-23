# Resumo da Solução para o Sistema de Extrato

Este documento descreve a estrutura de pastas e a comunicação esperada para a solução do desafio do sistema de extrato.

## Estrutura de Pastas Esperada

O projeto da solução deve ser criado em um diretório `solution` no mesmo nível da pasta `statement`, conforme a estrutura abaixo:

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

## Comunicação entre Projetos

A comunicação entre o projeto de validação (`statement`) e a sua solução será feita através de uma API REST. Sua solução deve atuar como um servidor HTTP rodando em `http://localhost:8080`.

### Endpoints a serem Implementados

| Método | Endpoint                                                       |Descrição                                                                     |
| :----- | :-----------------------------------------------------------   |:------------------------------------------------------------------------     |
| `POST` | `/events`                                                      | Recebe e processa uma nova transação (evento).                               |
| `GET`  | `/statement/{userId}/{accountType}/{currencyType}/{period}`    | Retorna o extrato de um usuário para uma conta, moeda e período específicos. |
| `GET`  | `/health`                                                      | Verifica se a aplicação está no ar.                                          |

### Estrutura de Dados (`POST /events`)

O endpoint `POST /events` receberá um JSON com a seguinte estrutura:

```json
{
  "id": "string",
  "user_id": "string",
  "account": "string",
  "currency": "string",
  "type": "string",
  "direction": "string",
  "amount": "float64",
  "balance": "float64",
  "metadata": {
    "description": "string",
    "source": "string",
    "reference": "string"
  },
  "processed_at": "time.Time",
  "created_at": "time.Time"
}
```


Pedi para o gemini gerar uma base para os arquivos, a maior parte do que ele trouxe foi ligar os arquivos aos que eles irão depender (nomear os packages e adicionar imports comentados para uso futuro)