package main

import (
	"github.com/wangle201210/mcp-server/helper"
)

func main() {
	s := helper.NewServer("text2sql")
	s.AddTool(getText2sqlTool(), text2sqlHandler)
	// {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"dt/texe2sql","arguments":{"question":"人员总数"}}}
	helper.ServerRun(s)
}
