# Glossário Git & GitHub

## Git
O Git é uma ferramenta onde eu versiono o meu código. Ele guarda todas as mudanças que foram feitas ao longo do tempo e se eu quiser ver algo que foi feito antes, onde foi mudado, tem como. Dá para voltar para a versão anterior. Quando estamos trabalhando em um projeto, cada pessoa pode trabalhar na sua parte sem interferir na do outro. O Git é diferente de GitHub. O Git fica instalado na minha máquina e não é como o GitHub que é uma plataforma.

## GitHub
O GitHub é a plataforma onde os códigos feitos ficam salvos, existem os repositórios e ele é como se fosse as pastas do meu Google Drive. Consigo compartilhar os meus projetos com outras pessoas, elas conseguem sugerir mudanças, comentar, colaborar, aprovar, então dá pra fazer bastante coisa. O GitHub usa o Git. Ele hospeda na nuvem os repositórios que eu crio localmente na minha máquina. O Git faz o versionamento e o GitHub guarda isso pra mim e todas as mudanças que foram feitas.

## Repositório
O repositório é o local onde tudo fica guardado. Como se fosse um armário e nesse armário eu guardo tudo relacionado a um assunto/projeto. E aí ficam guardadas todas as versões, etc dentro dele.

## Clone
Clone é quando eu faço uma cópia de um repositório do GitHub, para a minha máquina local. Dessa forma eu consigo trabalhar nele e mandar mudanças se preciso for.

## Branch
Branch é como se fosse uma aba separada de trabalho. Existe a branch main, que é a principal, e as demais branches que eu posso criar — um colega também pode criar as dele — e nós dois podemos estar trabalhando ao mesmo tempo, sem interferir no trabalho um do outro. Quando eu crio uma branch, ela nasce como uma cópia exata da main, então posso trabalhar nela sem bagunçar o que já está feito. Depois, se eu quiser, posso juntar essa branch de volta na main — mas isso só acontece se for aprovada num pull request.

## Commit
Commit é o que geralmente vem depois de você alterar uma branch. É registrar o que mudou, salvar o que foi feito. Quando eu faço commit, vem com uma mensagem descrevendo o que foi feito, e isso é importante para saber ao que cada mudança se refere. E eu consigo voltar ao que era antes se precisar.

## Push
Push é quando eu envio as mudanças feitas para o GitHub. Quando eu faço commit, essas mudanças ficam salvas somente na minha máquina, para subir para o GitHub eu tenho que fazer o push.

## Pull Request
Pull request é quando eu peço pra juntar a alteração que eu fiz na minha branch, junto a main. O pull request existe já para ter uma revisão de alguém. É como se fosse uma autorização de alguém para aprovar essa mudança pra main. No pull request as pessoas podem fazer comentários e pedir ajustes também antes de ter a completa aprovação.

## Review
O review é a análise da mudança que eu sugeri. A pessoa pode fazer sugestões, pedir algum ajuste ou aprovar. O PR é o pedido e o review é o ato de analisar o que foi pedido.

## Merge
Merge é a junção de uma branch a branch main. Por exemplo, eu fiz a minha alteração, pedi o pull request, alguém fez o review e foi aprovado, então eu posso fazer o merge, que é juntar a alteração que eu fiz na branch principal.

## Chave SSH
A chave SSH é uma forma segura do meu computador provar pro GitHub que ele é meu. A diferença entre a chave pública e a privada é que uma eu posso compartilhar e a outra não. A privada fica somente no meu computador, ninguém mais tem acesso e a pública é a que eu coloco no GitHub, é como uma fechadura. Já a chave privada é a chave real. Ela não deve ser compartilhada porque senão as pessoas podem se passar por mim e fazer alterações em repositórios como se fosse eu.

Eu testei a conexão dando o comando `ssh -T git@github.com` e respondi a pergunta de confirmação com `yes`, e depois disso tive a resposta `Hi juuhn! You've successfully authenticated...`, isso provou pro GitHub que a chave realmente era minha, assim eu pude clonar o repositório na minha máquina para trabalhar. Clonei rodando `git clone git@github.com:ferr3ira-org/ju-tpm-lab.git` e depois entrei na pasta com `cd ju-tpm-lab`. Confirmei que os arquivos apareceram com `ls`.
