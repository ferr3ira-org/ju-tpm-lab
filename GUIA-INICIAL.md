# Guia inicial — antes do primeiro desafio

Este guia é o começo da trilha. Ele foi escrito assumindo que você ainda está começando e talvez não saiba o que é terminal, Git, GitHub, SSH ou repositório.

A ideia é ir devagar: preparar o ambiente, clonar o repositório, abrir no VS Code e só depois começar os desafios.

## 0. Antes de tudo: o que você vai usar

Você vai usar quatro coisas principais:

| Ferramenta | Para que serve |
|---|---|
| **GitHub** | Site onde o repositório do projeto fica salvo. |
| **Git** | Ferramenta que controla versões dos arquivos. |
| **Terminal Linux** | Lugar onde você digita comandos. |
| **VS Code** | Editor onde você lê os arquivos, escreve respostas e abre o terminal. |

## 1. O que é terminal?

O **terminal** é uma tela onde você conversa com o computador usando comandos de texto.

Exemplo:

```bash
pwd
```

Esse comando mostra em qual pasta você está.

Outros comandos básicos:

```bash
ls
```

Mostra arquivos e pastas.

```bash
cd nome-da-pasta
```

Entra em uma pasta.

```bash
cd ..
```

Volta uma pasta.

```bash
mkdir estudos
```

Cria uma pasta chamada `estudos`.

Quando este guia disser “rode no terminal”, significa: abra o terminal e digite o comando exatamente como aparece.

## 2. Use o terminal Linux

Para estes passos iniciais, use um **terminal Linux**.

Pode ser:

- Linux direto no computador;
- WSL no Windows;
- terminal integrado do VS Code usando Linux/WSL;
- outro terminal Linux que o Gabriel/Doug configurar.

Se você estiver no Windows puro, peça ajuda para abrir o terminal correto com WSL.

## 3. O que é Git?

**Git** é uma ferramenta que guarda o histórico de mudanças dos arquivos.

Pense como um “histórico de versões” do projeto.

Com Git você consegue:

- ver o que mudou;
- salvar uma versão com uma mensagem;
- criar uma branch para trabalhar sem bagunçar a versão principal;
- enviar suas mudanças para o GitHub;
- abrir um Pull Request para revisão.

Alguns termos importantes:

| Termo | Explicação simples |
|---|---|
| **Repositório/repo** | Pasta do projeto controlada pelo Git. |
| **Clone** | Baixar uma cópia do repo do GitHub para seu computador. |
| **Branch** | Uma linha separada de trabalho. |
| **Commit** | Um pacote de mudanças salvo com uma mensagem. |
| **Push** | Enviar seus commits para o GitHub. |
| **Pull Request/PR** | Pedido para revisar e juntar sua mudança. |
| **Merge** | Juntar uma branch aprovada na versão principal. |

Você não precisa decorar tudo agora. Vai praticar isso nos desafios.

## 4. O que é GitHub?

**GitHub** é um site onde times guardam repositórios de código e documentação.

Neste desafio, o repo está aqui:

```text
https://github.com/ferr3ira-org/ju-tpm-lab
```

O GitHub vai ser usado para:

- guardar o projeto;
- abrir Pull Requests;
- pedir review do Doug;
- registrar os desafios em issues;
- acompanhar seu progresso.

## 5. Instalar/abrir o VS Code

Use o **Visual Studio Code** para estudar este repositório.

Por quê?

- Ele mostra arquivos e pastas com clareza.
- Ele tem terminal integrado.
- Ele consegue mostrar Markdown renderizado, como uma página bonita.
- Ele ajuda você a trabalhar com Git sem depender só de comandos.

## 6. O que é Markdown?

Os arquivos deste repo são escritos em **Markdown** (`.md`).

Markdown é um jeito simples de escrever documentação com títulos, listas, links e blocos de código.

Exemplo de Markdown:

```markdown
# Título

- item 1
- item 2
```

No VS Code você pode ver o Markdown renderizado como uma página mais bonita.

## 7. Ler Markdown renderizado no VS Code

No VS Code:

1. Abra um arquivo `.md`, por exemplo `README.md`.
2. Aperte `Ctrl+Shift+V` para abrir o preview.
3. Ou clique com o botão direito no arquivo e escolha **Open Preview**.

Opcionalmente instale a extensão:

- **Markdown Preview Enhanced**

Mas o preview nativo do VS Code já é suficiente.

## 8. Abrir o terminal no VS Code

No VS Code:

1. Vá no menu **Terminal**.
2. Clique em **New Terminal**.
3. Uma área de terminal vai abrir embaixo.

É nesse terminal que você vai rodar comandos como:

```bash
git status
```

Se você estiver usando Windows, confirme que o terminal aberto é Linux/WSL.

## 9. Verificar se Git existe

No terminal Linux, rode:

```bash
git --version
```

Se aparecer algo como:

```text
git version 2.x.x
```

então Git está instalado.

Se aparecer erro dizendo que o comando não existe, peça ajuda ao Gabriel/Doug para instalar Git.

## 10. Configurar seu nome e email no Git

O Git precisa saber quem está fazendo os commits.

No terminal:

```bash
git config --global user.name "Seu Nome"
git config --global user.email "seu-email-do-github@example.com"
```

Troque:

- `Seu Nome` pelo seu nome;
- `seu-email-do-github@example.com` pelo email da sua conta GitHub.

Confira:

```bash
git config --global user.name
git config --global user.email
```

## 11. O que é SSH?

**SSH** é uma forma segura de conectar seu computador a outro sistema.

Neste caso, vamos usar SSH para o seu computador conversar com o GitHub sem precisar digitar senha toda hora.

## 12. O que é uma chave SSH?

Uma **chave SSH** é uma forma segura do seu computador provar para o GitHub que ele é seu.

Pense como fechadura e chave:

- A **chave pública** é como uma fechadura. Você pode colocar no GitHub.
- A **chave privada** é a chave real. Ela fica só no seu computador.

Quando você tenta acessar o GitHub, seu computador usa a chave privada para provar que combina com a chave pública cadastrada lá.

Regra importante:

> Você pode compartilhar a chave pública. Nunca compartilhe a chave privada.

Neste guia, vamos criar a chave com o nome da Ju para ficar fácil de reconhecer:

```text
~/.ssh/ju_github      ← chave privada, não compartilhar
~/.ssh/ju_github.pub  ← chave pública, pode cadastrar no GitHub
```

O símbolo `~` significa sua pasta de usuário no Linux.

## 13. Verificar se você já tem chave SSH

No terminal:

```bash
ls ~/.ssh/*.pub
```

O comando `ls` lista arquivos.

Se aparecer algo como:

```text
/home/seu-usuario/.ssh/ju_github.pub
```

você já tem uma chave pública da Ju.

Se aparecer outra chave, como `id_ed25519.pub`, ela pode até funcionar, mas para este guia vamos usar `ju_github.pub` para ficar mais fácil de entender.

Se não aparecer `ju_github.pub`, crie uma chave nova no próximo passo.

## 14. Criar uma chave SSH

Crie uma chave com nome fácil de reconhecer:

```bash
ssh-keygen -t ed25519 -C "ju-github" -f ~/.ssh/ju_github
```

O que esse comando faz?

- `ssh-keygen` cria uma chave SSH.
- `-t ed25519` escolhe um tipo moderno de chave.
- `-C` adiciona um comentário, normalmente seu email.

Quando perguntar onde salvar, pode apertar `Enter` para aceitar o padrão.

Quando perguntar senha da chave, para estudos você pode apertar `Enter` duas vezes para deixar sem senha local.

## 15. Copiar a chave pública

Mostre a chave pública:

```bash
cat ~/.ssh/ju_github.pub
```

O comando `cat` mostra o conteúdo de um arquivo.

Copie tudo que aparecer. Começa com algo parecido com:

```text
ssh-ed25519 AAAA...
```

Atenção: copie o arquivo que termina com `.pub`. Esse é o público.

## 16. Cadastrar a chave pública no GitHub

No GitHub:

1. Clique na sua foto no canto superior direito.
2. Vá em **Settings**.
3. Vá em **SSH and GPG keys**.
4. Clique em **New SSH key**.
5. Em **Title**, coloque `ju_github` para ficar igual ao nome do arquivo da chave.
6. Em **Key**, cole a chave pública.
7. Clique em **Add SSH key**.

## 17. Testar conexão SSH com GitHub

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

Isso significa que seu computador conseguiu provar para o GitHub que a chave é sua.

Se deu erro, copie a mensagem e mande para o Doug.

## 18. Escolher onde salvar seus estudos

Agora escolha uma pasta para guardar o repo.

Exemplo:

```bash
mkdir -p ~/estudos
cd ~/estudos
```

Explicando:

- `mkdir -p ~/estudos` cria a pasta `estudos` se ela ainda não existir.
- `cd ~/estudos` entra nessa pasta.

Para confirmar onde você está:

```bash
pwd
```

## 19. Clonar este repositório

Clonar significa baixar uma cópia do repo do GitHub para seu computador.

Rode:

```bash
git clone git@github.com:ferr3ira-org/ju-tpm-lab.git
```

Depois entre na pasta:

```bash
cd ju-tpm-lab
```

Confirme que os arquivos apareceram:

```bash
ls
```

Você deve ver arquivos como:

```text
README.md
GUIA-INICIAL.md
COMO-ESTUDAR.md
COMO-ABRIR-PR.md
desafios
```

## 20. Abrir o repo no VS Code

Dentro da pasta `ju-tpm-lab`, rode:

```bash
code .
```

O ponto `.` significa “a pasta atual”.

Se `code .` não funcionar, abra o VS Code manualmente e escolha:

```text
File → Open Folder → ju-tpm-lab
```

## 21. Conferir se está tudo certo

No terminal do VS Code, rode:

```bash
git status
```

Se aparecer algo como:

```text
On branch main
Your branch is up to date with 'origin/main'.
nothing to commit, working tree clean
```

está tudo certo.

Não precisa entender tudo ainda. O importante é saber que o Git reconheceu o repo.

## 22. Começar a estudar

Leia nesta ordem, usando o preview de Markdown do VS Code:

1. `README.md`
2. `COMO-ESTUDAR.md`
3. `COMO-ABRIR-PR.md`
4. `desafios/01-ia-llm/README.md`

## 23. Regra de entrega

Cada desafio tem uma pasta `output/`.

Você deve colocar sua entrega sempre dentro da pasta `output/` do desafio atual.

Exemplo:

```text
desafios/01-ia-llm/output/resumo.md
```

Não avance para o próximo desafio sem PR aprovado pelo Doug.

## 24. Se travar

Se alguma coisa der errado:

1. Copie o comando que você tentou rodar.
2. Copie a mensagem de erro.
3. Mande para o Doug.
4. Explique em qual passo do guia você estava.

Exemplo:

```text
Estou no passo 17. Rodei ssh -T git@github.com e apareceu esse erro: ...
```
