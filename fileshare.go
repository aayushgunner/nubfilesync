package main

import (
	"fmt"
	"log"
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
		runServer()

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

func runServer() {

	if info, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal("Error creating file:", err)
		}
		defer file.Close() //defer used i don't know for now but might be useful later just may be safe for future purpose'
		fmt.Printf("Created %s\n", fileName)
	} else {
		fmt.Printf("File: %s\n", info.Name())
		fmt.Printf("Size: %d bytes\n", info.Size())
		fmt.Printf("Modified: %s\n", info.ModTime())
		fmt.Printf("Permissions: %s\n", info.Mode())
	}
	return
}
