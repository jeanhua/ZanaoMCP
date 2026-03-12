package zanao

import (
	"encoding/json"
	"fmt"
)

type basicData[T any] struct {
	ErrorNo  int    `json:"errno"`
	ErrorMsg string `json:"errmsg"`
	Data     T      `json:"data"`
}

type dataList[T any] basicData[struct {
	List []T `json:"list"`
}]

type Category struct {
	CateID  string `json:"cate_id"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

func (c Category) FriendlyText() string {
	return "[ID:" + c.CateID + "] " + c.Name + ": " + c.Summary
}

type cateList basicData[struct {
	CateList []Category `json:"cate_list"`
}]

// 帖子列表
// POST https://api.x.zanao.com/thread/v2/list?from_time=0&with_comment=false&with_reply=false
// 热门帖子
// POST https://api.x.zanao.com/thread/hot?count=10&type=3

// 实时搜索
// https://api.x.zanao.com/thread/v2/search?wd=xxx&cur_page=1&cate_id=10
// 历史搜索，range=1m表示搜索最近1个月，1d表示搜索最近1天，3d表示搜索最近3天，7d表示搜索最近7天，6m表示搜索最近6个月,1y表示搜索最近1年
// https://api.x.zanao.com/thread/v2/search?wd=xxx&cur_page=1&cate_id=20&range=1m

type Post struct {
	NickName     string      `json:"nickname"`
	Title        string      `json:"title"`
	Content      string      `json:"content"`
	CateName     string      `json:"cate_name"`
	ViewCount    json.Number `json:"view_count"`
	CommentCount json.Number `json:"c_count"`
	LikeCount    json.Number `json:"l_count"`
	PostTime     string      `json:"post_time"`
	PTime        string      `json:"p_time"`
	ThreadID     string      `json:"thread_id"`
}

func (p Post) FriendlyText() string {
	return "[ThreadID:" + p.ThreadID + "] [nickname: " + p.NickName + "] [" + p.CateName + "] [title: " + p.Title + "] " + p.Content +
		" (浏览" + fmt.Sprintf("%s", p.ViewCount.String()) + ", 评论" + fmt.Sprintf("%s", p.CommentCount.String()) +
		", 点赞" + fmt.Sprintf("%s", p.LikeCount.String()) + ", " + p.PostTime + "[timestamp: " + p.PTime + "])"
}

// 评论列表
// POST https://api.x.zanao.com/comment/list?id=xxx&sign=xxx
type Comment struct {
	NickName  string         `json:"nickname"`
	Content   string         `json:"content"`
	PostTime  string         `json:"post_time_text"`
	LikeNum   json.Number    `json:"like_num"`
	CommentID string         `json:"comment_id"`
	ReplyList []CommentReply `json:"reply_list"`
}

func (c Comment) FriendlyText() string {
	result := "[ID:" + c.CommentID + "] " + c.NickName + ": " + c.Content + " (点赞" + fmt.Sprintf("%s", c.LikeNum.String()) + ", " + c.PostTime + ")"
	if len(c.ReplyList) > 0 {
		result += "\n  回复:"
		for _, reply := range c.ReplyList {
			result += "\n    " + reply.FriendlyText()
		}
	}
	return result
}

type CommentReply struct {
	NickName       string      `json:"nickname"`
	Content        string      `json:"content"`
	PostTime       string      `json:"post_time_text"`
	LikeNum        json.Number `json:"like_num"`
	CommentID      string      `json:"comment_id"`
	ReplyCommentID string      `json:"reply_comment_id"`
}

func (cr CommentReply) FriendlyText() string {
	return "[ID:" + cr.CommentID + "] " + cr.NickName + ": " + cr.Content + " (点赞" + fmt.Sprintf("%s", cr.LikeNum.String()) + ", " + cr.PostTime + ")"
}

// 消息列表
// POST https://api.x.zanao.com/msg/list?from_time=0
type Message struct {
	MsgID      string `json:"msg_id"`
	MsgType    string `json:"msg_type"`
	MsgTitle   string `json:"msg_title"`
	CreateTime string `json:"create_time"`
	FromUser   struct {
		Nickname string `json:"nickname"`
	} `json:"from_user_info"`
	Thread      *ThreadInfo  `json:"thread_info,omitempty"`
	Comment     *CommentInfo `json:"comment_info,omitempty"`
	FromComment *CommentInfo `json:"from_comment_info,omitempty"`
}

type ThreadInfo struct {
	ThreadID string `json:"thread_id"`
	Title    string `json:"title"`
}

func (t *ThreadInfo) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '[' {
		return nil
	}
	type Alias ThreadInfo
	return json.Unmarshal(data, (*Alias)(t))
}

type CommentInfo struct {
	CommentID string      `json:"comment_id"`
	Content   string      `json:"content"`
	LikeNum   json.Number `json:"like_num"`
}

func (c *CommentInfo) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '[' {
		return nil
	}
	type Alias CommentInfo
	return json.Unmarshal(data, (*Alias)(c))
}

func (m Message) FriendlyText() string {
	result := "[ID:" + m.MsgID + "] [" + m.MsgType + "] " + m.MsgTitle + " (" + m.CreateTime + ")"
	result += "\n  来自: " + m.FromUser.Nickname
	if m.Thread != nil {
		result += "\n  帖子ID:" + m.Thread.ThreadID + " 标题: " + m.Thread.Title
	}
	if m.Comment != nil && m.Comment.CommentID != "" {
		result += "\n  评论ID:" + m.Comment.CommentID + " 内容: " + m.Comment.Content
	}
	if m.FromComment != nil && m.FromComment.CommentID != "" {
		result += "\n  回复给ID:" + m.FromComment.CommentID + " 内容: " + m.FromComment.Content
	}
	return result
}

// 用户信息
// POST https://api.x.zanao.com/user/info?from=mine
type UserInfo struct {
	Data struct {
		SchoolName string `json:"school_name"`
		Info       struct {
			NickName       string `json:"nickname"`
			UserLevelTitle string `json:"user_level_title"`
		} `json:"user_info"`
	} `json:"data"`
}

func (u UserInfo) FriendlyText() string {
	return u.Data.Info.NickName + " (等级头衔:" + u.Data.Info.UserLevelTitle + ") @ " + u.Data.SchoolName
}
