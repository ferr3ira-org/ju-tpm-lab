# Riscos — Glossário Tech para TPM

## Risco: Perda de dados ao reiniciar a API
- **Impacto:** Alto
- **Descrição:** A API reinicia, os dados cadastrados são perdidos, então as informações inseridas não estariam mais disponíveis.
- **Mitigação:** Pensar no futuro em um armazenamento em banco de dados, para que as informações não se percam caso a API reinicie.

## Risco: API não rodar em outra máquina
- **Impacto:** Alto
- **Descrição:** A API pode não rodar em outras máquinas por diferenças de ambiente (versão do Go, sistema operacional, dependências).
- **Mitigação:** Já mitigado — a API foi empacotada com Docker (desafio 06), garantindo que rode do mesmo jeito em qualquer máquina que tenha Docker instalado.

## Risco: Cadastro manual não escalar
- **Impacto:** Médio
- **Descrição:** Se o glossário crescer muito (ex: mais de 500 termos), cadastrar cada termo manualmente se torna cansativo e lento.
- **Mitigação:** Pensar em uma solução de preenchimento automático — por exemplo, um programa que sugira ou cadastre termos identificados automaticamente em outros sistemas.

## Risco: CORS / preflight não tratado
- **Impacto:** Alto — é um risco silencioso, o botão simplesmente não faz nada, dificultando o uso do programa para quem utiliza.
- **Descrição:** Requisições OPTIONS (preflight), enviadas automaticamente pelo navegador antes de POST/PUT/DELETE, não eram tratadas pela API — que respondia 405 Method Not Allowed, fazendo o navegador bloquear a ação de verdade, sem erro visível para quem utiliza. Isso só aparecia testando no navegador; pelo curl direto, a API parecia funcionar normalmente.
- **Mitigação:** Adicionado tratamento de requisições OPTIONS na API, respondendo com os headers de CORS corretos.
