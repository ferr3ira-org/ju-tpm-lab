# Descrição do funcionamento — Front-end do Glossário

Descrição do fluxo testado localmente (API em `http://localhost:8080`, front-end servido em `http://localhost:5500`).

## Tela inicial

A página abre com o título "Glossário TECH para TPM", um formulário no topo (campos Termo, Categoria, Explicação simples, Exemplo, Status, e o botão "Salvar") e uma tabela abaixo, vazia se ainda não existe nenhum termo cadastrado.

## Criar termo

Preenchi o formulário com um termo de exemplo (ex: "API REST" / "Backend" / explicação / exemplo / status "Estudando") e cliquei em "Salvar".

Resultado: o formulário limpou, e uma nova linha apareceu na tabela com os dados digitados, status "estudando", e três botões na coluna Ações: "Editar", "Remover" e "Marcar como entendido".

## Editar termo

Cliquei em "Editar" na linha do termo criado.

Resultado: o formulário foi preenchido com os dados daquele termo, o botão "Salvar" virou "Atualizar", e um botão "Cancelar" apareceu ao lado. Alterei a categoria e cliquei em "Atualizar".

Resultado: a linha na tabela atualizou com o novo valor, e o formulário voltou ao estado normal (vazio, botão "Salvar").

## Marcar como entendido

Cliquei em "Marcar como entendido" na linha do termo.

Resultado: o status da linha mudou de "estudando" para "entendido", e o botão "Marcar como entendido" desapareceu da linha (só fica visível quando o status ainda é "estudando"). Continuaram "Editar" e "Remover".

## Remover termo

Cliquei em "Remover" na linha do termo.

Resultado: apareceu uma confirmação do navegador ("Remover este termo?"); ao confirmar, a linha sumiu da tabela.

## Evidência via API (curl)

Testei também o ciclo direto na API, pra confirmar que o front-end está chamando os endpoints certos:

```
POST /terms    -> 201 Created, termo criado com id
GET  /terms    -> 200 OK, lista com o termo criado
PUT  /terms/1  -> 200 OK, status atualizado para "entendido"
DELETE /terms/1 -> 200 OK, {"mensagem":"termo removido"}
GET  /terms    -> 200 OK, lista vazia novamente
```

## Observação

Não anexei prints de tela nesta entrega — descrevi o comportamento observado em cada passo do teste manual, como alternativa indicada no README do desafio ("Prints/logs, se houver").
