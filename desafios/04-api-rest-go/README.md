# Desafio 04 — API REST em Go

## Objetivo

Construir uma API REST simples em Go para o Glossário Tech para TPM.

## Conceitos

Você precisa entender e explicar:

- O que é uma API
- O que é REST
- O que é endpoint
- O que é JSON
- O que são métodos HTTP
- O que são status codes
- Diferença entre frontend e backend

## CRUD esperado

Recurso: `terms`

| Método | Endpoint | Função |
|---|---|---|
| GET | `/terms` | Listar termos |
| GET | `/terms/{id}` | Buscar termo |
| POST | `/terms` | Criar termo |
| PUT | `/terms/{id}` | Atualizar termo |
| DELETE | `/terms/{id}` | Remover termo |

## Estrutura sugerida

```text
produto-final/api-go/
  go.mod
  main.go
  README.md
```

## Entregáveis

```text
produto-final/api-go/
desafios/04-api-rest-go/output/explicacao-api-rest.md
desafios/04-api-rest-go/output/testes-com-curl.md
```

## Critérios de aceite

- [ ] API roda localmente.
- [ ] Tem CRUD completo.
- [ ] Usa JSON.
- [ ] Retorna status codes adequados.
- [ ] README explica como rodar.
- [ ] Ju consegue explicar GET, POST, PUT e DELETE.
