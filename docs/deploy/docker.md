# Docker installation

## docker-compose (with no reverse proxy)
```bash
git clone https://github.com/nitmali/goshorly.git
cd goshorly
docker build -t goshorly:latest .
nano docker-compose.yml # Change the environment variables to your needs
docker-compose up -d
```

## docker-compose (with builtin reverse proxy)
```bash
git clone https://github.com/nitmali/goshorly.git
cd goshorly
docker build -t goshorly:latest .
mv docker-compose-proxy.yml docker-compose.yml
nano docker-compose.yml # Change the command line on caddy to your domain & environment variables to your needs
docker-compose up -d
```
