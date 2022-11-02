

run:
	docker-compose up --force-recreate --build -d

down:
	docker-compose down

build:
	go build -o client ./client/cmd/main.go