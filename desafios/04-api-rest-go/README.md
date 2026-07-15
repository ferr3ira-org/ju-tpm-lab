# Desafio 04 — API REST em Go

## Onde colocar sua entrega

Coloque **todos os arquivos deste desafio** em:

```text
desafios/04-api-rest-go/output/
```

Arquivos esperados:

- `desafios/04-api-rest-go/output/explicacao-api-rest.md`
- `desafios/04-api-rest-go/output/testes-com-curl.md`
- `desafios/04-api-rest-go/output/revisao-documentacao.md`

Não altere a pasta de outro desafio neste PR.


## Objetivo

Construir uma API REST simples em Go para o Glossário Tech para TPM.

Antes de implementar a API, leia e revise:

- [`docs/produto.md`](../../docs/produto.md)
- [`docs/prd.md`](../../docs/prd.md)
- [`docs/api-design.md`](../../docs/api-design.md)
- [`docs/data-modeling.md`](../../docs/data-modeling.md)

A implementação deve seguir essa documentação. Se você perceber que a documentação precisa mudar, explique o motivo no PR.

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
desafios/04-api-rest-go/output/revisao-documentacao.md
```

Em `revisao-documentacao.md`, explique:

- quais endpoints você implementou;
- se eles seguem o API Design;
- se o modelo de dados segue o Data Modeling;
- se algum ponto da documentação precisou mudar.

## Critérios de aceite

- [ ] Li PRD, API Design e Data Modeling antes de implementar.
- [ ] API roda localmente.
- [ ] Tem CRUD completo.
- [ ] Usa JSON.
- [ ] Retorna status codes adequados.
- [ ] README explica como rodar.
- [ ] Ju consegue explicar GET, POST, PUT e DELETE.


## Checklist antes de abrir PR

- [ ] Coloquei os arquivos na pasta `output/` deste desafio.
- [ ] Atualizei `docs/diario-de-estudos.md`.
- [ ] Sei explicar com minhas palavras o que fiz.
- [ ] Usei IA como ajuda, mas revisei o resultado.
- [ ] Abri PR pedindo review do Doug.
