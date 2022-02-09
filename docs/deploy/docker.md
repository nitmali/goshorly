# Docker installation

## Available Docker tags
- https://git.ucode.space/Phil/goshorly/container_registry/1

## docker-compose (with no reverse proxy)
```bash
mkdir goshorly
cd goshorly
wget https://git.ucode.space/Phil/goshorly/-/raw/main/docker-compose.yml
nano docker-compose.yml # Change the environment variables to your needs
docker-compose up -d
```

## docker-compose (with builtin reverse proxy)
```bash
mkdir goshorly
cd goshorly
wget https://git.ucode.space/Phil/goshorly/-/raw/main/docker-compose-proxy.yml
mv docker-compose-proxy.yml docker-compose.yml
nano docker-compose.yml # Change the command line on caddy to your domain & environment variables to your needs
docker-compose up -d
```
