package main

import (
	"errors"
	"fmt"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/wangle201210/dt/sdk/time"
)

func getDtTimeTool() *protocol.Tool {
	return &protocol.Tool{
		Name:        "dt_time",
		Description: "时间戳转时间",
		InputSchema: protocol.InputSchema{
			Type: protocol.Object,
			Properties: map[string]interface{}{
				"timestamp": map[string]string{
					"type":        "number",
					"description": "秒级时间戳",
				},
			},
			Required: []string{"timestamp"},
		},
	}
}

func dtTimeHandler(request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	timestamp, ok := request.Arguments["timestamp"].(float64)
	if !ok {
		return nil, errors.New("timestamp must be a int")
	}
	ts := int64(timestamp)
	res, err := time.ToDate(ts)
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
