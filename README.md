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

## Comece por aqui

1. Leia [`GUIA-INICIAL.md`](GUIA-INICIAL.md) para configurar VS Code, Markdown preview, SSH e clone do repo.
2. Leia [`COMO-ESTUDAR.md`](COMO-ESTUDAR.md).
3. Leia [`COMO-ABRIR-PR.md`](COMO-ABRIR-PR.md).
4. Comece pelo [`desafios/01-ia-llm/README.md`](desafios/01-ia-llm/README.md).

## Regra dos outputs

Cada desafio tem uma pasta `output/`. Sua entrega deve ficar sempre nessa pasta.

Exemplo:

```text
desafios/01-ia-llm/output/resumo.md
```

Não misture entregas de desafios diferentes no mesmo Pull Request.
