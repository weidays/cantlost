package vo

import (
	"projects/cantlost/models"
	"time"
)

type LostInfoVo struct {
	ID           int              `json:"id" form:"id"`
	PublishAt    *time.Time       `json:"publish_at" form:"publish_at"`   //发布时间
	UpdatedAt    *time.Time       `json:"updated_at" form:"updated_at"`   //更新时间
	UserID       int              `json:"user_id" form:"user_id"`         //用户id
	MemberInfo   *models.Member   `json:"member_info" form:"member_info"` //用户信息
	CategoryInfo *models.Category `json:"category_info"`                  //分类信息
	CategoryID   int              `json:"category_id" form:"category_id"` //分类ID
	Title        string           `json:"title" form:"title"`             //标题
	Content      string           `json:"content" form:"content"`         //内容
	LikeNum      int              `json:"like_num" form:"like_num"`       //点赞数
	DislikeNum   int              `json:"dislike_num" form:"disLike_num"` //不喜欢数
	CommentNum   int              `json:"comment_num" form:"comment_num"` //评论数据
	ShareNum     int              `json:"share_num" form:"share_num"`     //分享数
}
