package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/wangle201210/text2sql"
	"github.com/wangle201210/text2sql/eino"
)

func getText2sqlTool() *protocol.Tool {
	return &protocol.Tool{
		Name:        "text2sql",
		Description: "自然语言转数据库查询",
		InputSchema: protocol.InputSchema{
			Type: protocol.Object,
			Properties: map[string]interface{}{
				"question": map[string]string{
					"type":        "string",
					"description": "查询数据的问题",
				},
			},
			Required: []string{"question"},
		},
	}
}

func text2sqlHandler(request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	link := os.Getenv("link")
	if link == "" {
		return nil, fmt.Errorf("link is empty")
	}
	question, ok := request.Arguments["question"].(string)
	if !ok {
		return nil, errors.New("question must be a string")
	}
	cfg := &text2sql.Config{
		DbLink:    link,
		ShouldRun: true,
		Times:     3,
		Try:       3,
	}
	newEino, err := eino.NewEino(&openai.ChatModelConfig{
		APIKey:  "sk-******",
		BaseURL: "https://api.openai.com/v1",
		Model:   "gpt-4o-mini",
	})
	if err != nil {
		return nil, fmt.Errorf("eino err: %+v", err)
	}
	ts := text2sql.NewText2sql(cfg, newEino)
	sql, result, err := ts.Pretty(question)
	if err != nil {
		return nil, fmt.Errorf("text2sql err: %+v", err)
	}
	// 组装返回的数据
	res := fmt.Sprintf("sql: %s, result: %s", sql, result)
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: res,
			},
		},
	}, nil
}
