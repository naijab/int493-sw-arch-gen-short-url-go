install:
	go mod tidy

# Must install air for hot reload
# https://github.com/cosmtrek/air#installation
dev:
	air
