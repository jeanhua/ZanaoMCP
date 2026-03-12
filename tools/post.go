package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ListPostsTool 获取帖子列表工具
func ListPostsTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "list_posts",
		Description: "获取集市帖子列表，支持分页",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"from_time": map[string]interface{}{
					"type":        "string",
					"description": "起始时间戳，用于分页，设置为0表示从最新开始获取",
					"default":     "0",
				},
			},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			FromTime string `json:"from_time"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		fromTime := args.FromTime
		if fromTime == "" {
			fromTime = "0"
		}

		client := getClient()
		posts, err := client.GetPost(fromTime)
		if err != nil {
			return nil, err
		}

		var result strings.Builder
		for _, post := range *posts {
			result.WriteString(post.FriendlyText())
			result.WriteString("\n\n")
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: result.String()},
			},
		}, nil
	})
}

// HotPostsTool 获取热门帖子工具
func HotPostsTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "hot_posts",
		Description: "获取集市热门帖子列表",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := getClient()
		posts, err := client.GetHot()
		if err != nil {
			return nil, err
		}

		var result strings.Builder
		for _, post := range *posts {
			result.WriteString(post.FriendlyText())
			result.WriteString("\n\n")
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: result.String()},
			},
		}, nil
	})
}

// SearchPostsTool 搜索帖子工具
func SearchPostsTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "search_posts",
		Description: "在集市中搜索帖子（当前分类）",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"keyword": map[string]interface{}{
					"type":        "string",
					"description": "搜索关键词",
				},
				"page": map[string]interface{}{
					"type":        "number",
					"description": "页码",
					"default":     1,
				},
			},
			"required": []string{"keyword"},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			Keyword string `json:"keyword"`
			Page    int    `json:"page"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		if args.Keyword == "" {
			return nil, fmt.Errorf("keyword is required")
		}

		if args.Page == 0 {
			args.Page = 1
		}

		client := getClient()
		posts, err := client.Search(args.Keyword, args.Page)
		if err != nil {
			return nil, err
		}

		var result strings.Builder
		for _, post := range *posts {
			result.WriteString(post.FriendlyText())
			result.WriteString("\n\n")
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: result.String()},
			},
		}, nil
	})
}

// SearchHistoryPostsTool 搜索历史帖子工具
func SearchHistoryPostsTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "search_history_posts",
		Description: "在集市中搜索历史帖子，支持时间范围筛选",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"keyword": map[string]interface{}{
					"type":        "string",
					"description": "搜索关键词",
				},
				"page": map[string]interface{}{
					"type":        "number",
					"description": "页码",
					"default":     1,
				},
				"range": map[string]interface{}{
					"type":        "string",
					"description": "时间范围，可选值：1d(1天)、3d(3天)、7d(7天)、1m(1个月)、6m(6个月)、1y(1年)",
					"default":     "1m",
				},
			},
			"required": []string{"keyword"},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			Keyword string `json:"keyword"`
			Page    int    `json:"page"`
			Range   string `json:"range"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		if args.Keyword == "" {
			return nil, fmt.Errorf("keyword is required")
		}

		if args.Page == 0 {
			args.Page = 1
		}

		if args.Range == "" {
			args.Range = "1m"
		}

		client := getClient()
		posts, err := client.SearchHistory(args.Keyword, args.Page, args.Range)
		if err != nil {
			return nil, err
		}

		var result strings.Builder
		for _, post := range *posts {
			result.WriteString(post.FriendlyText())
			result.WriteString("\n\n")
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: result.String()},
			},
		}, nil
	})
}

// LikePostTool 点赞帖子工具
func LikePostTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "like_post",
		Description: "点赞指定的帖子",
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
		success, err := client.LikePost(args.ThreadID)
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

// UnlikePostTool 取消点赞帖子工具
func UnlikePostTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "unlike_post",
		Description: "取消点赞指定的帖子",
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
		success, err := client.UnLikePost(args.ThreadID)
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

// CreatePostTool 创建帖子工具
func CreatePostTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "create_post",
		Description: "在集市中创建新帖子",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"title": map[string]interface{}{
					"type":        "string",
					"description": "帖子标题",
				},
				"content": map[string]interface{}{
					"type":        "string",
					"description": "帖子内容",
				},
				"cate_id": map[string]interface{}{
					"type":        "string",
					"description": "分类ID",
				},
				"img_paths": map[string]interface{}{
					"type":        "string",
					"description": "图片路径（多个用逗号分隔）",
					"default":     "",
				},
				"contact_person": map[string]interface{}{
					"type":        "string",
					"description": "联系人",
					"default":     "",
				},
				"contact_phone": map[string]interface{}{
					"type":        "string",
					"description": "联系电话",
					"default":     "",
				},
				"contact_qq": map[string]interface{}{
					"type":        "string",
					"description": "联系QQ",
					"default":     "",
				},
				"contact_wx": map[string]interface{}{
					"type":        "string",
					"description": "联系微信",
					"default":     "",
				},
				"is_comment_close": map[string]interface{}{
					"type":        "number",
					"description": "是否关闭评论（0:否, 1:是）",
					"default":     0,
				},
			},
			"required": []string{"title", "content", "cate_id"},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			Title          string `json:"title"`
			Content        string `json:"content"`
			CateID         string `json:"cate_id"`
			ImgPaths       string `json:"img_paths"`
			ContactPerson  string `json:"contact_person"`
			ContactPhone   string `json:"contact_phone"`
			ContactQQ      string `json:"contact_qq"`
			ContactWX      string `json:"contact_wx"`
			IsCommentClose int    `json:"is_comment_close"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		if args.Title == "" {
			return nil, fmt.Errorf("title is required")
		}

		if args.Content == "" {
			return nil, fmt.Errorf("content is required")
		}

		if args.CateID == "" {
			return nil, fmt.Errorf("cate_id is required")
		}

		client := getClient()
		success, err := client.CreatePost(args.Title, args.Content, args.CateID, args.ImgPaths, args.ContactPerson, args.ContactPhone, args.ContactQQ, args.ContactWX, args.IsCommentClose)
		if err != nil {
			return nil, err
		}

		text := "帖子创建成功"
		if !success {
			text = "帖子创建失败"
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: text},
			},
		}, nil
	})
}

// ChangePostStatusTool 修改帖子状态工具
func ChangePostStatusTool(s *mcp.Server) {
	s.AddTool(&mcp.Tool{
		Name:        "change_post_status",
		Description: "修改帖子状态，如结束帖子需求并隐藏发帖人信息",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"thread_id": map[string]interface{}{
					"type":        "string",
					"description": "帖子ID",
				},
				"action": map[string]interface{}{
					"type":        "string",
					"description": "操作类型，如 finish(结束帖子需求并隐藏发帖人信息)",
					"default":     "finish",
				},
			},
			"required": []string{"thread_id"},
		},
	}, func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var args struct {
			ThreadID string `json:"thread_id"`
			Action   string `json:"action"`
		}
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, err
		}

		if args.ThreadID == "" {
			return nil, fmt.Errorf("thread_id is required")
		}

		if args.Action == "" {
			args.Action = "finish"
		}

		client := getClient()
		success, err := client.ChangePostStatus(args.ThreadID, args.Action)
		if err != nil {
			return nil, err
		}

		text := "帖子状态修改成功"
		if !success {
			text = "帖子状态修改失败"
		}

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: text},
			},
		}, nil
	})
}
