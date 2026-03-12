package zanao

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type ZanaoClient struct {
	client      *resty.Client
	schoolalias string
	token       string
}

func NewZanaoClient(token string, schoolAlias string) *ZanaoClient {
	return &ZanaoClient{
		client:      resty.New(),
		token:       token,
		schoolalias: schoolAlias,
	}
}

// ------------------------------------------------------------------------------- //
// 										读请求										//
// ------------------------------------------------------------------------------- //

func (c *ZanaoClient) fetchPost(url string) (*[]Post, error) {
	headers := getHeaders(c.token, c.schoolalias)
	var resp dataList[Post]
	_, err := c.client.R().
		SetHeaders(headers).
		SetResult(&resp).
		Post(url)

	if err != nil {
		return nil, err
	}
	return &resp.Data.List, nil
}

// GetPost 获取帖子列表
// fromTime: 起始时间戳，用于分页获取，设置为0表示从最新开始获取
func (c *ZanaoClient) GetPost(fromTime string) (*[]Post, error) {
	url := fmt.Sprintf("https://api.x.zanao.com/thread/v2/list?from_time=%s&with_comment=false&with_reply=false", fromTime)
	return c.fetchPost(url)
}

// GetHot 获取热门帖子列表
func (c *ZanaoClient) GetHot() (*[]Post, error) {
	return c.fetchPost("https://api.x.zanao.com/thread/hot?count=10&type=3")
}

// Search 搜索帖子（当前分类）
// keyword: 搜索关键词
// page: 页码
func (c *ZanaoClient) Search(keyword string, page int) (*[]Post, error) {
	keyword = url.QueryEscape(keyword)
	return c.fetchPost(fmt.Sprintf("https://api.x.zanao.com/thread/v2/search?wd=%s&cur_page=%d&cate_id=10", keyword, page))
}

// SearchHistory 搜索历史帖子
// keyword: 搜索关键词
// page: 页码
// rangeTime: 时间范围，可选值为1d、3d、7d、1m、6m、1y
func (c *ZanaoClient) SearchHistory(keyword string, page int, rangeTime string) (*[]Post, error) {
	keyword = url.QueryEscape(keyword)
	return c.fetchPost(fmt.Sprintf("https://api.x.zanao.com/thread/v2/search?wd=%s&cur_page=%d&cate_id=20&range=%s", keyword, page, rangeTime))
}

// GetComment 获取帖子的评论列表
// threadID: 帖子ID
// sign: 签名参数
func (c *ZanaoClient) GetComment(threadID string) (*[]Comment, error) {
	url := fmt.Sprintf("https://api.x.zanao.com/comment/list?id=%s", threadID)
	var resp dataList[Comment]
	_, err := c.client.R().
		SetHeaders(getHeaders(c.token, c.schoolalias)).
		SetResult(&resp).
		Post(url)
	if err != nil {
		return nil, err
	}
	return &resp.Data.List, nil
}

// GetMessage 获取用户消息列表
func (c *ZanaoClient) GetMessage() (*[]Message, error) {
	url := "https://api.x.zanao.com/msg/list?from_time=0"
	var resp dataList[Message]
	_, err := c.client.R().
		SetHeaders(getHeaders(c.token, c.schoolalias)).
		SetResult(&resp).
		Post(url)
	if err != nil {
		return nil, err
	}
	return &resp.Data.List, nil
}

// GetUserInfo 获取当前用户信息
func (c *ZanaoClient) GetUserInfo() (*UserInfo, error) {
	url := "https://api.x.zanao.com/user/info?from=mine"
	var resp UserInfo
	_, err := c.client.R().
		SetHeaders(getHeaders(c.token, c.schoolalias)).
		SetResult(&resp).
		Post(url)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCategory 获取帖子分类列表
func (c *ZanaoClient) GetCategory() (*[]Category, error) {
	url := "https://api.x.zanao.com/catelist?from=post&is_cross=0&cross_all=1"
	var resp cateList
	_, err := c.client.R().
		SetHeaders(getHeaders(c.token, c.schoolalias)).
		SetResult(&resp).
		Post(url)
	if err != nil {
		return nil, err
	}
	return &resp.Data.CateList, nil
}

// ------------------------------------------------------------------------------- //
//
//	写请求										//
//
// ------------------------------------------------------------------------------- //

func postData[T any](c *ZanaoClient, url string, data any) (*T, error) {
	var resp basicData[T]
	if data != nil {
		_, err := c.client.R().
			SetHeaders(getHeaders(c.token, c.schoolalias)).
			SetBody(data).
			SetResult(&resp).
			Post(url)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := c.client.R().
			SetHeaders(getHeaders(c.token, c.schoolalias)).
			SetResult(&resp).
			Post(url)
		if err != nil {
			return nil, err
		}
	}
	return &resp.Data, nil
}

// LikePost 点赞帖子
// threadID: 帖子ID
func (c *ZanaoClient) LikePost(threadID string) (bool, error) {
	url := fmt.Sprintf("https://api.x.zanao.com/thread/like?id=%s&comment_id=0&action=1", threadID)
	if _, err := postData[any](c, url, nil); err != nil {
		return false, err
	}
	return true, nil
}

// UnLikePost 取消点赞帖子
// threadID: 帖子ID
func (c *ZanaoClient) UnLikePost(threadID string) (bool, error) {
	url := fmt.Sprintf("https://api.x.zanao.com/thread/like?id=%s&comment_id=0&action=0", threadID)
	if _, err := postData[any](c, url, nil); err != nil {
		return false, err
	}
	return true, nil
}

// LikeComment 点赞评论
// threadID: 帖子ID
// commentID: 评论ID
func (c *ZanaoClient) LikeComment(threadID string, commentID string) (bool, error) {
	url := fmt.Sprintf("https://api.x.zanao.com/comment/like?id=%s&comment_id=%s&action=1", threadID, commentID)
	if _, err := postData[any](c, url, nil); err != nil {
		return false, err
	}
	return true, nil
}

// UnLikeComment 取消点赞评论
// threadID: 帖子ID
// commentID: 评论ID
func (c *ZanaoClient) UnLikeComment(threadID string, commentID string) (bool, error) {
	url := fmt.Sprintf("https://api.x.zanao.com/comment/like?id=%s&comment_id=%s&action=0", threadID, commentID)
	if _, err := postData[any](c, url, nil); err != nil {
		return false, err
	}
	return true, nil
}

// PostComment 发表评论
// threadID: 帖子ID
// content: 评论内容
// replyCommentID: 回复的评论ID（可为空，设置为0）
// rootCommentID: 根评论ID（用于楼中楼回复，可空，设置为0）
// useAnon: 是否匿名（0:否, 1:是）
func (c *ZanaoClient) PostComment(threadID string, content string, replyCommentID string, rootCommentID string, useAnon int) (bool, error) {
	if replyCommentID == "" {
		replyCommentID = "0"
	}
	if rootCommentID == "" {
		rootCommentID = "0"
	}
	url := fmt.Sprintf("https://api.x.zanao.com/comment/post?id=%s&content=%s&reply_comment_id=%s&root_comment_id=%s&use_anon=%d&from=detail", threadID, url.QueryEscape(content), replyCommentID, rootCommentID, useAnon)
	if _, err := postData[any](c, url, nil); err != nil {
		return false, err
	}
	return true, nil
}

// DeleteComment 删除评论
// threadID: 帖子ID
// commentID: 评论ID
func (c *ZanaoClient) DeleteComment(threadID string, commentID string) (bool, error) {
	url := fmt.Sprintf("https://api.x.zanao.com/comment/delete?comment_id=%s&id=%s", commentID, threadID)
	if _, err := postData[any](c, url, nil); err != nil {
		return false, err
	}
	return true, nil
}

// CreatePost 创建新帖子
// title: 帖子标题
// content: 帖子内容
// cateID: 分类ID
// imgPaths: 图片路径（多个用逗号分隔）
// contactPerson: 联系人
// contactPhone: 联系电话
// contactQQ: 联系QQ
// contactWX: 联系微信
// isCommentClose: 是否关闭评论（0:否, 1:是）
func (c *ZanaoClient) CreatePost(title string, content string, cateID string, imgPaths string, contactPerson string, contactPhone string, contactQQ string, contactWX string, isCommentClose int) (bool, error) {
	url := fmt.Sprintf("https://api.x.zanao.com/thread/post?contact_person=%s&contact_phone=%s&contact_qq=%s&contact_wx=%s&title=%s&content=%s&img_paths=%s&cate_id=%s&certShow=10&is_comment_close=%d",
		url.QueryEscape(contactPerson),
		url.QueryEscape(contactPhone),
		url.QueryEscape(contactQQ),
		url.QueryEscape(contactWX),
		url.QueryEscape(title),
		url.QueryEscape(content),
		url.QueryEscape(imgPaths),
		cateID,
		isCommentClose,
	)
	if _, err := postData[any](c, url, nil); err != nil {
		return false, err
	}
	return true, nil
}

// ChangePostStatus 修改帖子状态
// threadID: 帖子ID
// action: 操作类型 finish(结束帖子需求并隐藏发帖人信息)
func (c *ZanaoClient) ChangePostStatus(threadID string, action string) (bool, error) {
	url := fmt.Sprintf("https://api.x.zanao.com/thread/change?id=%s&act=%s", threadID, url.QueryEscape(action))
	if _, err := postData[any](c, url, nil); err != nil {
		return false, err
	}
	return true, nil
}
