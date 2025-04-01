package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/wangle201210/dt/sdk/time"
)

func getDtTimeTool() mcp.Tool {
	return mcp.NewTool("dt_time",
		mcp.WithDescription("时间戳转时间"),
		mcp.WithNumber("timestamp",
			mcp.Required(),
			mcp.Description("秒级时间戳"),
		),
	)
}

func dtTimeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	timestamp, ok := request.Params.Arguments["timestamp"].(float64)
	if !ok {
		return nil, errors.New("timestamp must be a int")
	}
	ts := int64(timestamp)
	res, err := time.ToDate(ts)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DoDecrypt err: %+v", err))
	}
	return mcp.NewToolResultText(res), nil
}
