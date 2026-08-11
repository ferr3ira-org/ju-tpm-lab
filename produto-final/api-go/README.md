# API Go — Glossário Tech para TPM

API REST em Go para cadastrar e consultar termos técnicos do Glossário Tech.

## Recurso principal

`terms`

## Endpoints esperados

| Método | Endpoint | Função |
|---|---|---|
| GET | `/terms` | Listar termos |
| GET | `/terms/{id}` | Buscar termo |
| POST | `/terms` | Criar termo |
| PUT | `/terms/{id}` | Atualizar termo |
| DELETE | `/terms/{id}` | Remover termo |

## Como rodar

### Pré-requisitos

- Go 1.22 ou superior instalado (`go version` para conferir)

### Rodando o servidor

```bash
cd produto-final/api-go
go run main.go
```

O servidor sobe em `http://localhost:8080`.

### Testando

Exemplo rápido com `curl`:

```bash
curl -i http://localhost:8080/terms
```

Mais exemplos de cada endpoint estão documentados em
[`desafios/04-api-rest-go/output/testes-com-curl.md`](../../desafios/04-api-rest-go/output/testes-com-curl.md).

### Observação sobre os dados

Os termos são armazenados **em memória** — ou seja, os dados somem quando o servidor é reiniciado. Não há banco de dados nesta etapa do desafio.
