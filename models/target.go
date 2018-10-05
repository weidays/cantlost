package models

import (
	db "projects/cantlost/dbs"
	"time"
)

//TargetForm
type TargetForm struct {
	Name       string                `json:"name"`        //对象名称
	Brand      string                `json:"brand"`       //对象品牌
	Value      string                `json:"value"`       //当前价值
	GetAt      *time.Time            `json:"get_at"`      //取得时间
	MainColor  string                `json:"main_color"`  //主颜色
	OtherColor string                `json:"other_color"` //次颜色
	OtherInfo  []TargetOtherInfoForm `json:"other_info"`  //其他自定义特征信息
}
type TargetVo struct {
	ID        int        `json:"id" form:"id"`
	CreateAt  *time.Time `json:"publish_at" form:"publish_at"` //发布时间
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at"` //更新时间
	// CategoryID string     `json:"category_id"`                  //分类（手机？电子产品？人？动物？）
	Name       string                `json:"name"`        //对象名称
	Brand      string                `json:"brand"`       //对象品牌
	Value      string                `json:"value"`       //当前价值
	GetAt      *time.Time            `json:"get_at"`      //取得时间
	MainColor  string                `json:"main_color"`  //主颜色
	OtherColor string                `json:"other_color"` //次颜色
	OtherInfo  []TargetOtherInfoForm `json:"other_info"`  //其他自定义特征信息
}

//Target 丢失对象信息
type Target struct {
	ID         int        `json:"id" form:"id"`
	CreateAt   *time.Time `json:"publish_at" form:"publish_at"` //发布时间
	UpdatedAt  *time.Time `json:"updated_at" form:"updated_at"` //更新时间
	CategoryID string     `json:"category_id"`                  //分类（手机？电子产品？人？动物？）
	Name       string     `json:"name"`                         //对象名称
	Brand      string     `json:"brand"`                        //对象品牌
	Value      string     `json:"value"`                        //当前价值
	GetAt      *time.Time `json:"get_at"`                       //取得时间
	MainColor  string     `json:"main_color"`                   //主颜色
	OtherColor string     `json:"other_color"`                  //次颜色
}

func (m *Target) AddTarget() (id int64, err error) {
	db.PGMaster.Create(m)
	return
}

func ListTarget(page, pageSize int, filters ...interface{}) (lists []Target, count int64, err error) {
	db.PGMaster.Model(&Target{}).Limit(pageSize).Offset((page - 1) * pageSize).Scan(&lists)
	count = (int64)(len(lists))
	return
}

func OneTarget(id int) (m Target, err error) {
	db.PGMaster.Where("id = ?", id).First(m)
	return
}

func (m *Target) UpdateTarget(id int) (n int64, err error) {
	return
}

func DeleteTarget(id int) (n int64, err error) {
	return
}
