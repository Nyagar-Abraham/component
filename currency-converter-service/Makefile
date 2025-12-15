.PHONY: build test run-service clean

build:
	go build -o bin/converter-service cmd/service/main.go

test:
	go test ./...

run-service:
	go run cmd/service/main.go

clean:
	rm -rf bin/