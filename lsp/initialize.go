package lsp


type InitializeRequest struct {
	Request 
	Params InitilizeRequestParams `json:"params"`
}



type InitilizeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// ... there's tons more that goes here
}


type ClientInfo struct {
	Name string `json:"name"`
	Version string `json:"version"`
}



type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}


type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo ServerInfo `json:"serverInfo"`
}
	
type ServerCapabilities struct {}

type ServerInfo struct {
	Name string `json:"name"`
	Version string `json:"version"`
}


func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID: id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{},
			ServerInfo: ServerInfo{
				Name: "educatinal lsp",
				Version: "0.0.1",
			},
		},
	}
}