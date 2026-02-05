package main

import (
	
	"os"
	"node-agent/internal/network"
	"log"

)

func main() {
	
	if len(os.Args) < 2 {
		log.Println("usage: go run main.go <port> !!!") // avoid 1-1023 ports
		return 
	}
	
	port := os.Args[1]
	network.StartServer(port)
	  
} 


