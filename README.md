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

## Produto único da trilha: Glossário Tech para TPM

Durante toda a trilha você vai desenvolver **um único produto** chamado **Glossário Tech para TPM**.

Você não precisa escolher uma aplicação. A aplicação já está definida para que você possa focar em aprender tecnologia, documentação, Git, IA, API, front e Docker.

O produto é uma aplicação simples para **cadastrar e consultar termos técnicos usados no dia a dia de tecnologia**.

Cada termo possui:

| Campo | O que significa |
|---|---|
| `termo` | Nome do conceito técnico, por exemplo `API`. |
| `categoria` | Grupo do termo, por exemplo `Backend`, `DevOps`, `Produto`, `Dados`. |
| `explicacao_simples` | Explicação em linguagem fácil. |
| `exemplo` | Exemplo prático de uso no dia a dia. |
| `status` | `estudando` ou `entendido`. |

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

Ao final da trilha, a aplicação terá:

- **API REST** para criar, listar, buscar, editar e remover termos;
- **front-end simples** para usar o glossário pelo navegador;
- **Docker Compose** para rodar tudo com um comando.

## Documentação antes da implementação

Antes de implementar qualquer funcionalidade, você vai aprender a **documentar o produto**.

Isso é importante porque uma das responsabilidades de uma TPM é ajudar times de produto e engenharia a transformar uma ideia em algo claro, planejado e executável.

A IA pode ser usada como copiloto durante todo esse processo, mas todas as decisões precisam ser compreendidas, revisadas e explicadas por você.

Durante a trilha, você vai produzir três documentos principais:

### 1. PRD — Product Requirements Document

Você vai aprender a documentar uma funcionalidade usando um PRD simples com:

- objetivo;
- problema;
- escopo;
- requisitos funcionais;
- critérios de aceite.

Arquivo guia: [`docs/prd.md`](docs/prd.md)

### 2. API Design

Antes de implementar a API, você vai desenhar os endpoints, requests e responses.

A ideia é entender como uma API é planejada antes de ser desenvolvida.

Arquivo guia: [`docs/api-design.md`](docs/api-design.md)

### 3. Data Modeling

Você vai identificar quais dados precisam ser armazenados e como essas informações se relacionam.

Para este produto, o modelo inicial será simples: um termo técnico com categoria, explicação, exemplo e status.

Arquivo guia: [`docs/data-modeling.md`](docs/data-modeling.md)

Essa documentação servirá como guia para o desenvolvimento e permitirá que você pratique um fluxo parecido com o usado por times de produto e engenharia no dia a dia.

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
