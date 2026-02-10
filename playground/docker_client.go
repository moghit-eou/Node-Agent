package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
)
// client != Client 
//client (lowercase) is a package (folder / namespace)
//Client (Capital C) is an exported struct type defined inside that package

//type Client struct {
	// contains filtered or unexported fields
	// definition 
	// from the documentaton
//}

 type DockerClient struct {
	cli *client.Client // to keep avoiding using client.client
}

func NewDockerClient() (*DockerClient, error) {
 	cli, err := client.NewClientWithOpts(client.FromEnv, 
		client.WithAPIVersionNegotiation() )  
		//client.FromEnv : looking for path of docker.sock and variables in geneeral
		//initializes a new API client 
	
		if err != nil {
		return nil, 
		fmt.Errorf("failed to create docker client: %w", err)
	}

	return &DockerClient{cli: cli}, nil
}

// Ping checks if Docker is actually alive and responding
//func (cli *Client) Ping(ctx context.Context) (types.Ping, error)
// why context -> it's documentation issue
func (d *DockerClient) Ping(ctx context.Context) error {
	_, err := d.cli.Ping(ctx)
	return err
}