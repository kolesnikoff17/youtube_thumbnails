

run:
	docker-compose up --force-recreate --build -d

down:
	docker-compose down

build:
	go build -o app client/cmd/main.go