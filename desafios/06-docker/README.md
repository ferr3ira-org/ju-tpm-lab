# Desafio 06 — Docker e Docker Compose

## Onde colocar sua entrega

Coloque **todos os arquivos deste desafio** em:

```text
desafios/06-docker/output/
```

Arquivos esperados:

- `desafios/06-docker/output/explicacao-docker.md`
- `desafios/06-docker/output/instalacao-docker.md`

Não altere a pasta de outro desafio neste PR.


## Antes de começar: verificar se Docker existe

No terminal Linux, rode:

```bash
docker --version
docker compose version
```

Resultado esperado:

```text
Docker version ...
Docker Compose version ...
```

Se os dois comandos funcionarem, Docker já está instalado.

Se não funcionar, siga a instalação abaixo.

## O que é Docker?

Docker é uma ferramenta para rodar aplicações em **containers**.

Pense em um container como uma “caixinha” com tudo que a aplicação precisa para funcionar: código, dependências e configuração.

Isso ajuda porque o projeto roda de forma parecida em computadores diferentes.

## O que é Docker Compose?

Docker Compose é uma forma de subir vários containers juntos usando um arquivo chamado:

```text
docker-compose.yml
```

Neste projeto, a ideia é subir:

- um container para a API Go;
- um container para o frontend.

## Instalar Docker no Linux/WSL

> Se você estiver usando Windows com WSL, confirme com Gabriel/Doug se deve usar Docker Desktop ou Docker Engine dentro do WSL.

### Opção simples em Ubuntu/Debian

No terminal Linux:

```bash
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg
```

Adicionar chave oficial do Docker:

```bash
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
```

Adicionar repositório do Docker:

```bash
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

Instalar Docker e Compose:

```bash
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

Testar:

```bash
sudo docker run hello-world
```

## Usar Docker sem sudo

Por padrão, talvez você precise usar `sudo docker`.

Para poder usar `docker` sem `sudo`:

```bash
sudo usermod -aG docker $USER
```

Depois disso, feche e abra o terminal de novo.

Teste:

```bash
docker run hello-world
```

Se der erro de permissão, reinicie o terminal/WSL ou peça ajuda ao Doug.

## Confirmar Docker Compose

Use o formato novo:

```bash
docker compose version
```

Atenção: neste projeto vamos usar `docker compose` com espaço, não `docker-compose` com hífen.

## Objetivo

Rodar API e front usando Docker Compose.

## Conceitos

Explique com suas palavras:

- imagem
- container
- Dockerfile
- porta
- volume
- docker compose

## Entregáveis

```text
produto-final/api-go/Dockerfile
produto-final/frontend/Dockerfile
produto-final/docker-compose.yml
desafios/06-docker/output/explicacao-docker.md
desafios/06-docker/output/instalacao-docker.md
```

## Critérios de aceite

- [ ] Docker está instalado ou documentei onde precisei de ajuda.
- [ ] `docker --version` funciona.
- [ ] `docker compose version` funciona.
- [ ] `docker compose up` sobe o projeto.
- [ ] API responde.
- [ ] Front abre no navegador.
- [ ] README explica como rodar e parar.


## Checklist antes de abrir PR

- [ ] Coloquei os arquivos na pasta `output/` deste desafio.
- [ ] Atualizei `docs/diario-de-estudos.md`.
- [ ] Sei explicar com minhas palavras o que fiz.
- [ ] Usei IA como ajuda, mas revisei o resultado.
- [ ] Abri PR pedindo review do Doug.
