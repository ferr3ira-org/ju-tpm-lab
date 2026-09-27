# Roadmap — Glossário Tech para TPM

## Agora/MVP (✅ Entregue)
- Cadastrar novo termo (com categoria, explicação simples e exemplo)
- Listar termos cadastrados (com categoria)
- Editar termo existente
- Remover termo
- Marcar status ("estudando" ou "entendido")

Essas 5 funcionalidades já foram implementadas nos desafios 04 (API), 05 (front) e 06 (Docker) — formam a base atual do produto.

## Próximo
- Busca avançada
- Anexar arquivos (prints e imagens)

**Por quê essa ordem:** os dois itens melhoram a experiência de quem já usa a ferramenta, sem exigir mudanças estruturais grandes (diferente de login/múltiplos usuários, que exigiria repensar autenticação e permissões).

Busca avançada vem primeiro porque resolve um problema de escala: conforme o glossário cresce, ficar mais rápido de localizar um termo se torna cada vez mais necessário. Anexar prints e imagens vem em seguida, pois torna o conteúdo mais visual e didático — a pessoa associa a imagem ao conceito, em vez de depender só de texto.

## Futuro
- Login de usuário
- Múltiplos usuários

**Por quê essa ordem:** ainda não é necessário porque a única usuária do glossário sou eu. O foco atual é melhorar o produto e suas funcionalidades antes de pensar em estrutura de autenticação e múltiplos usuários — que só fazem sentido quando (e se) mais pessoas passarem a usar a ferramenta.
