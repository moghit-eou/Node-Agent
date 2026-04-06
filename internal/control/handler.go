package control

import (
	"fmt"

	"node-agent/internal/execution"
	"context"
)

func HandleJob(commandPayLoad string) string {

	if commandPayLoad == "" {
		return "Error : Empty command"
	}

	exec, err := execution.NewDockerExecutor("alpine")
	
	if err != nil {
    	return fmt.Sprintf("error: %v", err)
	}	
	defer exec.Close()

	result, err := exec.Run(context.Background(), commandPayLoad)
	
	if err != nil {
		return fmt.Sprintf("Something went wrong: %v", err)
	}

	response := fmt.Sprintf("Stdout : %s \nStderr: %s\nExistcode: %d",
		result.Stdout, result.Stderr, result.ExitCode)

	return response
}
