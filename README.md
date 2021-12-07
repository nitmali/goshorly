![](https://git.ucode.space/Phil/goshorly/badges/main/pipeline.svg)
# goshorly

An easy self-hosted Link shortner in Golang with Redis <3

**WARNING:**
- goshorly is in an early stage, it is not an Final Version!

Installation with Docker-Compose (with no reverse proxy / own proxy):
```bash
mkdir goshorly
cd goshorly
wget https://git.ucode.space/Phil/goshorly/-/raw/main/docker-compose.yml
nano docker-compose.yml # Change the environment variables to your needs
docker-compose up -d
```

Installation with Docker-Compose (built in proxy / caddy as reverse proxy):
```bash
mkdir goshorly
cd goshorly
wget https://git.ucode.space/Phil/goshorly/-/raw/main/docker-compose-proxy.yml
mv docker-compose-proxy.yml docker-compose.yml
nano docker-compose.yml # Change the command line on caddy to your domain & environment variables to your needs
docker-compose up -d
```
