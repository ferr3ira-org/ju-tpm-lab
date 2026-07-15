# Desafio 02 — Claude Code e skills

## Onde colocar sua entrega

Coloque **todos os arquivos deste desafio** em:

```text
desafios/02-claude-code-skills/output/
```

Arquivos esperados:

- `desafios/02-claude-code-skills/output/claude-code.md`
- `desafios/02-claude-code-skills/output/skills.md`
- `desafios/02-claude-code-skills/output/testes-da-skill.md`
- `desafios/02-claude-code-skills/output/como-usei-claude-code.md`

Não altere a pasta de outro desafio neste PR.


## Regra obrigatória

Este desafio deve ser feito usando **Claude Code**.

Além dos arquivos principais, registre em `como-usei-claude-code.md` quais prompts você usou, o que a IA ajudou a entender e o que você revisou antes de aceitar.

## Objetivo

Aprender a usar Claude Code como par técnico e entender o conceito de skills.

## O que são skills?

Skills são instruções/procedimentos reutilizáveis que ajudam a IA a executar melhor uma tarefa específica.

## Passos

1. Abra o terminal Linux na pasta de estudos, não dentro do repo:

   ```bash
   cd ~/estudos
   claude --add-dir ./ju-tpm-lab
   ```

2. No Claude Code, peça primeiro orientação sem criar arquivos:

   ```text
   Estou no desafio 02 do repo ./ju-tpm-lab. Sou iniciante. Leia o README do desafio e me explique o que devo fazer. Não crie nem edite arquivos ainda.
   ```

3. Criar/usar uma conta de estudos para Claude Code.
4. Pedir para ele explicar a estrutura do repo.
5. Criar uma skill simples chamada `explicar-conceito-tpm`.
6. Testar a skill com 3 termos: API, Docker e Pull Request.
7. Quando for criar arquivos, peça o caminho completo dentro do repo, por exemplo:

   ```text
   Crie o arquivo ./ju-tpm-lab/desafios/02-claude-code-skills/output/claude-code.md. Não altere outros arquivos.
   ```

8. Escrever o que funcionou e o que ficou confuso.
9. Antes de commit, rodar `git status` e conferir se não apareceu arquivo sensível, token, `.env`, chave SSH ou configuração local.

## Entregáveis

```text
desafios/02-claude-code-skills/output/claude-code.md
desafios/02-claude-code-skills/output/skills.md
desafios/02-claude-code-skills/output/testes-da-skill.md
```

## Critérios de aceite

- [ ] Explica o que é Claude Code.
- [ ] Explica o que é uma skill.
- [ ] Criou ou descreveu uma skill útil para estudos.
- [ ] Mostrou testes com pelo menos 3 termos.
- [ ] Atualizou o diário de estudos.


## Checklist antes de abrir PR

- [ ] Coloquei os arquivos na pasta `output/` deste desafio.
- [ ] Atualizei `docs/diario-de-estudos.md`.
- [ ] Sei explicar com minhas palavras o que fiz.
- [ ] Usei IA como ajuda, mas revisei o resultado.
- [ ] Abri PR pedindo review do Doug.
