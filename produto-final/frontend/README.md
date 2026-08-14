# Frontend — Glossário Tech para TPM

Aqui será construída a interface do produto final.

## Funcionalidades esperadas

- listar termos;
- criar termo;
- editar termo;
- remover termo;
- marcar termo como entendido.

## Como rodar

Este é um front-end simples em HTML/JS puro, sem build — precisa só de um servidor estático e da API rodando.

1. Suba a API (na pasta `produto-final/api-go`):

   ```bash
   go run .
   ```

   Ela sobe em `http://localhost:8080`.

2. Em outro terminal, sirva esta pasta (`produto-final/frontend`) com um servidor estático. Exemplo com Python:

   ```bash
   python3 -m http.server 5500
   ```

3. Abra `http://localhost:5500/index.html` no navegador.

> Abrir `index.html` direto pelo navegador (`file://`) não funciona bem: o `fetch` para a API se comporta diferente vindo de um arquivo local. Sirva sempre por HTTP, mesmo que seja um servidor estático simples como o do passo 2.

## Tecnologia

HTML/JS simples, sem framework — decisão registrada em `docs/decisoes.md`.
