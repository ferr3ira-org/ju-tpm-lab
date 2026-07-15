# Como estudar nesta trilha

## 1. Use IA como tutora, não como cola

Você pode usar Claude Code ou outro assistente para:

- explicar conceitos;
- dar exemplos;
- revisar textos;
- ajudar a entender erros;
- sugerir próximos passos;
- gerar código inicial.

Mas você precisa conseguir explicar com suas palavras o que foi feito.

> IA ajuda, mas Ju decide. IA escreve, mas Ju entende. IA sugere, mas Ju valida.

## 2. Para cada desafio

1. Leia o README do desafio.
2. Crie uma branch.
3. Faça os arquivos pedidos dentro da pasta `output/`.
4. Faça commit.
5. Abra Pull Request.
6. Peça review do Doug.
7. Responda os comentários.
8. Faça merge só depois de aprovado.

## 3. Como pedir ajuda para IA

Bons prompts:

```text
Me explique API REST como se eu fosse iniciante em TPM, usando uma analogia simples.
```

```text
Quebre esse desafio em passos pequenos. Não resolva tudo por mim; me ajude a entender cada etapa.
```

```text
Revise minha explicação e me diga onde está confusa ou errada.
```

Prompts ruins:

```text
Faz tudo pra mim.
```

```text
Resolve esse desafio sem explicar.
```

## 5. Ferramentas sensíveis e segurança

Algumas ferramentas de IA, incluindo Claude Code, podem criar arquivos locais de configuração, memória, sessão ou cache.

Por isso:

- abra o Claude Code a partir da pasta `~/estudos`, não diretamente dentro do repo, quando estiver apenas estudando;
- sempre peça explicitamente onde ele deve criar arquivos;
- confira `git status` antes de cada commit;
- nunca commite `.env`, tokens, senhas, chaves SSH ou arquivos de configuração local.

Exemplo seguro:

```text
Crie o arquivo ./ju-tpm-lab/desafios/01-ia-llm/output/resumo.md. Não altere outros arquivos.
```

## 6. Diário de estudos

Atualize `docs/diario-de-estudos.md` sempre que concluir uma etapa.
