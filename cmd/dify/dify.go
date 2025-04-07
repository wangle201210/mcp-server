package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/cloudwego/eino-ext/components/retriever/dify"
)

func getDifyTool() *protocol.Tool {
	datasetName := os.Getenv("DIFY_DATASET_NAME")
	return &protocol.Tool{
		Name:        "dify_retriever",
		Description: fmt.Sprintf("检索 %s 知识库", datasetName),
		InputSchema: protocol.InputSchema{
			Type: protocol.Object,
			Properties: map[string]interface{}{
				"query": map[string]string{
					"type":        "string",
					"description": "检索内容",
				},
			},
			Required: []string{"query"},
		},
	}
}

func difyHandler(request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	query, ok := request.Arguments["query"].(string)
	if !ok {
		return nil, errors.New("query must be a string")
	}
	APIKey := os.Getenv("DIFY_DATASET_API_KEY")
	Endpoint := os.Getenv("DIFY_ENDPOINT")
	DatasetID := os.Getenv("DIFY_DATASET_ID")
	ret, err := dify.NewRetriever(context.Background(), &dify.RetrieverConfig{
		APIKey:    APIKey,
		Endpoint:  Endpoint,
		DatasetID: DatasetID,
	})
	if err != nil {
		return nil, err
	}
	docs, err := ret.Retrieve(context.Background(), query)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(docs)
	if err != nil {
		return nil, err
	}
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: string(marshal),
			},
		},
	}, nil
}
