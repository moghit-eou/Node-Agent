package network

import (
	"fmt"
	"net"
)

func StartServer(port string) {

	Listener , err := net.Listen ("tcp" , ":" + port)
	
	if err != nil {
		fmt.Println("Error starting server:", err)
		return 
	}

	fmt.Println("server is listening on port " + port + "")
	defer Listener.Close()

	for {
		conn , err := Listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue 
		}
		fmt.Println("new connection from", conn.RemoteAddr())
		go handleConnection(conn)
	}
} 

func handleConnection(conn net.Conn) {
	conn.Write([]byte("RECEIVED \n"))
}