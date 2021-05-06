install:
	go mod tidy

# Must install air for hot reload
# https://github.com/cosmtrek/air#installation
dev:
	air

start-app:
	docker-compose up -d shortener_api

stop-app:
	docker-compose down shortener_api --remove-orphans

start-redis:
	docker-compose up -d redis-cluster

stop-redis:
	docker-compose down redis-cluster --remove-orphans

status:
	docker-compose ps -a