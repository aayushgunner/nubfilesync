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
	if len(os.Args) < 2 {
		printUsage()
		return
	}
}
func printUsage() {
	fmt.Println("File Synchronization tool")
	fmt.Println("Usage:")
	fmt.Println(" go run fileshare.go server										- Start as server")
	fmt.Println(" go run filesync.go pull <ipv6_address>        - Pull file from server")
	fmt.Println(" go run filesync.go push <ipv6_address>        - Push file to server")
	fmt.Println("")
}
