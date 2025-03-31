package main

import (
	"github.com/wangle201210/mcp-server/helper"
)

func main() {
	s := helper.NewServer("text2sql")
	s.AddTool(getDifyTool(), difyHandler)
	// {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"dify/retriever","arguments":{"query":"test"}}}
	helper.ServerRun(s)
}
