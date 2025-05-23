package v1

import (
	. "github.com/alimy/mir/v5"

	"github.com/rocboss/paopao-ce/internal/model/web"
)

// Loose 宽松授权的服务
type Loose struct {
	Schema `mir:"v1,chain"`

	// Timeline 获取广场流
	Timeline func(Get, web.TimelineReq) web.TimelineResp `mir:"posts"`

	// GetUserTweets 获取用户动态列表
	GetUserTweets func(Get, web.GetUserTweetsReq) web.GetUserTweetsResp `mir:"user/posts"`

	// GetUserProfile 获取用户基本信息
	GetUserProfile func(Get, web.GetUserProfileReq) web.GetUserProfileResp `mir:"user/profile"`

	// TopicList 获取话题列表
	TopicList func(Get, web.TopicListReq) web.TopicListResp `mir:"tags"`

	// TweetComments 获取动态评论
	TweetComments func(Get, web.TweetCommentsReq) web.TweetCommentsResp `mir:"post/comments"`

	// TweetDetail 获取动态详情
	TweetDetail func(Get, web.TweetDetailReq) web.TweetDetailResp `mir:"post"`
}
