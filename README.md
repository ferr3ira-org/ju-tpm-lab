# Ju TPM Lab

Bem-vinda, Ju! 💜

Este repositório é uma trilha prática para você estudar tecnologia, IA e fundamentos de TPM construindo um produto simples com ajuda de IA.

## Objetivo

Você vai aprender, aos poucos:

1. IA e LLM básico
2. Claude Code e skills
3. Git, GitHub, branch, commit e PR
4. API REST em Go
5. Front básico
6. Docker e Docker Compose
7. TPM aplicado ao produto

A meta **não é você virar desenvolvedora**. A meta é você conseguir:

- conversar melhor com times técnicos;
- entender conceitos básicos de engenharia;
- usar IA com critério;
- quebrar demandas em etapas;
- validar se algo realmente funciona;
- apresentar um produto final simples com API + front.

## Regra principal

Cada desafio deve ser feito em uma branch separada e entregue via Pull Request.

Você só segue para o próximo desafio depois que o Doug revisar e aprovar o PR.

```text
Ler desafio → criar branch → fazer output → abrir PR → pedir review → corrigir → mergear → próximo desafio
```

## Produto final

Você vai construir um **Glossário Tech para TPM**:

- backend: API REST em Go com CRUD;
- frontend: tela simples para listar, criar, editar e remover termos;
- Docker: tudo rodando com `docker compose`.

Exemplo de termo:

```json
{
  "id": 1,
  "termo": "API",
  "categoria": "Backend",
  "explicacao_simples": "Uma forma de sistemas conversarem entre si.",
  "exemplo": "Um app de clima consulta uma API para buscar a previsão.",
  "status": "entendido"
}
```

## Como estudar

Leia primeiro:

- [`COMO-ESTUDAR.md`](COMO-ESTUDAR.md)
- [`COMO-ABRIR-PR.md`](COMO-ABRIR-PR.md)

Depois comece pelo desafio:

- [`desafios/01-ia-llm/README.md`](desafios/01-ia-llm/README.md)
