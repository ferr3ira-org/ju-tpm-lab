# Como usei o Claude Code neste desafio

## Prompts que usei

- "Estou no segundo desafio" — para retomar o contexto de onde eu tinha parado.
- "Pode me explicar o que é claude code/skills, de forma fácil de entender... Sem criar arquivo ainda, somente me explique" — pedi explicação antes de qualquer criação de arquivo, como o README recomendava.
- Ao responder a pergunta de validação, propus minha própria analogia (o psicólogo) e pedi ajustes em partes específicas, tipo "pode ajustar essa parte da analogia pra incluir a ideia de..." — refinei em várias rodadas até ficar clara.
- "Pode rodar a skill com esses 3 termos agora, sem criar arquivo ainda, só pra eu ver o resultado" — para testar a skill `explicar-conceito-tpm` com API, Docker e Pull Request antes de documentar.
- Depois de cada rascunho de arquivo (`testes-da-skill.md`, `claude-code.md`, `skills.md`), pedi trocas pontuais — exemplo prático, uma resposta mais completa, uma analogia diferente — antes de aprovar com "sim, pode criar".

## O que a IA me ajudou a entender

- A diferença entre um chat de IA comum e o Claude Code: ele não só fala, ele mexe direto nos arquivos e pastas do meu computador — algo que eu tinha esquecido no meio da conversa.
- O conceito de skill como um "padrão salvo" que a IA usa para responder sempre da mesma forma, sem precisar reexplicar tudo toda vez.
- Como aplicar analogias do meu dia a dia (psicólogo, receituário médico, receita de pudim, TCC em grupo) para fixar conceitos técnicos como API, Docker e Pull Request.

## O que revisei antes de aceitar

- Toda analogia que eu mesma criei, revisei e ajustei em mais de uma rodada antes de considerar "pronta" (a do Claude Code como psicólogo, por exemplo, passou por três versões).
- Em cada arquivo (`testes-da-skill.md`, `claude-code.md`, `skills.md`), pedi para ver o rascunho completo antes de autorizar a criação, e troquei trechos que não refletiam bem o que eu queria dizer (exemplos práticos, respostas de validação, frases da seção "como aparece em times").
- Só autorizei a criação de cada arquivo depois de revisar o conteúdo final e confirmar explicitamente com "sim, pode criar".
- Esqueci de rodar o `git status` antes do commit, mas fui lembrada disso, e seguimos em frente rodando o comando para conferir que não subiu nada sensível (token, `.env`, chave SSH) antes de commitar.
