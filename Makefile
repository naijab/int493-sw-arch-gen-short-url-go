install:
	go mod tidy

install-docker:
	sudo apt update
	sudo apt install git -y
	sudo apt install apt-transport-https ca-certificates curl software-properties-common
	curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
	sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
	sudo apt update
	sudo apt install docker-ce -y
	sudo usermod -aG docker ${USER}
	su - ${USER}

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