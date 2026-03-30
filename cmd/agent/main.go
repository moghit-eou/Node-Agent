package main

import (
	"log"
	"os"

	"node-agent/internal/network"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("usage: go run main.go <port> !!!") // should avoid 1-1023 ports
		return
	}

	port := os.Args[1]
	network.StartServer(port)
}
