package models

import (
	db "projects/cantlost/dbs"
	"time"
)

type LostInfoForm struct {
	MemberID     int64       `json:"user_id" form:"user_id"`         //用户id
	CategoryID   int64       `json:"category_id" form:"category_id"` //分类ID
	Title        string      `json:"title" form:"title"`             //标题
	CoverPics    string      `json:"cover_pics"`                     //封面图片（多张，分号分割））
	TagertInfo   *TargetForm `json:"target_info"`                    //丢失对象详情
	LostTimeStr  string      `json:"lost_time_str"`                  //丢失时间 2018-10-09 11:15:31 字符串格式
	LostPlaceStr string      `json:"lost_place_str"`                 //丢失地点描述
	LostPlaceLon string      `json:"lost_place_lon"`                 //丢失地点经度
	LostPlaceLat string      `json:"lost_place_lat"`                 //丢失地点纬度
	LostEvent    string      `json:"lost_event" form:"lost_event"`   //丢失事件经过,详细描述，前后发生了什么事情
	TargetStory  string      `json:"target_stroy"`                   //和Ta的回忆
	ThenEnd      string      `json:"the_end"`                        //故事结局
}
type LostInfoVo struct {
	ID           int64      `json:"id" form:"id"`
	PublishAt    *time.Time `json:"publish_at" form:"publish_at"`   //发布时间
	UpdatedAt    *time.Time `json:"updated_at" form:"updated_at"`   //更新时间
	MemberID     int64      `json:"member_id"`                      //用户id
	MemberInfo   *Member    `json:"member_info" form:"member_info"` //用户信息
	CategoryID   int64      `json:"category_id" form:"category_id"` //分类ID
	CategoryInfo *Category  `json:"category_info"`                  //分类信息
	Title        string     `json:"title" form:"title"`             //标题
	CoverPics    string     `json:"cover_pics"`                     //封面图片（多张，分号分割））
	LostTime     *time.Time `json:"lost_time"`                      //丢失时间
	LostPlaceStr string     `json:"lost_place_str"`                 //丢失地点描述
	LostPlaceLon string     `json:"lost_place_lon"`                 //丢失地点经度
	LostPlaceLat string     `json:"lost_place_lat"`                 //丢失地点纬度
	TargetID     int64      `json:"target_id" form:"target_id"`     //丢失对象信息id
	TagertInfo   *Target    `json:"target_info"`                    //丢失对象详情
	LostEvent    string     `json:"lost_event" form:"lost_event"`   //丢失事件经过,详细描述，前后发生了什么事情
	TargetStory  string     `json:"target_stroy"`                   //和Ta的回忆
	ThenEnd      string     `json:"the_end"`                        //故事结局
	LikeNum      int64      `json:"like_num" form:"like_num"`       //点赞数
	DislikeNum   int64      `json:"dislike_num" form:"disLike_num"` //不喜欢数
	CommentNum   int64      `json:"comment_num" form:"comment_num"` //评论数据
	ShareNum     int64      `json:"share_num" form:"share_num"`     //分享数
}

type LostInfo struct {
	ID           int64      `sql:"primerykey" json:"id" form:"id"`
	PublishAt    *time.Time `json:"publish_at" form:"publish_at"`   //发布时间
	UpdatedAt    *time.Time `json:"updated_at" form:"updated_at"`   //更新时间
	MemberID     int64      `json:"member_id"`                      //用户id
	CategoryID   int64      `json:"category_id" form:"category_id"` //分类ID
	Title        string     `json:"title" form:"title"`             //标题
	CoverPics    string     `json:"cover_pics"`                     //封面图片（多张，分号分割））
	LostTime     *time.Time `json:"lost_time"`                      //丢失时间
	LostPlaceStr string     `json:"lost_place_str"`                 //丢失地点描述
	LostPlaceLon string     `json:"lost_place_lon"`                 //丢失地点经度
	LostPlaceLat string     `json:"lost_place_lat"`                 //丢失地点纬度
	TargetID     int64      `json:"target_id" form:"target_id"`     //丢失对象信息id
	LostEvent    string     `json:"lost_event" form:"lost_event"`   //丢失事件经过,详细描述，前后发生了什么事情
	TargetStory  string     `json:"target_stroy"`                   //和Ta的回忆
	ThenEnd      string     `json:"the_end"`                        //故事结局
	LikeNum      int        `json:"like_num" form:"like_num"`       //点赞数
	DislikeNum   int64      `json:"dislike_num" form:"disLike_num"` //不喜欢数
	CommentNum   int64      `json:"comment_num" form:"comment_num"` //评论数据
	ShareNum     int64      `json:"share_num" form:"share_num"`     //分享数
}

func (m *LostInfo) CreateTable() {
	db.PGMaster.CreateTable(m)
}

func (m *LostInfo) AddLostInfo() (id int64, err error) {
	db.PGMaster.Create(m)
	return
}

func ListLostInfo(page, pageSize int, filters ...interface{}) (lists []LostInfoVo, count int64, err error) {
	db.PGMaster.Model(&LostInfo{}).Limit(pageSize).Offset((page - 1) * pageSize).Scan(&lists)
	count = (int64)(len(lists))
	return
}

func OneLostInfo(id int) (m LostInfoVo, err error) {
	db.PGMaster.Where("id = ?", id).First(m)
	return
}

func (m *LostInfo) UpdateLostInfo(id int) (n int64, err error) {
	return
}

func DeleteLostInfo(id int) (n int64, err error) {
	return
}
