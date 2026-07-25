# O que é uma Skill (no Claude Code)

## Explicação simples

Uma skill é um "modo de trabalho" salvo, com instruções prontas, que posso chamar sempre que quiser que a IA siga um jeito específico de fazer algo. Em vez de explicar tudo de novo toda vez, as instruções ficam escritas uma vez num arquivo, e depois é só acionar essa skill.

## Minha analogia

É como um receituário pré-salvo que um médico usa: não precisa tirar tudo da cabeça de novo, já tem o remédio, a dosagem — a única alteração seria o nome do paciente. Ou então como uma receita de pudim já pronta: eu não preciso mudar nada, só copiar o que já está escrito.

## A skill que usei: `explicar-conceito-tpm`

No repo já existia a skill `explicar-conceito-tpm`, guardada em `.claude/skills/explicar-conceito-tpm/SKILL.md`. Ela segue sempre este formato de resposta:

1. Explicação simples
2. Analogia
3. Exemplo prático
4. Como isso aparece em times de tecnologia
5. Pergunta para validar entendimento

Com tom simples, sem jargão desnecessário e com exemplos concretos.

## Exemplo prático

Testei essa skill com 3 termos técnicos (API, Docker e Pull Request) — os testes completos estão em `testes-da-skill.md`. Em todos os casos, a resposta seguiu o mesmo padrão, o que me ajudou a comparar os conceitos e responder as perguntas de validação com minhas próprias palavras.

## Como isso aparece em times de tecnologia

Times de tecnologia criam skills para tarefas que se repetem, revisar PR, ver bugs. Isso garante que qualquer pessoa do time, ou a IA, siga o mesmo padrão, sem lembrar tudo de cabeça todas as vezes.
