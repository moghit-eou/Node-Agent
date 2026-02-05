package network

import (
	"net"
	"bufio"
	"log"
	"fmt"
)

func StartServer(port string) {

	Listener , err := net.Listen ("tcp" , ":" + port)
	
	if err != nil {
		log.Println("Error starting server:", err)
		return 
	}

	log.Println("server is listening on port " + port + "")
	defer Listener.Close()

	for {
		conn , err := Listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue 
		}
		log.Println("new connection from", conn.RemoteAddr())
		go handleConnection(conn)
	}
} 

func handleConnection(conn net.Conn) error {
	defer conn.Close()
	
	// I still have to pay attention to DDOS
	// to do : set a timeout for the connection
 
	reader := bufio.NewReader(conn)
	message , err := reader.ReadString('\n')
	
	if err != nil {
		log.Printf("Error reading message %v", err)
		return err 
	}

	// chiwiwi

	response := fmt.Sprintf("Echo -> %s" , message) 

	_, err = conn.Write([]byte(response))
		// here basicaly converting String into bytes and send it back to the client
	if err != nil {
		log.Printf("sending response failed: %v", err)
	}
	
	log.Printf("Received message: %s", message)
	log.Println("Echoed message back to", conn.RemoteAddr())

	return nil	
		 
}