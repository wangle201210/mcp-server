package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/wangle201210/text2sql"
	"github.com/wangle201210/text2sql/eino"
)

func getText2sqlTool() mcp.Tool {
	return mcp.NewTool("dt_texe2sql",
		mcp.WithDescription("数据库查询"),
		mcp.WithString("question",
			mcp.Required(),
			mcp.Description("查询数据的问题"),
		),
	)
}

func text2sqlHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if err := checkOpenEnv(); err != nil {
		return nil, err
	}
	link := os.Getenv("link")
	tryStr := os.Getenv("try")
	runStr := os.Getenv("run")
	timeStr := os.Getenv("times")
	if link == "" {
		return nil, fmt.Errorf("link is empty")
	}
	question, ok := request.Params.Arguments["question"].(string)
	if !ok {
		return nil, errors.New("question must be a string")
	}
	try, _ := strconv.Atoi(tryStr)
	run, _ := strconv.ParseBool(runStr)
	times, _ := strconv.Atoi(timeStr)
	if try == 0 {
		try = 3
	}
	if times == 0 {
		times = 3
	}
	cfg := &text2sql.Config{
		DbLink:    link,
		Try:       try,
		ShouldRun: run,
		Times:     times,
		// OnlyView:  true,
	}
	newEino, err := eino.NewEino(&openai.ChatModelConfig{
		APIKey:  os.Getenv("OPENAI_API_KEY"),
		BaseURL: os.Getenv("OPENAI_BASE_URL"),
		Model:   os.Getenv("OPENAI_MODEL_NAME"),
	})
	if err != nil {
		return nil, fmt.Errorf("eino err: %+v", err)
	}
	ts := text2sql.NewText2sql(cfg, newEino)
	sql, result, err := ts.Pretty(question)
	if err != nil {
		return nil, fmt.Errorf("text2sql err: %+v", err)
	}
	res := fmt.Sprintf("sql: %s, \nresult: %s", sql, result)
	return mcp.NewToolResultText(res), nil
}

func checkOpenEnv() error {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return errors.New("OPENAI_API_KEY is empty")
	}
	modelName := os.Getenv("OPENAI_MODEL_NAME")
	if modelName == "" {
		return errors.New("OPENAI_MODEL_NAME is empty")
	}
	baseUrl := os.Getenv("OPENAI_BASE_URL")
	if baseUrl == "" {
		return errors.New("OPENAI_BASE_URL is empty")
	}
	return nil
}
