# Plano de Release — Glossário Tech para TPM

## O que está sendo entregue (escopo)
Esta release cobre o MVP do Glossário Tech:
- Cadastrar novo termo (com categoria, explicação simples e exemplo)
- Listar termos cadastrados (com categoria)
- Editar termo existente
- Remover termo
- Marcar status ("estudando" ou "entendido")

## Como a entrega é feita
1. Buildar e subir os containers com `docker compose up --build`.
2. Testar no navegador para garantir que tudo está funcionando (API + front rodando, termo sendo criado/listado com sucesso).

## Critérios de aceite
Cada funcionalidade do MVP tem seus próprios critérios de sucesso e erro, detalhados em `criterios-de-aceite.md` (cadastrar, listar, editar, remover, marcar status).

Saberei que a release está pronta para usar quando: API e front estiverem rodando e eu conseguir criar e listar um termo com sucesso pelo navegador.

## Quem precisa saber / ser avisado
- Doug precisa revisar e confirmar que está tudo ok.
- Eu preciso avisar/saber quando está pronto para usar.

## O que fazer se algo der errado
Se a API não subir, paro o `docker compose` e revejo o `Dockerfile`.
