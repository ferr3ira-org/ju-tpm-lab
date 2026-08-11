# Revisão da Documentação — Desafio 04

## Quais endpoints foram implementados

- `GET /terms` — para listar todos os termos.
- `GET /terms/{id}` — para buscar um termo específico.
- `POST /terms` — para criar um termo.
- `PUT /terms/{id}` — para atualizar um termo.
- `DELETE /terms/{id}` — para deletar um termo.

## Os endpoints seguem o API Design?

Sim, seguem o `api-design.md` que revisamos antes. Os status codes esperados na documentação batem com os que foram implementados — por exemplo, `404` quando o termo não existe e `201` quando um termo é criado.

## O modelo de dados segue o Data Modeling?

Sim, os campos batem com os que definimos no `data-modeling.md`: os campos do termo (`termo`, `categoria`, `explicacao_simples`, `exemplo`, `status`) e os status permitidos (`estudando` ou `entendido`).

## Algum ponto da documentação precisou mudar?

Não precisou mudar. A implementação seguiu o que já estava planejado no PRD, API Design e Data Modeling sem nenhuma divergência.
