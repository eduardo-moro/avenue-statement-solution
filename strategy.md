Este arquivo tem como objetivo documentar minha linha de raciocínio para resolver o problema proposto no projeto 'statement', apresentado no evento hands-on de Golang da comunidade Golang SP em parceria com a Avenue.



---



Para facilitar o processo de desenvolvimento, irei aplicar o [modelo c4](https://c4model.com/) sugerido pela colega de equipe Juli no dia do evento.



Minha idéia inicial é aplicar apenas a estrutura básica do sistema, apenas a titulo de poder rodar health checks.



Primeiramente, irei organizar os requisitos e fluxogramas necessários no projeto, não como forma de overengeniering, mas como forma de visualizar a solução com maior clareza.



Criei inicialmente um repositório git com os componentes que irei utilizar: `statement` para o repositório original com os testes (clonei ele como um submódulo do git, para evitar alterações), `steatement-answer` (talvez eu mude esse nome ainda...) para a solução ào problema, e o `statement-client` que irá se comunicar com a solução, para nos mostrar o extrato em tempo real, optei fazer uma solução TUI para manter o Golang em uso ao longo do projeto, e por que eu gosto de TUIs para visualização de dados :).



Não sou muito habituado à utilização do Make em projetos Go, mas percebi uma utilização ampla da ferramenta, em projetos como do kubernetes, docker e este projeto, então vou tentar utiliza-lo como oportunidade para adquirir esta experiência.

> objetivo: Desenvolver um sistema que gera extratos atualizados instantaneamente para os usuários, permitindo acompanhar movimentações em tempo real, para contas/moedas diferentes.

> os makefiles (até por obvio) são compatíveis com Linux, estou trabalhando no windows no momento, eu posso rodar wsl toda vez que for testar, ou posso utilizar um docker já recebendo os comandos, vale a pena criar um dockerfile só para isso? acho que vou manter o wsl por enquanto, se me irritar eu crio um dockerfile rápido.



Analisando o makefile, consigo entender um pouco de como é a estrutura que os testes esperam encontrar para poderem trabalhar em cima, vou criar a estrutura de pastas para o projeto, com o que já consigo visualizar.

> pedi para a IA gerar uma arvore dos arquivos neste momento: 



```
C:/Users/eduar/source/avenue-golang/
├── statement/
│   ├── Makefile
│   ├── README.md
│   └── script/
│       └── ...
└── solution/
    ├── go.mod
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
    │   │   └── statement\\\_service.go
    │   │
    │   └── repository/         # Camada de acesso a dados (banco de dados, cache)
    │       └── transaction\\\_repo.go
    │
    └── pkg/
        └── model/              # Modelos de dados (pode reutilizar a estrutura do validador)
            └── transaction.go
``` 

