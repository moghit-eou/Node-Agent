package network


type Request struct {
	Type string `json:"type"`
	Payload string `json:"payload"`
}


type Response struct { 
	Status string `json:"status"`
	Message string `json:"message"`

}


