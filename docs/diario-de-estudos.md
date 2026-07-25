# Diário de estudos

Use este arquivo para registrar o que você estudou.

## Modelo

```markdown
## YYYY-MM-DD — Tema estudado

### O que estudei

### O que entendi

### O que ainda ficou confuso

### Como usei IA

### Próximo passo
```

## 2026-07-21 — Desafio 01: IA e LLM básico

### O que estudei

Conceitos básicos de IA e LLM: IA generativa, LLM, prompt, contexto, alucinação, eval e por que não confiar cegamente na IA. Também pratiquei escrever prompts bons e melhorar prompts ruins.

### O que entendi

Entendi que um bom prompt precisa de contexto (quem está perguntando, o que já sabe) e de um pedido específico (o que exatamente eu quero de volta, em que formato). Também entendi que a IA responde com base em probabilidade, por isso pode alucinar, e que por isso preciso validar o que ela responde com meu próprio conhecimento.

### O que ainda ficou confuso

Nada, os conceitos ficaram claros.

### Como usei IA

Usei o Claude Code para me explicar os conceitos antes de escrever qualquer coisa, para dar feedback nos meus resumos e nos meus exemplos de prompt bom/ruim, e para me ajudar a organizar meus textos em markdown sem mudar o que eu quis dizer.

### Próximo passo

Revisar todos os entregáveis do desafio 01, abrir o PR e pedir review do Doug.

## 2026-07-25 — Desafio 02: Claude Code e skills

### O que estudei

O que é o Claude Code e como ele difere de um chat de IA comum, o conceito de skill (instrução reutilizável salva), e testei a skill `explicar-conceito-tpm` com 3 termos: API, Docker e Pull Request.

### O que entendi

Entendi que o Claude Code não só conversa, ele mexe direto nos arquivos e pastas do computador — cria, edita, roda comandos — sempre pedindo minha autorização antes. Entendi também que uma skill é um formato de resposta salvo, que a IA usa como padrão sempre que é acionado, sem precisar reexplicar tudo de novo (parecido com um receituário médico pré-salvo ou uma receita pronta).

### O que ainda ficou confuso

No início esqueci que o Claude Code tinha acesso real aos meus arquivos, achei que fosse só mais um chat tipo o ChatGPT que só devolve texto pronto. Também esqueci de rodar o `git status` antes do commit, mas fui lembrada e segui o passo certo.

### Como usei IA

Usei o Claude Code para me explicar os conceitos antes de criar qualquer arquivo, revisei e ajustei minhas próprias analogias em várias rodadas antes de aceitar, testei a skill com 3 termos e só autorizei a criação de cada arquivo depois de ver o rascunho completo.

### Próximo passo

Rodar `git status`, commitar os 4 entregáveis do desafio 02, abrir o PR e pedir review do Doug.
