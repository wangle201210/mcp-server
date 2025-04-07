package main

import (
	"errors"
	"fmt"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/wangle201210/dt/sdk/decode"
)

func getDtDecodeTool() *protocol.Tool {
	return &protocol.Tool{
		Name:        "dt_decrypt",
		Description: "解密数据",
		InputSchema: protocol.InputSchema{
			Type: protocol.Object,
			Properties: map[string]interface{}{
				"requestId": map[string]string{
					"type":        "string",
					"description": "requestId是解密时使用的key",
				},
				"data": map[string]string{
					"type":        "string",
					"description": "data是需要被解密的数据",
				},
			},
			Required: []string{"requestId", "data"},
		},
	}
}

func dtDecodeHandler(request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	key, ok := request.Arguments["requestId"].(string)
	if !ok {
		return nil, errors.New("requestId must be a string")
	}
	data, ok := request.Arguments["data"].(string)
	if !ok {
		return nil, errors.New("data must be a string")
	}
	res, err := decode.DoDecrypt(key, data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DoDecrypt err: %+v", err))
	}
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: res,
			},
		},
	}, nil
}
