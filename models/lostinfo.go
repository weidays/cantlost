package models

import (
	db "projects/cantlost/dbs"
	"time"
)

type LostInfo struct {
	ID         int        `json:"id" form:"id"`
	PublishAt  *time.Time `json:"publish_at" form:"publish_at"`   //发布时间
	UpdatedAt  *time.Time `json:"updated_at" form:"updated_at"`   //更新时间
	UserID     int        `json:"user_id" form:"user_id"`         //用户id
	CategoryID int        `json:"category_id" form:"category_id"` //分类ID
	Title      string     `json:"title" form:"title"`             //标题
	Content    string     `json:"content" form:"content"`         //内容
	LikeNum    int        `json:"like_num" form:"like_num"`       //点赞数
	DislikeNum int        `json:"dislike_num" form:"disLike_num"` //不喜欢数
	CommentNum int        `json:"comment_num" form:"comment_num"` //评论数据
	ShareNum   int        `json:"share_num" form:"share_num"`     //分享数
}

func (m *LostInfo) CreateTable() {
	db.PGMaster.CreateTable(m)
}

func (m *LostInfo) AddLostInfo() (id int64, err error) {
	db.PGMaster.Create(m)
	return
}

func ListLostInfo(page, pageSize int, filters ...interface{}) (lists []LostInfo, count int64, err error) {
	db.PGMaster.Model(&LostInfo{}).Limit(pageSize).Offset((page - 1) * pageSize).Scan(&lists)
	count = (int64)(len(lists))
	return
}

func OneLostInfo(id int) (m LostInfo, err error) {
	db.PGMaster.Where("id = ?", id).First(m)
	return
}

func (m *LostInfo) UpdateLostInfo(id int) (n int64, err error) {
	return
}

func DeleteLostInfo(id int) (n int64, err error) {
	return
}
