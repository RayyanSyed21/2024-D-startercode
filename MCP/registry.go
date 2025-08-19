package main

import (
	"github.com/softeng-starter-code-api/mcp-server/config"
	"github.com/softeng-starter-code-api/mcp-server/models"
	tools_general "github.com/softeng-starter-code-api/mcp-server/tools/general"
	tools_api "github.com/softeng-starter-code-api/mcp-server/tools/api"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_general.CreateGetTool(cfg),
		tools_general.CreatePostTool(cfg),
		tools_api.CreateGet_api_high_scoreTool(cfg),
		tools_api.CreatePost_api_high_scoreTool(cfg),
	}
}
