package main
// go run main.go 
// nc localhost <port> < request.json
import (
	
	"os"
	//"node-agent/internal/network"
	//"log"
	
	"node-agent/internal/execution"
	"node-agent/internal/control"
	"os/exec"
	"fmt"
	"bytes"

)

func main() {
	
	if len(os.Args) < 2 {
//		log.Println("usage: go run main.go <port> !!!") // I should avoid 1-1023 ports
//		return 
	}

	//We use sh -c to bridge the gap. 
	//It takes your single string and executes it like a real terminal.
	command := "ls -al"
	cmd := exec.Command("sh", "-c", command) 
	var stderr,stdout bytes.Buffer

	// 3. Connect our buckets to the command's pipes
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 4. Run it!
	err := cmd.Run()
	fmt.Printf("stdout :\n%s\n --- stder -- \n %s\n", stdout.String(), stderr.String()	)

	fmt.Printf("error : %v\n", err)
	
	/* stdout :
		total 0
	drwxrwxrwx 1 elmmo elmmo 4096 Feb  7 23:17 .
	drwxrwxrwx 1 elmmo elmmo 4096 Feb  5 14:41 ..
	-rwxrwxrwx 1 elmmo elmmo  285 Feb  7 23:17 main.go
	-rwxrwxrwx 1 elmmo elmmo  683 Feb  7 23:19 main2.go
	-rwxrwxrwx 1 elmmo elmmo  196 Feb  7 22:10 request.json 
	*/

//	port := os.Args[1]
//	network.StartServer(port)

	result , err := execution.RunCommand("ls -al")	

	fmt.Println("\n-------\n-stdout-\n------------\n")
	fmt.Println("Stdout :")
	fmt.Println(result.Stdout)

	fmt.Println("Stderr :")
	fmt.Println(result.Stderr)

	fmt.Println("ExitCode :")
	fmt.Println(result.ExitCode)

	fmt.Println("\n error output: \n")
	fmt.Println("-", err , "-")

	fmt.Println("\n\n----------------\n\n")



	fmt.Println("%s" , control.HandleJob("ls -al"))
} 


