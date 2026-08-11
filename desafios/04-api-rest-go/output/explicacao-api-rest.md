# Explicação — API REST

## O que é uma API

API é uma forma de um sistema conversar com o outro, busca informação, envia, altera, remove dados.

## O que é REST

REST é uma API organizada com endpoints como `/terms` e ações como GET, POST, PUT e DELETE.

## O que é endpoint

Endpoint é o endereço/caminho dentro de uma API. No Glossário Tech, `/terms` é o endpoint principal para trabalhar com os termos do glossário.

## O que é JSON

JSON é um jeito organizado de escrever dados para sistemas trocarem informações, é um formato de texto.

## O que são métodos HTTP

O método HTTP diz a ação que deve ser feita. São ações usadas para conversar com uma API. No que estou estudando posso citar GET, POST, PUT e DELETE. Usarei esses métodos para listar, criar, editar e remover termos do meu glossário.

## O que são status codes

Status codes são códigos de resposta que a API devolve para dizer se uma requisição feita deu certo ou deu erro. Por exemplo:

- `200 OK` — sucesso na busca/listagem/edição.
- `201 Created` — termo criado.
- `400 Bad Request` — dados errados.
- `404 Not Found` — quando procura algo que não existe.
- `500 Internal Server Error` — quando algo deu errado dentro da API, um erro inesperado.

## Diferença entre frontend e backend

Frontend é tudo aquilo visto pelo usuário, como por exemplo o glossário. Backend é aquilo que fica por trás, como se fossem os bastidores — o cliente não tem acesso aos códigos que existem por trás para que o glossário exista, ele só vê pronto.

A API REST em Go que vou construir agora é o backend: é ela que vai guardar os termos do glossário e processar. Enquanto o frontend é a tela em que a pessoa usa para cadastrar os termos.
