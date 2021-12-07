![](https://git.ucode.space/Phil/goshorly/badges/main/pipeline.svg)
# goshorly

An easy self-hosted Link shortner in Golang with Redis <3

Installation with Docker-Compose (no reverse proxy):
```bash
mkdir goshorly
cd goshorly
wget https://git.ucode.space/Phil/goshorly/-/raw/main/docker-compose.yml
docker-compose up -d
```

Installation with Docker-Compose (caddy as reverse proxy):
```bash
mkdir goshorly
cd goshorly
wget https://git.ucode.space/Phil/goshorly/-/raw/main/docker-compose-proxy.yml
mv docker-compose-proxy.yml docker-compose.yml
nano docker-compose.yml # Change the command line on caddy (Line 20) to your domain
docker-compose up -d
```
