# Data Modeling

**Data Modeling** é pensar quais dados precisam ser armazenados e como eles se organizam.

Antes de implementar, você precisa entender quais informações fazem parte do produto.

## Entidade principal

A entidade principal do Glossário Tech para TPM é:

```text
Termo
```

## Campos do Termo

| Campo | Tipo simples | Obrigatório? | Exemplo |
|---|---|---|---|
| `id` | número | sim | `1` |
| `termo` | texto | sim | `API` |
| `categoria` | texto | sim | `Backend` |
| `explicacao_simples` | texto | sim | `Uma forma de sistemas conversarem.` |
| `exemplo` | texto | sim | `Um app consulta uma API de clima.` |
| `status` | texto | sim | `estudando` |

## Status permitido

O campo `status` deve aceitar apenas:

- `estudando`
- `entendido`

## Relacionamentos

No início, o modelo não terá relacionamento com outras tabelas.

Ou seja: cada termo é independente.

Mais para frente, poderiam existir outras entidades, como:

- usuário;
- tags;
- histórico de estudo;
- fonte do termo.

Mas isso está fora do escopo inicial.

## Perguntas para revisar o modelo

- Quais dados são obrigatórios?
- Quais dados podem ficar vazios?
- Existe algum campo que só pode aceitar valores específicos?
- O modelo está simples o suficiente para o MVP?
- Consigo explicar por que cada campo existe?

## Uso da IA

Você pode pedir para a IA sugerir campos e revisar o modelo, mas precisa validar se cada campo realmente faz sentido para o produto.
