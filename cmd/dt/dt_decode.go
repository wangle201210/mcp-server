package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/wangle201210/dt/sdk/decode"
)

func getDtDecodeTool() mcp.Tool {
	return mcp.NewTool("dt/decrypt",
		mcp.WithDescription("解密数据"),
		mcp.WithString("requestId",
			mcp.Required(),
			mcp.Description("requestId是解密时使用的key"),
		),
		mcp.WithString("data",
			mcp.Required(),
			mcp.Description("data是需要被解密的数据"),
		),
	)
}

func dtDecodeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	key, ok := request.Params.Arguments["requestId"].(string)
	if !ok {
		return nil, errors.New("requestId must be a string")
	}
	data, ok := request.Params.Arguments["data"].(string)
	if !ok {
		return nil, errors.New("data must be a string")
	}
	res, err := decode.DoDecrypt(key, data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DoDecrypt err: %+v", err))
	}
	return mcp.NewToolResultText(res), nil
}
