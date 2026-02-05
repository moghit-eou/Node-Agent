package main

import (
	
	"os"
	"fmt"
	"node-agent/internal/network"

)

func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port> !!!") // avoid 1-1023 ports
		return 
	}
	
	port := os.Args[1]
	network.StartServer(port)
	  
} 

