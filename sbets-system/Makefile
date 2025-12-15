.PHONY: build test run-sbets clean

build:
	go build -o bin/sbets cmd/sbets/main.go

test:
	go test ./...

run-sbets:
	CONVERTER_SERVICE_URL=http://localhost:8080 go run cmd/sbets/main.go

clean:
	rm -rf bin/