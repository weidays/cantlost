package models

import (
	db "projects/cantlost/dbs"
	"time"
)

type TargetOtherInfoForm struct {
	Option string `json:"option"` //对象名称
	Value  string `json:"value"`
}

//TargetOtherInfo 丢失对象信息
type TargetOtherInfo struct {
	ID                int        `json:"id" form:"id"`
	CreateAt          *time.Time `json:"publish_at" form:"publish_at"` //发布时间
	UpdatedAt         *time.Time `json:"updated_at" form:"updated_at"` //更新时间
	TargetOtherInfoID string     `json:"TargetOtherInfo_id"`           //分类（手机？电子产品？人？动物？）
	Option            string     `json:"option"`                       //对象名称
	Value             string     `json:"value"`
}

func ListTargetOtherInfo(page, pageSize int, filters ...interface{}) (lists []TargetOtherInfo, count int64, err error) {
	db.PGMaster.Model(&TargetOtherInfo{}).Limit(pageSize).Offset((page - 1) * pageSize).Scan(&lists)
	count = (int64)(len(lists))
	return
}

func OneTargetOtherInfo(id int) (m TargetOtherInfo, err error) {
	db.PGMaster.Where("id = ?", id).First(m)
	return
}

func (m *TargetOtherInfo) UpdateTargetOtherInfo(id int) (n int64, err error) {
	return
}

func DeleteTargetOtherInfo(id int) (n int64, err error) {
	return
}
