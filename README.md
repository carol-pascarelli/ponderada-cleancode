# Ponderada - Clean Code

Projeto da ponderada desenvolvido por Laura e Carol com foco em separação de responsabilidades, injeção de dependência manual e organização em camadas.


## Visão geral

A API gerencia figurinhas e foi organizada em uma estrutura parecida com a discutida em aula:

```text
main.go
Domain/figurinha.go
Repository/figurinha_repo.go
Service/figurinha_service.go
Handler/figurinha_handler.go
```

Essa divisão ajuda a deixar cada parte do sistema com uma responsabilidade clara:

- `Domain`: entidade e tipos do negócio
- `Repository`: acesso ao banco com SQLite/GORM
- `Service`: regras de negócio e validações
- `Handler`: HTTP, requests, responses e tradução de erros para status code
- `main.go`: composição da aplicação, conexão com banco e registro de rotas

## Motivação das decisões

Escolhemos essa estrutura porque ela deixa o fluxo mais fácil de entender e de manter. Em vez de concentrar tudo em um único arquivo, cada camada passou a cuidar de uma parte específica do sistema. Isso facilita leitura, testes e evolução do código.

Também optamos por usar interfaces entre as camadas principais.  Assim, o service não depende de um banco específico, e o handler não precisa conhecer detalhes de persistência. Essa escolha ajuda muito em testes unitários, porque permite trocar o repository por uma implementação fake sem precisar de banco real.

A conexão com o SQLite ficou no `main.go` porque ele funciona como ponto de composição da aplicação. O `main` cria o banco, monta o repository, injeta no service e injeta o service no handler. Isso mantém o controle da aplicação em um único lugar e evita que as outras camadas saibam como o banco é aberto.

Usamos GORM e SQLite porque a proposta da atividade pedia um backend local com banco local e porque essa combinação reduz configuração. O SQLite cria um arquivo local e o GORM simplifica a criação da tabela com `AutoMigrate`, o que foi suficiente para o objetivo da ponderada.

Os erros de domínio foram nomeados para que o handler possa mapear cada caso para o status HTTP correto. Isso evita depender de texto solto de erro e deixa a tradução para HTTP mais confiável.


## Aprendizados em cada etapa

### Domain

Começamos o desenvolvimento pelo `Domain`. Definimos a struct `Figurinha` com os campos dados e tags GORM para persistência e declaramos os valores válidos de tipo e posição para a validação na camada de serviço.

### Repository

Em seguida, desenvolvemos o `Repository`, criando a interface de acesso ao banco e a implementação com SQLite/GORM. A interface permite trocar a tecnologia de persistência no futuro sem alterar o restante do código.


### Service

Depois, desenvolvemos o `Service`, onde ficaram as validações de campos obrigatórios, tipo válido e posição válida. Também centralizamos a criação, leitura, atualização e exclusão com regras de negócio antes de tocar no repository.


### Handler

Por último fizemos o `Handler`, que recebe requests HTTP, chama o service e converte os erros de domínio em respostas HTTP adequadas, como `400`, `404` e `500`.


### Main

Por fim, o `main.go` foi conectamos tudo: banco, repository, service, handler e rotas.

Aprendizado principal: o `main` funciona como composição da aplicação, não como lugar de regra de negócio.

## Como rodar localmente

### Pré-requisitos

- Go instalado na máquina
- Git instalado para clonar o repositório

O projeto usa as dependências declaradas no `go.mod`, principalmente:

- `github.com/gin-gonic/gin`
- `gorm.io/gorm`
- `gorm.io/driver/sqlite`

### Instalar dependências

Na raiz do projeto, rode:

```powershell
go mod download
```

Se quiser garantir que o arquivo de módulos fique coerente, você também pode usar:

```powershell
go mod tidy
```

### Subir o servidor

Na raiz do projeto:

```powershell
go run .
```

O servidor sobe em:

```text
http://localhost:8080
```

O SQLite cria o arquivo `figurinhas.db` na própria pasta do projeto, caso ele ainda não exista.


## Rotas da API

### GET /

Retorna uma mensagem simples para confirmar que a API está no ar.

Exemplo:

```powershell
Invoke-RestMethod "http://localhost:8080/"
```

Resposta esperada:

```json
{
	"message": "API de figurinhas funcionando",
	"routes": [
		"POST /figurinha",
		"GET /figurinha",
		"GET /figurinha/:id",
		"PUT /figurinha/:id",
		"DELETE /figurinha/:id"
	]
}
```

### POST /figurinha

Cria uma figurinha.

Exemplo:

```powershell
Invoke-RestMethod `
	-Method Post `
	-Uri "http://localhost:8080/figurinha" `
	-ContentType "application/json" `
	-Body '{"numero":"BRA 15","tipo":"comum","posicao":"atacante"}'
```

### GET /figurinha

Lista todas as figurinhas.

Exemplo:

```powershell
Invoke-RestMethod "http://localhost:8080/figurinha"
```

### GET /figurinha?posicao=...&tipo=...

Lista figurinhas com filtro.

Exemplo:

```powershell
Invoke-RestMethod "http://localhost:8080/figurinha?tipo=comum&posicao=atacante"
```

### GET /figurinha/:id

Busca uma figurinha pelo ID.

Exemplo:

```powershell
Invoke-RestMethod "http://localhost:8080/figurinha/1"
```

### PUT /figurinha/:id

Atualiza uma figurinha existente.

Exemplo:

```powershell
Invoke-RestMethod `
	-Method Put `
	-Uri "http://localhost:8080/figurinha/1" `
	-ContentType "application/json" `
	-Body '{"numero":"BRA 10","tipo":"brilhante","posicao":"atacante"}'
```

### DELETE /figurinha/:id

Remove uma figurinha existente.

Exemplo:

```powershell
Invoke-RestMethod -Method Delete "http://localhost:8080/figurinha/1"
```

## Resposta dos erros

O handler traduz os erros de domínio para respostas HTTP previsíveis:

- `404` quando a figurinha não é encontrada
- `400` quando os dados são inválidos ou obrigatórios estão faltando
- `500` para erros internos inesperados

## Estrutura do banco

O projeto usa SQLite e o `AutoMigrate` cria a tabela `figurinhas` com base na struct `Figurinha` do domínio.
