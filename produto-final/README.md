# Produto final completo

Rodando API + frontend via Docker Compose.

## Pré-requisitos

- Docker e Docker Compose instalados

## Como subir o projeto

```bash
cd produto-final
docker compose up --build
```

O `--build` reconstrói as imagens a partir dos Dockerfiles — útil caso o código tenha mudado desde a última vez.

## Onde acessar

- API: http://localhost:8080
- Frontend: http://localhost:3000

## Como parar

```bash
docker compose down
```

Esse comando para e remove os containers, mas não apaga as imagens que já foram construídas.
