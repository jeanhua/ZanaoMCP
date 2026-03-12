package tools

import (
	"context"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetMessagesTool 获取用户消息工具
func GetMessagesTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "campus_market_get_messages",
		Description: "获取当前用户的消息列表",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
			"required":   []string{""},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := getClient()
		messages, err := client.GetMessage()
		if err != nil {
			return nil, err
		}

		var result strings.Builder
		for _, msg := range *messages {
			result.WriteString(msg.FriendlyText())
			result.WriteString("\n\n")
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: result.String()},
			},
		}, nil
	})
}

// GetUserInfoTool 获取用户信息工具
func GetUserInfoTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "campus_market_get_user_info",
		Description: "获取当前登录用户的信息",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
			"required":   []string{""},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := getClient()
		userInfo, err := client.GetUserInfo()
		if err != nil {
			return nil, err
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: userInfo.FriendlyText()},
			},
		}, nil
	})
}

// GetCategoriesTool 获取分类列表工具
func GetCategoriesTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "campus_market_get_categories",
		Description: "获取集市帖子的分类列表",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
			"required":   []string{""},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := getClient()
		categories, err := client.GetCategory()
		if err != nil {
			return nil, err
		}

		var result strings.Builder
		for _, cate := range *categories {
			result.WriteString(cate.FriendlyText())
			result.WriteString("\n")
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: result.String()},
			},
		}, nil
	})
}
