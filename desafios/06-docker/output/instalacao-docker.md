# Instalação do Docker

## Como instalei

Segui os passos do README para instalar o Docker no Linux Mint. Durante a instalação, tive dois problemas que precisei resolver.

## Problema 1: erro 404 ao adicionar o repositório

Ao adicionar o repositório do Docker, tive um erro 404. O README usa `$VERSION_CODENAME` do `/etc/os-release`, e no Linux Mint isso resolve para `zena` — só que o repositório do Docker não tem esse codename, ele só usa os codenomes do Ubuntu.

**Solução:** troquei para `$UBUNTU_CODENAME` (que resolve para `noble`) no arquivo `/etc/apt/sources.list.d/docker.list`. Depois disso, a instalação seguiu normalmente.

## Problema 2: rodar Docker sem sudo não funcionava

Depois de instalar, tentei rodar `docker` sem `sudo`, mas não funcionava. Mesmo depois de rodar `sudo usermod -aG docker $USER` e reabrir o terminal, o grupo `docker` continuava não aparecendo no `groups`.

**Solução:** só funcionou depois de um reboot completo da máquina — reabrir o terminal não foi suficiente.

## Como usei IA

Encaminhei os erros que apareciam para o Claude, que foi me ajudando a identificar a causa de cada um para eu conseguir seguir em frente.

## Resultado final

Instalação confirmada com:

```bash
docker --version
# Docker version 29.7.2

docker compose version
# Docker Compose version v5.5.0

docker run hello-world
# rodou sem sudo, sem erros
```
