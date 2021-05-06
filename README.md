# URL Shortener Project

> URL Shortener API's with Golang (Fiber) and Redis

## Must installed before run
- `sudo apt install make` for installing make
- `make install-docker` for installing docker
- install docker-compose with
  
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.28.5/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

## How to run
1. copy `.env.example` to `.env` and edit variables
2. use command `go run main.go`

## Production
1. run `make start-app` and `make start-redis`