# 赞哦入门配置（MCP）

## 配置 MCP Server

在 MCP 客户端（Claude Desktop / Cursor 等）配置文件中添加：

```json
{
  "mcpServers": {
    "zanao campus market": {
      "command": "/path/to/zanao-mcp",
      "env": {
        "ZANAO_TOKEN": "your_token",
        "ZANAO_SCHOOL_ALIAS": "scu"
      }
    }
  }
}
```

**如何获取 token**：登录赞哦电脑端微信小程序，用抓包工具（Fiddler、Charles、mitmproxy 均可）捕获任意请求，取请求头 `X-Sc-Od` 的值即为 token。学校别名如四川大学为 `scu`。

## 工具速查

```
# 浏览
campus_market_list_posts(from_time="0")        # 最新帖子，from_time 传上一页看到的 timestamp 可翻页
campus_market_hot_posts()                      # 热门帖子
campus_market_search_posts(keyword="二手书", page=1)          # 实时搜索
campus_market_search_history_posts(keyword="考研", page=1, range="1m")  # 历史搜索
campus_market_thread_info(thread_id="xxx")     # 帖子详情（含图片）

# 评论
campus_market_get_comments(thread_id="xxx")    # 查看评论

# 分类
campus_market_get_categories()                 # 查看所有分类及 ID（发帖时用）

# 用户
campus_market_get_user_info()                  # 当前用户信息
campus_market_get_messages()                   # 消息通知
```

## 图片字段

帖子类接口返回的图片信息：

- 文本输出：`[图片: https://... https://...]`，多张以空格分隔
- JSON 字段：`img_paths`（数组，完整 https 链接）、`headimgurl`（发帖人头像）
- 帖子详情文本中为单独一行 `图片: https://...`

**取值注意**：赞哦图床禁止访问原图，不带样式后缀会返回 403（AccessDenied）。`img_paths` 的链接会自动补上 `@!common`，头像自带 `@!sm_w100_h100`。这些后缀**不能去掉**，手动拼 URL 时也必须带上。`@!common` 是官方网页端使用的规格，保留原始宽高比、长边最高 1080、不裁剪。

## 角色要求

在代为发帖或评论时，扮演**四川大学在校本科生**：

- **身份**：川大在校生（望江/华西/江安任一校区均可）
- **语气**：口语化，可带点不耐烦或幽默，偶尔用网络梗（绝绝子、蚌埠住了、抽象）
- **禁忌**：不用"我来帮您"、"当然可以"、"以下是……"等 AI 腔，不用 Markdown 加粗或列表
- **地名梗**：望江、江安、磨子桥、望鹤、老图、华西坝 等地名用起来更自然
