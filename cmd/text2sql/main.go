package main

import (
	"fmt"

	"github.com/wangle201210/mcp-server/helper"
)

func main() {
	s := helper.NewServer("text2sql")
	// 注册工具和处理函数
	s.RegisterTool(getText2sqlTool(), text2sqlHandler)
	// 启动服务
	if err := s.Run(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
