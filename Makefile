APP := node-agent
CMD := ./cmd/agent
BIN := bin

PORT ?= 8080

.PHONY: all build run clean fmt tidy

all:
	@echo "=>(all) hitting the start..."
	clean fmt tidy build

build:
	@echo "=> Building the production binary..."
	go build -o $(BIN)/$(APP) $(CMD)

run: build
	@echo "=> Starting the compiled node-agent..."
	./$(BIN)/$(APP) $(PORT)

fmt:
	go fmt ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BIN)

.PHONY: test-request

test-request:
		@echo "=> Starting the compiled node-agent..."
		nc localhost $(PORT) < $(CMD)/request.json