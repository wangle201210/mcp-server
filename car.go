package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func getCarTool() mcp.Tool {
	return mcp.NewTool("sicau/car",
		mcp.WithDescription("车辆进校预约"),
		mcp.WithString("car_number",
			mcp.Required(),
			mcp.Description("车牌号"),
		),
		mcp.WithString("time",
			mcp.Required(),
			mcp.Description("预约时间"),
		),
		mcp.WithString("username",
			mcp.Required(),
			mcp.Description("预约人"),
		),
	)
}

func carHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	car_number, ok := request.Params.Arguments["car_number"].(string)
	if !ok {
		return nil, errors.New("car_number must be a string")
	}
	time, ok := request.Params.Arguments["time"].(string)
	if !ok {
		return nil, errors.New("time must be a string")
	}
	username, ok := request.Params.Arguments["username"].(string)
	if !ok {
		return nil, errors.New("username must be a string")
	}

	return mcp.NewToolResultText(fmt.Sprintf("%s 预约了 %s 在 %s 到校的车辆进校申请", username, car_number, time)), nil
}
