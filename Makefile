include .env


.PHONY: start build run lint clean doc dev docker-build docker-image docker-image-staging docker-image-dev start-docker-dev stop-docker-dev


default: build

start:
	@GIN_MODE=release ./bin/$(API_NAME) 

build:
	@go build -o bin/$(API_NAME) 

run:
	@./bin/$(API_NAME) 

lint:
	@golint

clean:
	@go clean && rm -rf ./bin/$(API_NAME) && rm -f gin-bin

doc:
	godoc -http=:6060 -index

dev:
	@gin -a 8080 -p 3030 run main.go

docker-build: clean
	@docker-compose -f docker/development/docker-compose.yml run --rm api make build

docker-image: docker-build 
	@docker build -t $(API_NAME):latest .

docker-image-staging: docker-build 
	@docker build -t $(API_NAME):staging .

docker-image-dev:
	@docker-compose -f docker/development/docker-compose.yml run --rm api dep ensure -v

start-docker-dev:
	@docker-compose up -d

stop-docker-dev:
	@docker-compose down
