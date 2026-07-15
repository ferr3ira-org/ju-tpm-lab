# Como abrir branch, commit e Pull Request

Este tutorial é para você conseguir entregar cada desafio.


## Conceitos rápidos

- **Branch**: uma cópia paralela do projeto para você trabalhar sem mexer direto na `main`.
- **Commit**: um pacote de mudanças salvo com uma mensagem.
- **Push**: enviar sua branch e commits para o GitHub.
- **Pull Request (PR)**: pedido para revisar e juntar sua mudança na branch principal.
- **Review**: comentários/aprovação de outra pessoa antes do merge.
- **Merge**: juntar sua branch aprovada na `main`.

Regra da trilha:

> Você só faz merge depois que o Doug aprovar o PR.

## Uma vez no começo

Clone o repo:

```bash
git clone git@github.com:ferr3ira-org/ju-tpm-lab.git
cd ju-tpm-lab
```

Se SSH ainda não estiver configurado, use o desafio `03-git-ssh` como guia e peça ajuda ao Doug.

## Fluxo para cada desafio

### 1. Atualizar sua branch main

```bash
git checkout main
git pull origin main
```

### 2. Criar uma branch nova

Exemplo para o desafio 01:

```bash
git checkout -b desafio-01-ia-llm
```

### 3. Fazer os arquivos do desafio

Cada desafio tem uma pasta `output/`. Coloque sua entrega lá.

Exemplo:

```text
desafios/01-ia-llm/output/respostas.md
```

### 4. Ver o que mudou

```bash
git status
```

### 5. Adicionar arquivos

```bash
git add .
```

### 6. Criar commit

```bash
git commit -m "desafio 01: IA e LLM"
```

### 7. Enviar branch para o GitHub

```bash
git push -u origin desafio-01-ia-llm
```

### 8. Abrir Pull Request

No GitHub, clique em **Compare & pull request**.

No texto do PR, preencha:

```markdown
## O que fiz
- ...

## O que aprendi
- ...

## Onde tive dúvida
- ...

## Checklist
- [ ] Li o README do desafio
- [ ] Coloquei os outputs na pasta certa
- [ ] Atualizei o diário de estudos
- [ ] Consigo explicar com minhas palavras
```

### 9. Pedir review do Doug

Marque o Doug/Gabriel no PR ou envie o link para o Doug no Telegram.

### 10. Depois de aprovado

Clique em **Merge pull request** no GitHub.

Depois atualize seu computador:

```bash
git checkout main
git pull origin main
```

Agora pode seguir para o próximo desafio.
