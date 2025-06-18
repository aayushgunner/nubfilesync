package main

//package imports
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// some constants defined
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

// notify the users of the command and their usage
func printUsage() {
	fmt.Println("File Synchronization tool")
	fmt.Println("Usage:")
	fmt.Println(" go run fileshare.go server                    - Start as server")
	fmt.Println(" go run fileshare.go pull <ipv6_address>        - Pull file from server")
	fmt.Println(" go run fileshare.go push <ipv6_address>        - Push file to server")
	fmt.Println("")
}

// checks the command of the users if they are valid for our program or not
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

// function to run the server
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

	fmt.Println("=== File Sync Server ===")
	fmt.Printf("File: %s\n", fileName)
	fmt.Printf("Port :%s\n", port)
	fmt.Printf("Server starting on port %s ..\n", port)
	// http.HandleFunc("/file", handleFile)
	// http.HandleFunc("/status", handleStatus)
	// http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(port, nil))

}

func handleFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		content, err := os.ReadFile(fileName)
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(content)
		fmt.Printf("File sent to client (%d bytes)\n", len(content))

	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		err = os.WriteFile(fileName, body, 0644)
		if err != nil {
			http.Error(w, "Error writing file", http.StatusInternalServerError)
			fmt.Printf("Error writing file: %v\n", err)
			return
		}
		fmt.Printf("File updated by client (%d bytes)\n", len(body))
		fmt.Fprintf(w, "File uploaded successfully")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}

}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	info, err := os.Stat(fileName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "File: %s\nSize: %d bytes\nModified: %s\n",
		fileName, info.Size(), info.ModTime().Format("2006-01-02 15:04:05"))

}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "File Sync Server Running\nUse /file endpoint for sync operations\nUse /status for file info\n")

}
