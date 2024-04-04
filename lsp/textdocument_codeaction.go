package lsp


type CodeActionRequest struct {
	Request
	Params TextDocumentCodeActionParams `json:"params"`
}

type TextDocumentCodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range Range `json:"range"`
	Context CodeActionContext `json:"context"`
}

type CodeActionContext struct {
	// add fields for code aciton context as needed
}

type CodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}



type CodeAction struct {
	Title string `json:"title"`
	Edit *WorkspaceEdit `json:"edit"`
	Command *Command `json:"command"`
}



type Command struct {
	Title     string        `json:"title"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}
	
