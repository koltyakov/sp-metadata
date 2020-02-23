install:
	go get -u ./... && go mod tidy

format:
	gofmt -s -w .

get:
	go run ./cmd/get

gen:
	go run ./cmd/gen

publish:
	./scripts/publish.sh

build:
	go build -o bin/get ./cmd/get
	go build -o bin/gen ./cmd/gen