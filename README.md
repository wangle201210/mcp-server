# mcp-server
## 安装

```bash
go install github.com/wangle201210/mcp-server
```

## 配置
```json
{
    "mcpServers": {
      "dt-stdio": {
        "descriptions": "自然语言转sql",
        "icons": "📁",
        "autoApprove": [
          "all"
        ],
        "type": "stdio",
        "command": "mcp-server",
        "args": [],
        "env": {
          "link": "root:@tcp(127.0.0.1:3306)/goadmin?charset=utf8mb4&parseTime=True&loc=Local",
          "run": "true",
          "OPENAI_API_KEY": "sk-******",
          "OPENAI_MODEL_NAME": "gpt-4o-mini",
          "OPENAI_BASE_URL": "https://api.openai.com/v1"
        }
      }
    }
}
```