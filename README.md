# Go Gateway API

## Descrição

Este projeto é uma API Gateway escrita em Go. Ele fornece funcionalidades para gerenciar contas, incluindo criação de contas, consulta de informações e atualização de saldo. A API utiliza um banco de dados PostgreSQL para persistência e segue uma arquitetura modular com camadas de domínio, serviço, repositório e web.

## Estrutura do Projeto

```
go.mod
go.sum
README.md
cmd/
    app/
        main.go
internal/
    domain/
        account.go
        errors.go
        repository.go
    dto/
        account.go
    repository/
        account_repository.go
    service/
        account_service.go
    web/
        handlers/
            handlers.go
        server/
            server.go
```

## Pré-requisitos

- Go 1.24.1 ou superior
- PostgreSQL
- Arquivo `.env` configurado com as variáveis de ambiente necessárias:

```
HTTP_PORT=8080
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=gateway
DB_SSL_MODE=disable
```

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/gabrielalmir/go-gateway-api.git
cd go-gateway-api
```

2. Instale as dependências:

```bash
go mod tidy
```

3. Configure o banco de dados PostgreSQL e atualize o arquivo `.env` com as credenciais corretas.

## Execução

1. Inicie o servidor:

```bash
go run cmd/app/main.go
```

2. A API estará disponível em `http://localhost:8080` (ou na porta configurada no `.env`).

## Endpoints

### Criar Conta

**POST** `/api/v1/accounts`

- **Headers:**
  - `Content-Type: application/json`

- **Body:**

```json
{
  "name": "Nome do Usuário",
  "email": "email@exemplo.com"
}
```

- **Response:**

```json
{
  "id": "<ID>",
  "name": "Nome do Usuário",
  "email": "email@exemplo.com",
  "balance": 0,
  "api_key": "<API_KEY>",
  "created_at": "<DATA>",
  "updated_at": "<DATA>"
}
```

### Consultar Conta

**GET** `/api/v1/accounts`

- **Headers:**
  - `X-API-Key: <API_KEY>`

- **Response:**

```json
{
  "id": "<ID>",
  "name": "Nome do Usuário",
  "email": "email@exemplo.com",
  "balance": 0,
  "created_at": "<DATA>",
  "updated_at": "<DATA>"
}
```
