# Guia inicial — antes do primeiro desafio

Este guia é o começo da trilha. Ele existe para você conseguir preparar o ambiente, abrir o repositório no VS Code e entender como entregar os desafios.

## 1. Instalar/abrir o VS Code

Use o **Visual Studio Code** para estudar este repositório.

Por quê?

- Ele mostra arquivos e pastas com clareza.
- Ele tem terminal integrado.
- Ele consegue mostrar Markdown renderizado, como uma página bonita.
- Ele ajuda você a trabalhar com Git sem depender só de comandos.

## 2. Ler Markdown renderizado

Os arquivos deste repo são escritos em **Markdown** (`.md`). Markdown é um formato simples para escrever documentação.

No VS Code você pode ver o Markdown renderizado assim:

- Abra um arquivo `.md`, por exemplo `README.md`.
- Aperte `Ctrl+Shift+V` para abrir o preview.
- Ou clique com o botão direito no arquivo e escolha **Open Preview**.

Opcionalmente instale a extensão:

- **Markdown Preview Enhanced**

Mas o preview nativo do VS Code já é suficiente.

## 3. Configurar seu nome e email no Git

No terminal do VS Code:

```bash
git config --global user.name "Seu Nome"
git config --global user.email "seu-email-do-github@example.com"
```

Confira:

```bash
git config --global user.name
git config --global user.email
```

## 4. O que é uma chave SSH?

Uma **chave SSH** é uma forma segura do seu computador provar para o GitHub que ele é seu.

Pense como uma fechadura e uma chave:

- A **chave pública** é como uma fechadura que você pode colocar no GitHub.
- A **chave privada** é a chave real, que fica só no seu computador.

Quando você tenta acessar o GitHub, seu computador usa a chave privada para provar que combina com a chave pública cadastrada lá.

Regra importante:

> Você pode compartilhar a chave pública. Nunca compartilhe a chave privada.

Arquivos comuns:

```text
~/.ssh/id_ed25519      ← chave privada, não compartilhar
~/.ssh/id_ed25519.pub  ← chave pública, pode cadastrar no GitHub
```

## 5. Verificar se você já tem chave SSH

No terminal:

```bash
ls ~/.ssh/*.pub
```

Se aparecer algo como `id_ed25519.pub`, você já tem uma chave pública.

Se não aparecer, crie uma.

## 6. Criar uma chave SSH

Troque o email pelo email da sua conta GitHub:

```bash
ssh-keygen -t ed25519 -C "seu-email-do-github@example.com"
```

Quando perguntar onde salvar, pode apertar `Enter` para aceitar o padrão.

Quando perguntar senha, para estudos você pode apertar `Enter` duas vezes para deixar sem senha local.

## 7. Copiar a chave pública

Mostre a chave pública:

```bash
cat ~/.ssh/id_ed25519.pub
```

Copie tudo que aparecer. Começa com algo parecido com:

```text
ssh-ed25519 AAAA...
```

## 8. Cadastrar no GitHub

No GitHub:

1. Clique na sua foto no canto superior direito.
2. Vá em **Settings**.
3. Vá em **SSH and GPG keys**.
4. Clique em **New SSH key**.
5. Em **Title**, coloque algo como `notebook-ju`.
6. Em **Key**, cole a chave pública.
7. Clique em **Add SSH key**.

## 9. Testar conexão SSH

No terminal:

```bash
ssh -T git@github.com
```

Se aparecer uma pergunta de confirmação, digite:

```text
yes
```

Resultado esperado:

```text
Hi SEU_USUARIO! You've successfully authenticated...
```

Se deu erro, mande o erro para o Doug.

## 10. Clonar este repositório

Escolha uma pasta no seu computador e rode:

```bash
git clone git@github.com:ferr3ira-org/ju-tpm-lab.git
cd ju-tpm-lab
```

Abra no VS Code:

```bash
code .
```

Se `code .` não funcionar, abra o VS Code manualmente e escolha **File → Open Folder**.

## 11. Começar a estudar

Leia nesta ordem:

1. `README.md`
2. `COMO-ESTUDAR.md`
3. `COMO-ABRIR-PR.md`
4. `desafios/01-ia-llm/README.md`

## 12. Regra de entrega

Cada desafio tem uma pasta `output/`.

Você deve colocar sua entrega sempre dentro da pasta `output/` do desafio atual.

Exemplo:

```text
desafios/01-ia-llm/output/resumo.md
```

Não avance para o próximo desafio sem PR aprovado pelo Doug.
