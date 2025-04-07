package main

import (
	"fmt"

	"github.com/wangle201210/mcp-server/helper"
)

func main() {
	s := helper.NewServer("text2sql")
	s.RegisterTool(getDifyTool(), difyHandler)
	// {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"dify_retriever","arguments":{"query":"test"}}}
	// 启动服务
	if err := s.Run(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
