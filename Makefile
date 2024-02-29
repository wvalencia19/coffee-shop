run:
	go build -o bin/coffe-shop ./cmd/
	./bin/coffe-shop
test:
	go test -race -v ./...
	
.PHONY: run	