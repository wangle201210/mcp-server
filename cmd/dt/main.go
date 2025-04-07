package main

import (
	"fmt"

	"github.com/wangle201210/mcp-server/helper"
)

func main() {
	s := helper.NewServer("data_transfer")
	s.RegisterTool(getDtTimeTool(), dtTimeHandler)
	s.RegisterTool(getDtDecodeTool(), dtDecodeHandler)
	// 启动服务
	if err := s.Run(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
