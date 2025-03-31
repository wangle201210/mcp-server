package helper

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

func NewServer(name string) *server.MCPServer {
	s := server.NewMCPServer(
		name,
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithLogging(),
	)
	return s
}

func ServerRun(s *server.MCPServer) {
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
	// sseServer := server.NewSSEServer(s, server.WithBaseURL("http://localhost:8088"))
	// log.Printf("SSE server listening on :8088")
	// if err := sseServer.Start(":8088"); err != nil {
	// 	log.Fatalf("Server error: %v", err)
	// }
}
