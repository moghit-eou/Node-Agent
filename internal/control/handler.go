package control

import (
	"node-agent/internal/execution"
	"fmt"
)

func HandleJob(commandPayLoad string) string {

	if commandPayLoad == "" {
		return "Error : Empty command"
	}

	result , err := execution.RunCommand(commandPayLoad)

	if err != nil {
		return fmt.Sprintf("Something went wrong: %v" , err )

	}

	response := fmt.Sprintf("Stdout : %s \nStderr: %s\nExistcode: %d",
				result.Stdout , result.Stderr , result.ExitCode)

	return response 

}