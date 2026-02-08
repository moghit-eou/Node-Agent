package network


type Request struct {
	Type string `json:"type"`
	Payload string `json:"payload"`
}

type Response struct { 
	Status string `json:"status"`
	Message string `json:"message"` // json name is used during testing calls
									// Message in GO , message in json
									// it's a mapping layer
}

