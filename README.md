# mcp-server

## text2sql

### 安装
```bash
go install github.com/wangle201210/mcp-server/cmd/text2sql@latest
```

### 配置
```json
{
    "mcpServers": {
      "text2sql": {
        "descriptions": "自然语言转sql",
        "icons": "📁",
        "autoApprove": [
          "all"
        ],
        "type": "stdio",
        "command": "text2sql",
        "args": [],
        "env": {
          "link": "root:@tcp(127.0.0.1:3306)/goadmin?charset=utf8mb4&parseTime=True&loc=Local",
        }
      }
    }
}
```

## dify retriever

### 安装
```bash
go install github.com/wangle201210/mcp-server/cmd/dify@latest
```

### 配置
```json
{
    "mcpServers": {
      "dify_retriever": {
        "descriptions": "dify知识库检索",
        "icons": "📁",
        "autoApprove": [
          "all"
        ],
        "type": "stdio",
        "command": "dify",
        "args": [],
        "env":{
          "DIFY_DATASET_API_KEY": "dataset-***",
          "DIFY_ENDPOINT": "your DIFY_ENDPOINT",
          "DIFY_DATASET_ID": "your DIFY_DATASET_ID",
          "DIFY_DATASET_NAME": "deepchat 快捷键介绍"
        }
      }
    }
}
```


## 数据转换

### 安装
```bash
go install github.com/wangle201210/mcp-server/cmd/dt@latest
```

### 配置
```json
{
    "mcpServers": {
      "data_transfer": {
        "descriptions": "数据转换",
        "icons": "📁",
        "autoApprove": [
          "all"
        ],
        "type": "stdio",
        "command": "dt",
        "args": [],
        "env":{
        }
      }
    }
}
```
