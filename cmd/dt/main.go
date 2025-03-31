package main

import (
	"github.com/wangle201210/mcp-server/helper"
)

func main() {
	s := helper.NewServer("data_transfer")
	s.AddTool(getDtTimeTool(), dtTimeHandler)
	s.AddTool(getDtDecodeTool(), dtDecodeHandler)
	helper.ServerRun(s)
}
