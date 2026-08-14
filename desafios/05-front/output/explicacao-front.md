# Explicação — Front-end

## O que é front-end

Front-end é a parte do sistema que a pessoa usuária vê e usa diretamente: a tela, os botões, o formulário. No glossário, é a página onde eu vejo a lista de termos e consigo criar, editar, remover e marcar como entendido.

## Diferença entre front-end e back-end, na prática

No desafio 04 eu já tinha construído o back-end (a API em Go), que guarda os termos e sabe responder GET/POST/PUT/DELETE. Mas sem um front-end, só dava pra usar a API pelo terminal com `curl`.

Agora o front-end é uma página HTML com JavaScript que, em vez de eu digitar `curl` no terminal, faz essas mesmas chamadas pra API quando eu clico em um botão. O front não guarda nada sozinho — ele só pede pro back-end e mostra o resultado na tela.

## Como o front conversa com a API

Uso a função `fetch` do JavaScript pra fazer as requisições HTTP pra API, do mesmo jeito que o `curl` fazia:

- `fetch("http://localhost:8080/terms")` — busca a lista de termos (GET).
- `fetch(..., { method: "POST", body: JSON.stringify(termo) })` — cria um termo novo.
- `fetch(..., { method: "PUT", ... })` — atualiza um termo (uso isso tanto pra editar quanto pra marcar como entendido, só mudando o campo `status`).
- `fetch(..., { method: "DELETE" })` — remove um termo.

A resposta da API vem em JSON, e eu uso esses dados pra montar as linhas da tabela na tela.

## Decisão de tecnologia

Escolhi HTML/JS simples, sem framework (React ficou de fora por enquanto). O raciocínio completo está em `docs/decisoes.md` — resumindo, meu objetivo agora é entender o que cada linha de código faz, e um framework ia esconder isso.

## O que foi surpresa pra mim: CORS

A parte mais interessante desse desafio foi um bug que eu não esperava. Minha API já tinha um header de CORS (`Access-Control-Allow-Origin`), então eu achava que criar/editar/remover ia funcionar direto. Só que ao testar de verdade no navegador, o botão "Salvar" não fazia nada — sem erro visível, o formulário só ficava do jeito que estava.

O motivo: quando o `fetch` manda um corpo em JSON (`Content-Type: application/json`), o navegador manda antes uma requisição `OPTIONS` (chamada de "preflight"), perguntando pra API se ela aceita esse tipo de requisição. Minha API não tinha um handler pra `OPTIONS`, então respondia `405 Method Not Allowed` — e o navegador bloqueava a requisição de verdade (POST/PUT/DELETE) antes mesmo dela sair.

Isso só apareceu testando no navegador — pelo `curl` direto (sem passar pelo preflight do navegador) a API parecia estar funcionando. Foi um lembrete de que testar só via terminal não é suficiente pra front-end; o comportamento do navegador é diferente.

A correção foi adicionar um tratamento pra requisições `OPTIONS` na API, respondendo com os headers de CORS certos.
