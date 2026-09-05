# Docker e Docker Compose — Conceitos

## O que é Docker?

Docker é uma ferramenta para rodar aplicações em containers.

## Imagem

Imagem é o molde. Ela contém tudo que a aplicação precisa para rodar: código, dependências e configuração, mas ainda não está em execução.

## Container

Container é o que está rodando — uma instância da imagem em execução.

## Dockerfile

Dockerfile é o arquivo que ensina o Docker a montar a imagem da aplicação.

## Porta

Porta é como se fosse uma entrada para acessar uma aplicação — a porta onde o projeto roda.

Geralmente existe um mapeamento entre a porta de fora (na minha máquina) e a porta de dentro (no container). Exemplo: `8080:80` significa acesso pela porta 8080 da minha máquina, que aponta para a porta 80 dentro do container.

## Volume

Volume guarda dados ou compartilha arquivos. É como se fosse uma mochila, um armário externo — assim as coisas não ficam pra trás esquecidas.

Sem volume, se eu apago ou recrio o container, os dados somem junto. Com volume, os dados sobrevivem mesmo se o container for destruído — por isso ele é muito importante.

## Docker Compose

Docker Compose é uma forma de subir vários containers juntos, usando um arquivo `docker-compose.yml`.

No meu desafio, o Docker Compose vai ajudar a rodar a API e o frontend como um comando só.
