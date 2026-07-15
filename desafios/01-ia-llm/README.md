# Desafio 01 — IA e LLM básico

## Onde colocar sua entrega

Coloque **todos os arquivos deste desafio** em:

```text
desafios/01-ia-llm/output/
```

Arquivos esperados:

- `desafios/01-ia-llm/output/resumo.md`
- `desafios/01-ia-llm/output/prompts.md`
- `desafios/01-ia-llm/output/reflexao.md`
- `desafios/01-ia-llm/output/como-usei-claude-code.md`

Não altere a pasta de outro desafio neste PR.


## Regra obrigatória

Este desafio deve ser feito usando **Claude Code**.

Além dos arquivos principais, registre em `como-usei-claude-code.md` quais prompts você usou, o que a IA ajudou a entender e o que você revisou antes de aceitar.

## Objetivo

Entender conceitos básicos de IA para usar Claude Code e outros assistentes com mais segurança.

## Conceitos

Você precisa explicar com suas palavras:

- O que é IA generativa
- O que é LLM
- O que é prompt
- O que é contexto
- O que é alucinação
- O que é eval
- Por que não devemos confiar cegamente na IA

## Passos

1. Abra o Claude Code a partir da pasta de estudos:

   ```bash
   cd ~/estudos
   claude --add-dir ./ju-tpm-lab
   ```

2. Peça explicações simples usando Claude Code. Comece com:

   ```text
   Estou no desafio 01 do repo ./ju-tpm-lab. Sou iniciante em IA. Leia o README do desafio e me explique os conceitos sem criar arquivos ainda.
   ```

3. Escreva seu próprio resumo.
4. Crie 5 exemplos de bons prompts.
5. Crie 5 exemplos de prompts ruins e depois melhore cada um.
6. Escreva uma reflexão: “como uma TPM pode usar IA no dia a dia?”.
7. Registre em `como-usei-claude-code.md` quais prompts você usou e o que revisou antes de aceitar.

## Entregáveis

```text
desafios/01-ia-llm/output/resumo.md
desafios/01-ia-llm/output/prompts.md
desafios/01-ia-llm/output/reflexao.md
```

## Critérios de aceite

- [ ] Explica os conceitos com linguagem simples.
- [ ] Tem exemplos próprios.
- [ ] Mostra diferença entre prompt ruim e prompt melhorado.
- [ ] Fala dos riscos de alucinação.
- [ ] Atualizou `docs/diario-de-estudos.md`.


## Checklist antes de abrir PR

- [ ] Coloquei os arquivos na pasta `output/` deste desafio.
- [ ] Atualizei `docs/diario-de-estudos.md`.
- [ ] Sei explicar com minhas palavras o que fiz.
- [ ] Usei IA como ajuda, mas revisei o resultado.
- [ ] Abri PR pedindo review do Doug.
