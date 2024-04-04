package lsp 


type DidChangeNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
	TextDocument VersionTextDocumentIdentifier `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}


type TextDocumentContentChangeEvent struct {
	// the new text of the whole document 
	Text string `json:"text"`
}	
	
