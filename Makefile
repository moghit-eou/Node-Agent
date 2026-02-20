APP := node-agent
CMD := ./cmd/agent
BIN := bin

.PHONY: run clean fmt tidy

run:
	go run $(CMD) 8080

fmt:
	go fmt ./...

tidy:
	go mod tidy



clean:
	rm -rf $(BIN)
