install:
	go get -u ./... && go mod tidy

format:
	gofmt -s -w .

get:
	go run ./cmd/get

publish:
	make get && ./scripts/publish.sh