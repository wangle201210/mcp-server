package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"wanna`s mcp server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithLogging(),
	)

	// Add tool handler
	s.AddTool(getDtTimeTool(), dtTimeHandler)
	s.AddTool(getDtDecodeTool(), dtDecodeHandler)
	s.AddTool(getCarTool(), carHandler)
	// // Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
	// sseServer := server.NewSSEServer(s, server.WithBaseURL("http://localhost:8088"))
	// log.Printf("SSE server listening on :8088")
	// if err := sseServer.Start(":8088"); err != nil {
	// 	log.Fatalf("Server error: %v", err)
	// }
}
