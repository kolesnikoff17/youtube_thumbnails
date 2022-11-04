all: run build

run:
	docker-compose up --build -d

down:
	docker-compose down

build:
	go build -o app client/cmd/main.go

test:
	go test -v --race ./... -coverprofile cover.out

cover.out:
	make test

report: cover.out
	go tool cover -html=cover.out