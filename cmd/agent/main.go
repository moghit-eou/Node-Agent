package main

import (
	
	"os"
	"fmt"
	"net"
)

func main() {
	
	if len(os.Args) < 2 {
		println("Usage: go run main.go <port>")
		return 
	}
	
	port := os.Args[1]
	fmt.Println("Port:", port)

	listener , err := net.Listen ("tcp" , ":" + port)
	
	if err != nil {
		fmt.Println("Error starting server:", err)
		return 
	}

	fmt.Println("Server listening on port", listener.Addr().String())
} 