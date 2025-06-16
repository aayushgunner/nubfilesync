package main

import (
	"fmt"
	"os"
)

const (
	fileName = "new.txt"
	port     = ":8000"
)

func main() {
	// fmt.Println("the number of arguments is ", len(os.Args))
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	var command string
	command = os.Args[1]
	fmt.Printf("The command is %s\n", command) //found out that \n is needed at the end because fmt.Printf doesn't have a newline at the end
	commandChecker(command)

}

func printUsage() {
	fmt.Println("File Synchronization tool")
	fmt.Println("Usage:")
	fmt.Println(" go run fileshare.go server                    - Start as server")
	fmt.Println(" go run fileshare.go pull <ipv6_address>        - Pull file from server")
	fmt.Println(" go run fileshare.go push <ipv6_address>        - Push file to server")
	fmt.Println("")
}

func commandChecker(command string) {

	switch command {
	case "server":
		fmt.Println("Server is running")
		// runServer()

	case "pull":
		if len(os.Args) < 3 {
			fmt.Println("Error: IPv6 address required ")
			fmt.Println("Usage: go run fileshare.go pull <ipv6_address>")
			return
		}
		fmt.Println("Pulling is working")

	case "push":
		if len(os.Args) < 3 {
			fmt.Println("Error: IPv6 address required ")
			fmt.Println("Usage: go run fileshare.go push <ipv6_address>")
			return
		}
		fmt.Println("Pushing is working")

	default:

		printUsage()
	}

}
