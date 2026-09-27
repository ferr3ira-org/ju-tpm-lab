# Critérios de Aceite — Glossário Tech para TPM

## Cadastrar termo
- [ ] Consigo preencher todos os campos (termo, categoria, explicação, exemplo, status) e salvar.
- [ ] Ao salvar, o termo aparece na listagem.
- [ ] Se eu deixar um campo obrigatório vazio, o sistema não deixa cadastrar e mostra `400 Bad Request` + mensagem de erro.

## Listar termos
- [ ] Consigo ver todos os termos cadastrados.
- [ ] Após remover um termo, ele some da listagem (`200 OK` + `[]`, se não houver mais nenhum).

## Editar termo
- [ ] Consigo atualizar um termo com outra informação e o resultado é `200 OK` + o termo atualizado.
- [ ] Se eu tentar atualizar com um campo obrigatório vazio, aparece `400 Bad Request` + mensagem de erro.

## Remover termo
- [ ] Consigo remover um termo e o resultado é `200 OK` + mensagem confirmando que foi removido com sucesso.
- [ ] Depois de remover, o termo não aparece mais ao listar de novo.

## Marcar status
- [ ] Consigo marcar o status como `estudando` ou `entendido` — apenas essas duas opções são aceitas.
- [ ] Se eu tentar salvar com um status inválido (diferente desses dois), aparece `400 Bad Request` + mensagem de erro.
