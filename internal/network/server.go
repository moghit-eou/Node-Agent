package network

import (
	"net"
	"log"
 	"encoding/json"
	"io"
)

func StartServer(port string) {

	Listener , err := net.Listen ("tcp" , ":" + port)
	
	if err != nil {
		log.Println("Error starting server:", err)
		return 
	}

	log.Println("server is listening on port " + port + "")
	defer Listener.Close() // runs after function exist

	for {
		
		// conn : tcp socket 
		// Listener : door to the socket ( bound to a port )

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

	
	// I still have to pay attention to DDOS attcks 
	// to do : set a timeout for the connection
	

	// conn is a stream of bytes
	// I need to decode it into a Request struct and encode the Response struct back to the client
	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)
	
	for 
	{
				
		request := Request{}
		err := decoder.Decode(&request)
		
		if err == io.EOF {
			log.Printf("client %s disconnected", conn.RemoteAddr())
			return nil 
		}

		if err != nil {
			log.Printf("decoding request failed: %v", err)
			return err 
		}
		
		log.Printf("Received request : Type = %s, Payload = %s", request.Type, request.Payload)

		var response Response

		switch request.Type {
			
			case "health-check":
				response = Response{
					Status : "ok",
					Message  : request.Payload,
				}
			
			case "job":
				response = Response {
					Status : "pending",
					Message : "Job is being processed " + request.Payload,
				}	
				
			default:
				response = Response {
					Status : "error",
					Message : "Unknown request type",
				}	
			
		}
	 

		err = encoder.Encode(response)

		if err != nil {
			log.Printf("sending response failed: %v", err)
		}

	}
 	return nil	
		 
}