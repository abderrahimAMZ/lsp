package lsp 



type Request struct {
	RPC string `json:"rpc"`
	ID int `json:"id"`
	method string `json:"method"`

	// we will just specify the type of the params in all the request types 
	// params 
}

type Response struct {
	RPC string `json:"rpc"`
	ID int `json:"id"`

	//result string `json:"result"`
	//error string `json:"error"`

	// we will just specify the type of the params in all the request types 
	// params 
}


type Notification struct {
	RPC string `json:"rpc"`
	method string `json:"method"`
}

