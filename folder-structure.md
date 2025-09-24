```
solution/
│── docker-compose.yml          # Orquestração de todos os serviços
│── .env                        # Variáveis de ambiente globais
│
├── services/
│   ├── api-write/              # API de comandos (transações)
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── handlers/       # Handlers HTTP (POST /transaction)
│   │   │   ├── services/       # Lógica de negócio
│   │   │   ├── repository/     # Persistência MongoDB
│   │   │   └── events/         # Publicação no RabbitMQ
│   │   ├── go.mod
│   │   └── Dockerfile
│   │
│   ├── api-read/               # API de consultas (extrato)
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── handlers/       # Handlers HTTP (GET /statement)
│   │   │   ├── services/       # Lógica de leitura
│   │   │   ├── repository/     # Persistência PostgreSQL
│   │   │   └── consumer/       # Consome eventos do RabbitMQ
│   │   ├── go.mod
│   │   └── Dockerfile
│   │
│   └── client-tui/             # Cliente TUI em Go
│       ├── cmd/
│       │   └── main.go
│       ├── internal/
│       │   ├── ui/             # Interface TUI (bubbletea)
│       │   └── api/            # Cliente HTTP para chamar APIs (/write, /read)
│       ├── go.mod
│       └── Dockerfile
│
├── pkg/                        # Código compartilhado
│   ├── models/                 # Structs: User, Account, TransactionEvent, Metadata
│   ├── events/                 # Definição de eventos (TransacaoRegistrada)
│   └── utils/                  # Funções utilitárias (UUID, logger, etc.)
│
├── infra/                      # Infraestrutura
│   ├── mongo/                  # DBWrite
│   │   └── Dockerfile (opcional)
│   ├── postgres/               # DBRead
│   │   └── init.sql             # DDL inicial
│   ├── rabbitmq/               # Mensageria
│   └── nginx/                  # Proxy reverso (API Gateway)
│       └── nginx.conf           # Rotas /write → api-write, /read → api-read
```