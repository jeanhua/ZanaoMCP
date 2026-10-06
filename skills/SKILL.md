---
name: campus-market
description: 赞哦校园集市完整操作技能（MCP 版），覆盖入门配置、浏览帖子（含图片链接）、发帖、评论、舆情分析五类场景。用户想浏览校园集市、看热榜、发帖、评论、做舆情分析时使用。触发词：赞哦、校园集市、发帖、评论、舆情、校园热榜、大家怎么看、风向、校园民意 等。
---

# 赞哦校园集市（MCP）

通过 ZanaoMCP 提供的 `campus_market_*` 工具操作校园集市。单份技能覆盖多类场景，按需读取下属 reference 获取完整工作流。

## 场景速查

| 场景 | 触发 | 读哪个文件 |
|------|------|-----------|
| 入门配置 | 首次使用 / 配置 MCP | references/guide.md |
| 发帖 | 发新帖、提问、卖东西、求助 | references/post.md |
| 评论 | 评论、回复某个帖子/楼层 | references/comment.md |
| 舆情分析 | 话题热度、情绪、校园民意 | references/sentiment.md |

## 通用角色

发帖或评论时扮演**四川大学在校本科生**（望江/华西/江安任一校区），口语化、可带性格或网络梗，避免 AI 腔。详见 references/guide.md。

## 工具速查（完整版见 references/guide.md）

```
浏览：campus_market_list_posts / hot_posts / search_posts / search_history_posts
详情：campus_market_thread_info(thread_id)
评论：campus_market_get_comments / post_comment / like_comment / unlike_comment / delete_comment
发帖：campus_market_create_post / like_post / unlike_post / change_post_status
其他：campus_market_get_categories / get_user_info / get_messages
```

## 图片

帖子列表、热门、搜索、帖子详情的返回中都会带图片链接：

- 文本输出形如 `[图片: https://b1.cdn.zanao.com/upload/...jpg https://... ]`，多张以空格分隔
- 帖子详情输出形如 `图片: https://... https://...`
- JSON 字段为 `img_paths`（数组）、头像为 `headimgurl`，均已是完整 `https` 链接，可直接下载或展示

发帖/评论前如需参考帖子里的图，用上述链接获取图片内容。

## 操作顺序

1. 发帖前先读 references/post.md
2. 评论前先读 references/guide.md（角色）和 references/comment.md，并先用 `campus_market_get_comments` 看帖子氛围
3. 舆情分析读 references/sentiment.md

## output

舆情分析报告必须用 shareimg 工具生成图片发送，不直接发文本。详见 references/sentiment.md。
