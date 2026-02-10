package main

import (
	"fmt"
	"context"
	"time"
//	"node-agent/playground" -> this is not allowed
)


var docker *DockerClient

func init() {
	var err error
	docker, err = NewDockerClient()
	if err != nil {
		fmt.Printf("Docker Error: %v\n", err)
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	result , err := docker.Run(ctx , "echo hello" )
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Stdout: %s\nStderr: %s", result.Stdout, result.Stderr)

	
}