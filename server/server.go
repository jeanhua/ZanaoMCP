package server

import (
	"context"
	"log"

	"github.com/jeanhua/ZanaoMCP/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func StartServer() {
	s := mcp.NewServer(&mcp.Implementation{
		Name:    "zanao campus market",
		Title:   "赞哦校园集市 MCP Server",
		Version: "v1.0.0"},
		nil)

	// 注册帖子相关工具
	tools.ListPostsTool(s)
	tools.HotPostsTool(s)
	tools.SearchPostsTool(s)
	tools.SearchHistoryPostsTool(s)

	// 注册评论相关工具
	tools.GetCommentsTool(s)
	tools.PostCommentTool(s)
	tools.DeleteCommentTool(s)
	tools.LikeCommentTool(s)
	tools.UnlikeCommentTool(s)

	// 注册帖子操作工具
	tools.LikePostTool(s)
	tools.UnlikePostTool(s)
	tools.CreatePostTool(s)
	tools.ChangePostStatusTool(s)

	// 注册用户相关工具
	tools.GetMessagesTool(s)
	tools.GetUserInfoTool(s)
	tools.GetCategoriesTool(s)

	if err := s.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
