package main

import (
	"bufio"
	"fmt"
	"log"
	"lsp/rpc"
	"lsp/lsp"
	"os"
	"encoding/json"
)


func main() {
	logger := getLogger("/home/amz/lsp/log.txt")
	logger.Println("Hey, I started!")
	fmt.Println("Hello, World!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
 
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println("got an error: %s", err)
			continue 
		}

		handleMessage(logger,method, contents)
	
}

}

func handleMessage(logger *log.Logger,method string,  content []byte) {
	logger.Printf("Received msg with method : %s ", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("Hey, we couldn't parse this : %s", err)
		}

		logger.Printf("Connected to %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)
		// hey ... let's reply!
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout 
		writer.Write([]byte(reply))
		logger.Printf("Sent the reply")
}


}

func getLogger(filename string) (*log.Logger) {
	logfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("hey, you didn't give me a good file")
	}
	return log.New(logfile, "[lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
	






