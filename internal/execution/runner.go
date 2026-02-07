package execution

import (

	"os/exec"
	"bytes"

)
type Result struct 
{
	Stdout string 
	Stderr string 
	ExitCode int 
}

func runCommand(command string ) (*Result , error)
{

	cmd := exec.command("sh" , "-c" , command )

	var stdout , stderr bytes.buffer 

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	exitCode := 0 

	if err != nil {

	}


	return &Result
	{
		Stdout : stdout,
		Stderr : stderr,
		ExiCode : exitCode,
	} , nil 

}