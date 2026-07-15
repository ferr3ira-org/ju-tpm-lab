# API Design

**API Design** é o planejamento da API antes de escrever código.

Antes de implementar, você vai desenhar:

- endpoints;
- métodos HTTP;
- request body;
- response body;
- status codes.

## Recurso principal

```text
terms
```

Um `term` representa um termo técnico do glossário.

## Modelo inicial de um termo

```json
{
  "id": 1,
  "termo": "API",
  "categoria": "Backend",
  "explicacao_simples": "Uma forma de sistemas conversarem entre si.",
  "exemplo": "Um app de clima consulta uma API para buscar a previsão.",
  "status": "estudando"
}
```

## Endpoints planejados

| Método | Endpoint | Objetivo |
|---|---|---|
| `GET` | `/terms` | Listar todos os termos. |
| `GET` | `/terms/{id}` | Buscar um termo específico. |
| `POST` | `/terms` | Criar um termo. |
| `PUT` | `/terms/{id}` | Atualizar um termo. |
| `DELETE` | `/terms/{id}` | Remover um termo. |

## Exemplo de request — criar termo

```http
POST /terms
Content-Type: application/json
```

```json
{
  "termo": "API",
  "categoria": "Backend",
  "explicacao_simples": "Uma forma de sistemas conversarem entre si.",
  "exemplo": "Um app de clima consulta uma API para buscar a previsão.",
  "status": "estudando"
}
```

## Exemplo de response — criado com sucesso

```http
201 Created
```

```json
{
  "id": 1,
  "termo": "API",
  "categoria": "Backend",
  "explicacao_simples": "Uma forma de sistemas conversarem entre si.",
  "exemplo": "Um app de clima consulta uma API para buscar a previsão.",
  "status": "estudando"
}
```

## Status codes esperados

| Status | Quando usar |
|---|---|
| `200 OK` | Busca/listagem/edição com sucesso. |
| `201 Created` | Termo criado com sucesso. |
| `400 Bad Request` | Dados inválidos. |
| `404 Not Found` | Termo não encontrado. |
| `500 Internal Server Error` | Erro inesperado. |

## Uso da IA

Você pode pedir para a IA revisar se os endpoints fazem sentido, mas precisa entender cada método e cada response.
