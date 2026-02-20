APP := node-agent
CMD := ./cmd/agent
BIN := bin

.PHONY: run  fmt tidy

run:
	go run $(CMD) 8080

fmt:
	go fmt ./...

tidy:
	go mod tidy


 
