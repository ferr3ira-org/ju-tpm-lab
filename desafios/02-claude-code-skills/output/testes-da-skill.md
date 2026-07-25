# Testes da skill `explicar-conceito-tpm`

Testei a skill `explicar-conceito-tpm` com 3 termos técnicos, para validar se o formato de resposta (explicação simples, analogia, exemplo prático, como aparece em times, pergunta de validação) funciona bem para uma iniciante.

## Teste 1 — API

**Explicação simples:** API é um jeito de dois programas conversarem entre si, seguindo regras combinadas.

**Analogia:** É como um garçom de restaurante — você não entra na cozinha, faz o pedido ao garçom, ele leva e traz a resposta pronta.

**Exemplo prático:** A Alexa, por exemplo — quando você pergunta "que horas são?", ela não sabe sozinha. Ela faz uma chamada de API para consultar essa informação em algum lugar, e devolve a resposta pra você.

**Pergunta de validação:** Se o garçom fosse trocado por outro, mas continuasse entregando os mesmos pratos, o cliente perceberia?

**Minha resposta:** Acredito que o cliente final não perceberia a diferença, só se o prato dele viesse diferente. A troca do garçom não interfere diretamente no prato dele, que foi o que ele pediu e está com expectativa. Se mudasse o prato, sim, aí faria diferença, pois não foi aquilo que foi solicitado.

## Teste 2 — Docker

**Explicação simples:** Docker empacota um programa com tudo que ele precisa para funcionar, num pacote só que roda igual em qualquer computador.

**Analogia:** É como uma marmita pronta — não importa em qual micro-ondas você esquenta, o resultado é sempre o mesmo.

**Exemplo prático:** No desafio 06, vou rodar a API + banco de dados do Glossário Tech em containers Docker, evitando o problema de "na minha máquina funciona".

**Pergunta de validação:** Por que "funciona na minha máquina, mas não na do meu colega" é um problema comum sem Docker, e como a marmita resolve isso?

**Minha resposta:** É um problema comum sem Docker porque as configurações ficam diferentes em cada computador — sem padrão, não tem como saber se vai funcionar em outra máquina. A marmita resolve porque já vai tudo do mesmo jeito, então não tem como faltar ou trocar algum item.

## Teste 3 — Pull Request

**Explicação simples:** Pull Request é um pedido formal para juntar as mudanças feitas num código com a versão principal do projeto.

**Analogia:** É como entregar um TCC na faculdade. Se estamos fazendo em grupo, cada um faz a sua parte, coloca no projeto e antes de salvar, o professor orientador avalia se o que colocamos faz sentido e permite ou não colocarmos.

**Exemplo prático:** Foi o que fiz no desafio 1, criei a branch, abri o PR e pedi pro Doug revisar. Só depois que ele revisou, que eu mergeei.

**Pergunta de validação:** Por que é mais seguro passar pelo Doug antes de mergear, em vez de cada um colocar as mudanças direto na main?

**Minha resposta:** É mais seguro porque às vezes pode ter alguma coisa errada que quem fez não percebeu, e ter mais um olhar evita erros, bugs e problemas lá na frente. Isso diminui os riscos de subir algo e prejudicar o produto inteiro.

## Conclusão

Os 3 testes mostraram que a skill segue um formato consistente e me ajudou a entender os conceitos de verdade — não só decorar, já que consegui responder as perguntas de validação com minhas próprias palavras.
