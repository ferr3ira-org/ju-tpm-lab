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

## 2026-07-28 — Desafio 03: Git, SSH, branch e Pull Request

### O que estudei

Hoje eu estudei sobre Git, GitHub, repositório, clone, branch, commit, push, pull request, review, merge e chave SSH. Montei um glossário com todos os seus significados para que eu possa fixar o conteúdo.

### O que entendi

Entendi melhor a diferença entre Git e GitHub, pois antes eu achava que eles eram praticamente iguais e vi que não. E através do glossário deu pra enxergar toda a sequência de fluxo que um código passa para ser publicado.

### O que ainda ficou confuso

O que ficou confuso ainda foi explicar o passo a passo do fluxo de abrir um PR. Nas duas primeiras vezes o Claude me ajudou, então eu não tenho certeza se fiz o comando `git checkout main` + `git pull origin main` e `git add`. Acredito que eu precise voltar nesse assunto depois para entender/gravar melhor.

### Como usei IA

Usei a IA nesse desafio como guia, mas eu quem escrevi as respostas do meu jeito, a IA só foi me falando o que eu podia acrescentar, dando alguns insights.

### Próximo passo

Revisar os entregáveis do desafio 03, abrir o PR e pedir review do Doug.

## 2026-08-10 — Desafio 04: API REST em Go

### O que estudei

Hoje estudei a construção de API REST em Go, os conceitos de API, REST, endpoint, JSON, métodos HTTP, status codes e a diferença entre front e backend. Como funciona o roteamento de uma API e como ela responde em JSON.

### O que entendi

Entendi que a struct vira JSON automaticamente com as tags.

### O que ainda ficou confuso

Achei esse desafio mais difícil e complexo. Ainda ficou confuso, se não fosse pelo Claude, seria bem difícil de dar os comandos corretos. Os conceitos de cada coisa foi fácil de entender, mas na hora de aplicar, não.

### Como usei IA

Usei a IA para me guiar e me dar exemplos, tirar dúvidas, melhorar as minhas respostas, revisando tudo o que era escrito.

### Próximo passo

Fazer o PR e ver a aprovação do Doug.

## 2026-08-13 — Desafio 05: Front-end

### O que estudei

Construí a interface do glossário (HTML/JS simples, sem framework) que consome a API do desafio 04: listar, criar, editar, remover e marcar termo como entendido. Estudei como o `fetch` do JavaScript faz as mesmas chamadas HTTP que eu fazia com `curl`, e como usar o resultado em JSON pra montar a tabela na tela.

### O que entendi

Entendi melhor, na prática, a diferença entre front e back: o front não guarda nada sozinho, ele só pede pro back-end e mostra o que volta. E entendi o que é CORS e por que existe.

### O que ainda ficou confuso

Tive um bug de CORS que não esperava: minha API já tinha o header `Access-Control-Allow-Origin`, mas mesmo assim criar/editar/remover não funcionava no navegador, sem erro visível. Foi confuso descobrir que o problema era o preflight `OPTIONS`, que o navegador manda antes do POST/PUT quando o corpo é JSON — e que testar só com `curl` não pega esse tipo de problema, porque o `curl` não faz esse preflight sozinho.

### Como usei IA

Usei o Claude Code pra escrever o front-end junto comigo, testar o fluxo completo no navegador (não só por `curl`), e pra investigar e corrigir o bug de CORS. Revisei o código e os textos antes de aceitar.

### Próximo passo

Fazer o PR e ver a aprovação do Doug.

## 2026-09-05 — Desafio 06: Docker e Docker Compose

### O que estudei

Estudei Docker: container, imagem, Dockerfile, docker-compose e portas. Nesse desafio 6 aprendi entre várias outras coisas, sobre o dockerfile, como ele funciona, o que colocar, como definir onde vou trabalhar. Os significados de FROM, WORKDIR, COPY, RUN, EXPOSE e CMD. A diferença do dockerfile da API para o do front-end. A diferença entre as portas onde fica cada coisa. Também tive dois problemas na instalação do Docker (um de codinome da distribuição, outro de permissão de grupo) que consegui resolver com ajuda do Claude Code. Consegui testar o docker compose para ver se tinha dado certo e abri o localhost:3000.

### O que entendi

Entendi que API e front precisam de portas diferentes e o motivo de eu ter colocado a de front como 3000. O Claude sugeriu, mas eu não sabia muito bem o porquê e ele me explicou. Duas coisas diferentes não podem estar na mesma porta ao mesmo tempo na minha máquina. Se a API já ocupa 8080, o front precisa de outra e não podia ser 80, porque 80 costuma ser a porta padrão da internet. E às vezes ela já está ocupada por outro programa, então não faria sentido colocar, porque poderia dar problema depois. Ele me explicou que 3000 é uma convenção do mercado, que muitas ferramentas de frontend usam 3000 como padrão de desenvolvimento, então quem visse ia reconhecer que 3000 é o front. Entendi que o Dockerfile é como se fosse a receita — tudo que está ali é pra mostrar onde e como minha aplicação vai rodar —, e que o docker-compose é o cardápio da refeição toda: ele diz quais containers eu quero que sirvam juntos e como eles se conectam.

### O que ainda ficou confuso

Foi principalmente a sintaxe do YAML, como eu não tenho conhecimento técnico e ainda não peguei muito de escrever código, pra mim fica um pouco difícil entender o que eu coloco, como coloco, pra que serve aquilo, etc. É como se eu tivesse que decorar, mas eu sei que não precisa, o que preciso é entender o conceito, que sinto que ainda preciso estudar mais sobre.

### Como usei IA

Usei a IA para me ajudar em diversas fases, o Claude me ajudou a entender os conceitos antes de eu escrever, leu os meus rascunhos e me deu toques do que eu podia melhorar para só depois eu aprovar. Ajudou a rodar o teste real para confirmar que tudo estava funcionando, e confirmamos que estava. Rodei o localhost e deu tudo certo. Também me ajudou a tentar entender melhor o docker compose na parte de sintaxe e montar os códigos da maneira correta.

### Próximo passo

Commitar tudo, dar push e abrir o PR.

## 2026-09-27 — Desafio 07: TPM — revisão de produto

### O que estudei

Tratei o Glossário Tech como um produto de verdade: revisei os 4 documentos de produto (produto.md, prd.md, api-design.md, data-modeling.md) pra confirmar se ainda batiam com o que foi construído, e produzi 6 entregáveis: backlog, critérios de aceite, riscos, roadmap, plano de release e revisão de documentação do produto.

### O que entendi

Entendi a diferença entre backlog (tudo que quero fazer, sem ordem de tempo) e roadmap (quando e em que ordem, com o porquê da prioridade). Entendi que critérios de aceite da entrega do produto são diferentes dos documentos entregáveis do desafio — no início confundi os dois no plano de release, mas consegui corrigir depois de entender a diferença. Também entendi melhor como apresentar o mesmo produto de formas diferentes: pra um dev, falando de tecnologia e endpoints; pra um stakeholder, falando do que o produto resolve e do valor, sem termos técnicos.

### O que ainda ficou confuso

No início do desafio, no arquivo de riscos, entendi errado que os dados sobreviveriam a um restart mesmo com armazenamento em memória — me corrigi depois de entender melhor o conceito.

### Como usei IA

Usei o Claude Code pra me explicar cada conceito antes de escrever (roadmap, plano de release, revisão de documentação), dei meus rascunhos, recebi feedback específico do que ajustar, e só depois de revisar aprovei a criação de cada arquivo.

### Próximo passo

Commitar os 6 entregáveis, dar push e abrir o PR — esse é o último desafio da trilha.
