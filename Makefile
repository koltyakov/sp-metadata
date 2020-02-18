install:
	go get -u ./... && go mod tidy

format:
	gofmt -s -w .

get:
	go run ./cmd/get

publish:
	./scripts/publish.sh

build:
	go build -o bin/job ./cmd/get