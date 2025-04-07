package helper

import (
	"log"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
)

func NewServer(name string) *server.Server {
	// 创建MCP服务器
	s, err := server.NewServer(
		transport.NewStdioServerTransport(),
		server.WithServerInfo(protocol.Implementation{
			Name:    name,
			Version: "1.0.0",
		}),
	)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	return s
}
