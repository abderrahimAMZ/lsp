package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"lsp/analysis"
	"lsp/lsp"
	"lsp/rpc"
	"os"
	"io"
)


func main() {
	logger := getLogger("/home/amz/lsp/log.txt")
	logger.Println("Hey, I started!")
	fmt.Println("Hello, World!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
 
	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println("got an error: %s", err)
			continue 
		}

		handleMessage(logger,writer, state, method, contents)
	
}

}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string,  content []byte) {
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
		writeResponse(writer, msg)
		logger.Printf("Sent the reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/didOpen : %s", err)
		}
		logger.Printf(" %s",
			request.Params.TextDocument.URI)
		// hey ... let's reply!
		logger.Printf(" %s",
			request.Params.TextDocument.Text)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	case "textDocument/didChange":
		var request lsp.DidChangeNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/didChange : %s", err)
		}
		logger.Printf("Changed: %s", request.Params.TextDocument.URI)
	for _, change := range request.Params.ContentChanges {
		state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest 
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/hover : %s", err)
		}


		// Create a response 
		response := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		writeResponse(writer, response)
	case "textDocument/definition":
		var request lsp.DefinitionRequest 
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/definition : %s", err)
		}


		// Create a response 
		response := state.Definition(request.ID, request.Params.TextDocument.URI, request.Params.Position)

		writeResponse(writer, response)

	case "textDocument/codeAction":
		var request lsp.CodeActionRequest 
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/definition : %s", err)
		}

		response := state.TextDocumentCodeAction(request.ID, request.Params.TextDocument.URI)

		writeResponse(writer, response)
	case "textDocument/completion":
		var request lsp.CompletionRequest 
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("textDocument/definition : %s", err)
		}

		response := state.TextDocumentCompletion(request.ID, request.Params.TextDocument.URI)

		writeResponse(writer, response)
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
	
func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}





