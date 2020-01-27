BIN:=./bin
SERVICE:=$(BIN)/service

build: $(SERVICE)

$(SERVICE):
	go get ./...
	go build -o $(SERVICE) cmd/service/main.go

.PHONY: run
run:
	mkdir -p data
	./bin/service

.PHONY: test
test:
	go test -v ./pkg/...
