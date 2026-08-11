# Testes com curl — API Glossário Tech

Testes manuais dos 5 endpoints da API REST em Go, rodando localmente em `http://localhost:8080`.

## GET /terms — listar termos (lista vazia)

```
curl -s -i http://localhost:8080/terms
```

**Resposta:**
```
HTTP/1.1 200 OK
Content-Type: application/json

[]
```

## POST /terms — criar termo

```
curl -s -i -X POST http://localhost:8080/terms \
  -H "Content-Type: application/json" \
  -d '{"termo":"API","categoria":"Backend","explicacao_simples":"Uma forma de sistemas conversarem entre si.","exemplo":"Um app de clima consulta uma API para buscar a previsão.","status":"estudando"}'
```

**Resposta:**
```
HTTP/1.1 201 Created
Content-Type: application/json

{"id":1,"termo":"API","categoria":"Backend","explicacao_simples":"Uma forma de sistemas conversarem entre si.","exemplo":"Um app de clima consulta uma API para buscar a previsão.","status":"estudando"}
```

## GET /terms/{id} — buscar termo existente

```
curl -s -i http://localhost:8080/terms/1
```

**Resposta:**
```
HTTP/1.1 200 OK
Content-Type: application/json

{"id":1,"termo":"API","categoria":"Backend","explicacao_simples":"Uma forma de sistemas conversarem entre si.","exemplo":"Um app de clima consulta uma API para buscar a previsão.","status":"estudando"}
```

## GET /terms/{id} — termo que não existe

```
curl -s -i http://localhost:8080/terms/999
```

**Resposta:**
```
HTTP/1.1 404 Not Found
Content-Type: application/json

{"erro":"termo não encontrado"}
```

## POST /terms — status inválido

```
curl -s -i -X POST http://localhost:8080/terms \
  -H "Content-Type: application/json" \
  -d '{"termo":"Docker","categoria":"Infra","explicacao_simples":"x","exemplo":"y","status":"invalido"}'
```

**Resposta:**
```
HTTP/1.1 400 Bad Request
Content-Type: application/json

{"erro":"status deve ser 'estudando' ou 'entendido'"}
```

## PUT /terms/{id} — atualizar termo

```
curl -s -i -X PUT http://localhost:8080/terms/1 \
  -H "Content-Type: application/json" \
  -d '{"termo":"API","categoria":"Backend","explicacao_simples":"Uma forma de sistemas conversarem entre si.","exemplo":"Um app de clima consulta uma API para buscar a previsão.","status":"entendido"}'
```

**Resposta:**
```
HTTP/1.1 200 OK
Content-Type: application/json

{"id":1,"termo":"API","categoria":"Backend","explicacao_simples":"Uma forma de sistemas conversarem entre si.","exemplo":"Um app de clima consulta uma API para buscar a previsão.","status":"entendido"}
```

## DELETE /terms/{id} — remover termo

```
curl -s -i -X DELETE http://localhost:8080/terms/1
```

**Resposta:**
```
HTTP/1.1 200 OK
Content-Type: application/json

{"mensagem":"termo removido"}
```

## GET /terms — lista após remoção

```
curl -s -i http://localhost:8080/terms
```

**Resposta:**
```
HTTP/1.1 200 OK
Content-Type: application/json

[]
```

## Resumo

| Endpoint testado | Status esperado | Resultado |
|---|---|---|
| `GET /terms` (vazio) | 200 | OK |
| `POST /terms` (criar) | 201 | OK |
| `GET /terms/{id}` (existente) | 200 | OK |
| `GET /terms/{id}` (inexistente) | 404 | OK |
| `POST /terms` (status inválido) | 400 | OK |
| `PUT /terms/{id}` | 200 | OK |
| `DELETE /terms/{id}` | 200 | OK |
| `GET /terms` (após remoção) | 200 | OK |
