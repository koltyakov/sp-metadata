install:
	go get -u ./... && go mod tidy

format:
	gofmt -s -w .

get:
	go run ./cmd/get

compare:
	go run ./cmd/compare

publish:
	./scripts/publish.sh

build:
	go build -o bin/get ./cmd/get
	go build -o bin/compare ./cmd/compare