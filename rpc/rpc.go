package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)

}

func DecodeMessage(msg []byte) (string, []byte, error) {

	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return "", nil, errors.New("Did not find separator")
	}

	// Conent-Length : <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))

	if err != nil {
		return "", nil, err
	}

	var BaseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &BaseMessage); err != nil {
		return "", nil, err
	}

	return BaseMessage.Method, content, nil

}


func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return 0, nil, nil
	}

	// Conent-Length : <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))

	if err != nil {
		return 0, nil, err
	}

	if (len(content) < contentLength) {
		return 0, nil, nil
	}

	totalLength := len(header) + len(content) + 4
	return totalLength, data[:totalLength], nil
}