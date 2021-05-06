install:
	go mod tidy

# Must install air for hot reload
# https://github.com/cosmtrek/air#installation
dev:
	air

build-app:
	docker-compose up -d --build shortener_api

start-app:
	docker-compose up -d shortener_api

stop-app:
	docker-compose stop shortener_api

start-redis:
	docker-compose up -d redis-cluster

stop-redis:
	docker-compose stop redis-cluster

status:
	docker-compose ps -a

down:
	docker-compose down --remove-orphans