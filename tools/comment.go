package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetCommentsTool 获取帖子评论工具
func GetCommentsTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "campus_market_get_comments",
		Description: "获取指定帖子的评论列表",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"thread_id": map[string]interface{}{
					"type":        "string",
					"description": "帖子ID",
				},
			},
			"required": []string{"thread_id"},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			ThreadID string `json:"thread_id"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		if args.ThreadID == "" {
			return nil, fmt.Errorf("thread_id is required")
		}

		client := getClient()
		comments, err := client.GetComment(args.ThreadID)
		if err != nil {
			return nil, err
		}

		var result strings.Builder
		for _, comment := range *comments {
			result.WriteString(comment.FriendlyText())
			result.WriteString("\n\n")
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: result.String()},
			},
		}, nil
	})
}

// PostCommentTool 发表评论工具
func PostCommentTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "campus_market_post_comment",
		Description: "在指定帖子下发表评论",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"thread_id": map[string]interface{}{
					"type":        "string",
					"description": "帖子ID",
				},
				"content": map[string]interface{}{
					"type":        "string",
					"description": "评论内容",
				},
				"reply_comment_id": map[string]interface{}{
					"type":        "string",
					"description": "回复的评论ID（可为空，用于回复特定评论）",
					"default":     "0",
				},
				"root_comment_id": map[string]interface{}{
					"type":        "string",
					"description": "根评论ID（用于楼中楼回复，可为空）",
					"default":     "0",
				},
				"use_anon": map[string]interface{}{
					"type":        "number",
					"description": "是否匿名（0:否, 1:是）",
					"default":     0,
				},
			},
			"required": []string{"thread_id", "content"},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			ThreadID       string `json:"thread_id"`
			Content        string `json:"content"`
			ReplyCommentID string `json:"reply_comment_id"`
			RootCommentID  string `json:"root_comment_id"`
			UseAnon        int    `json:"use_anon"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		if args.ThreadID == "" {
			return nil, fmt.Errorf("thread_id is required")
		}

		if args.Content == "" {
			return nil, fmt.Errorf("content is required")
		}

		if args.ReplyCommentID == "" {
			args.ReplyCommentID = "0"
		}

		if args.RootCommentID == "" {
			args.RootCommentID = "0"
		}

		client := getClient()
		success, err := client.PostComment(args.ThreadID, args.Content, args.ReplyCommentID, args.RootCommentID, args.UseAnon)
		if err != nil {
			return nil, err
		}

		text := "评论发表成功"
		if !success {
			text = "评论发表失败"
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: text},
			},
		}, nil
	})
}

// DeleteCommentTool 删除评论工具
func DeleteCommentTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "campus_market_delete_comment",
		Description: "删除指定的评论",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"thread_id": map[string]interface{}{
					"type":        "string",
					"description": "帖子ID",
				},
				"comment_id": map[string]interface{}{
					"type":        "string",
					"description": "评论ID",
				},
			},
			"required": []string{"thread_id", "comment_id"},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			ThreadID  string `json:"thread_id"`
			CommentID string `json:"comment_id"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		if args.ThreadID == "" {
			return nil, fmt.Errorf("thread_id is required")
		}

		if args.CommentID == "" {
			return nil, fmt.Errorf("comment_id is required")
		}

		client := getClient()
		success, err := client.DeleteComment(args.ThreadID, args.CommentID)
		if err != nil {
			return nil, err
		}

		text := "评论删除成功"
		if !success {
			text = "评论删除失败"
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: text},
			},
		}, nil
	})
}

// LikeCommentTool 点赞评论工具
func LikeCommentTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "campus_market_like_comment",
		Description: "点赞指定的评论",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"thread_id": map[string]interface{}{
					"type":        "string",
					"description": "帖子ID",
				},
				"comment_id": map[string]interface{}{
					"type":        "string",
					"description": "评论ID",
				},
			},
			"required": []string{"thread_id", "comment_id"},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			ThreadID  string `json:"thread_id"`
			CommentID string `json:"comment_id"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		if args.ThreadID == "" {
			return nil, fmt.Errorf("thread_id is required")
		}

		if args.CommentID == "" {
			return nil, fmt.Errorf("comment_id is required")
		}

		client := getClient()
		success, err := client.LikeComment(args.ThreadID, args.CommentID)
		if err != nil {
			return nil, err
		}

		text := "点赞成功"
		if !success {
			text = "点赞失败"
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: text},
			},
		}, nil
	})
}

// UnlikeCommentTool 取消点赞评论工具
func UnlikeCommentTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "campus_market_unlike_comment",
		Description: "取消点赞指定的评论",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"thread_id": map[string]interface{}{
					"type":        "string",
					"description": "帖子ID",
				},
				"comment_id": map[string]interface{}{
					"type":        "string",
					"description": "评论ID",
				},
			},
			"required": []string{"thread_id", "comment_id"},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			ThreadID  string `json:"thread_id"`
			CommentID string `json:"comment_id"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		if args.ThreadID == "" {
			return nil, fmt.Errorf("thread_id is required")
		}

		if args.CommentID == "" {
			return nil, fmt.Errorf("comment_id is required")
		}

		client := getClient()
		success, err := client.UnLikeComment(args.ThreadID, args.CommentID)
		if err != nil {
			return nil, err
		}

		text := "取消点赞成功"
		if !success {
			text = "取消点赞失败"
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: text},
			},
		}, nil
	})
}
