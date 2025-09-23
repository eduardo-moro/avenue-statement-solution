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

Pedi para o gemini gerar uma base para os arquivos, a maior parte do que ele trouxe foi ligar os arquivos aos que eles irão depender (nomear os packages e adicionar imports comentados para uso futuro)

após criar uma base de arquivos para processar a comunicação com o projeto `tatement`, notei que parte dos testes são apenas placeholders, testes como por exemplo:
``` 
// Verificação de consistência
func RunConsistencyCheck() {
	fmt.Println("\n🔍 Executando verificação de consistência...")
	fmt.Println("==================================================")

	checks := []string{
		"Verificação de saldos por usuário/conta/moeda",
		"Verificação de soma de transações",
		"Verificação de integridade referencial",
		"Verificação de timestamps",
		"Verificação de tipos de transação",
	}

	allPassed := true
	for i, check := range checks {
		time.Sleep(100 * time.Millisecond)
		passed := rand.Float64() > 0.1
		if !passed {
			allPassed = false
		}
		status := "✅"
		if !passed {
			status = "❌"
		}
		fmt.Printf("  %s %d. %s\n", status, i+1, check)
	}

	fmt.Println("\n🎯 Resultado da verificação de consistência:")
	if allPassed {
		fmt.Println("✅ Todas as verificações passaram - Sistema consistente")
	} else {
		fmt.Println("❌ Algumas verificações falharam - Investigar inconsistências")
	}
	fmt.Println("==================================================")
}
```
não oferecem realmente um teste de consistência dos dados recebidos.
Por enquanto, vamos manter estes testes, após ter um entendimento melhor das regras de negócio do projeto, eu volto re-implementando os testes com validações que façam mais sentido no contexto do projeto.


após a criação da base de arquivos, já temos um projeto que retorna os testes com sucesso: 
``` 
✅ 1000 transações enviadas com sucesso

🎯 Resumo da execução:
✅ Sucessos: 1000
❌ Erros: 0
📊 Taxa de sucesso: 100.00%

📈 Transações por tipo:
  CAMBIO: 369 transações
  WIRE: 181 transações
  ACAO: 101 transações
  TRANSACAO DE CARTAO: 174 transações
  TED: 92 transações
  PIX: 83 transações

🔄 Transações por direção:
  DEBITO: 516 transações
  CREDITO: 484 transações

🏦 Transações por conta:
  CONTA INVESTIMENTO: 215 transações
  CONTA BANKING: 530 transações
  CONTA BRASILEIRA: 255 transações

💱 Transações por moeda:
  EUR: 267 transações
  BRL: 255 transações
  USD: 478 transações

🔍 Executando verificação de consistência final...

🔍 Executando verificação de consistência...
==================================================
  ✅ 1. Verificação de saldos por usuário/conta/moeda
  ✅ 2. Verificação de soma de transações
  ✅ 3. Verificação de integridade referencial
  ✅ 4. Verificação de timestamps
  ✅ 5. Verificação de tipos de transação

🎯 Resultado da verificação de consistência:
✅ Todas as verificações passaram - Sistema consistente
==================================================

🎉 Validação do desafio concluída!
==================================================
``` 

Isso certamente está muito errado!
Não criamos nenhum banco de dados, nenhum cache, nenhuma validação, sanitização de dados, nenhum serviço!

Vou seguir criando o que o sistema se propõe a resolver, e conforme evoluo a solução, implemento os testes reais.