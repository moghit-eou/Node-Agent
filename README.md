# Node-Agent

node-agent/
├── cmd/
│   └── agent/
│       └── main.go            # Entry Point (Config & Startup)
│
├── internal/
│   ├── network/               # LAYER 1: The Door
│   │   ├── server.go          # TCP Listener 
│   │   └── protocol.go        # JSON Structs 
│   │
│   ├── control/               #
│   │   └── handler.go         # Logic 
│   │
│   └── execution/             
│       └── runner.go          # Docker/Process Execution
│
├── go.mod
├── .gitignore
└── README.md