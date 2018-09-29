package models

import (
	db "projects/cantlost/dbs"
	"time"
)

type Member struct {
	ID          int        `json:"id"`
	RegistedAt  *time.Time `json:"registed_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	PicURL      string     `json:"pic_url"`
	NickName    string     `json:"nick_name"`
	RealName    string     `json:"real_name"`
	PhoneNumber string     `json:"phone_number"`
	Gender      string     `json:"gender"`
	BirthDay    *time.Time `json:"birthDay"`
	LoginName   string     `json:"login_name" form:"login_name"`
	Password    string     `json:"password" form:"password"`
}

func (m *Member) CreateTable() {
	db.PGMaster.CreateTable(m)
}

func (m *Member) AddMember() (id int64, err error) {
	db.PGMaster.Create(m)
	return
}

func ListMember(page, pageSize int, filters ...interface{}) (lists []Member, count int64, err error) {
	db.PGMaster.Limit(pageSize).Offset((page - 1) * pageSize).Scan(lists)
	return
}

func OneMember(id int) (m Member, err error) {
	//TODO 加入redis缓存
	db.PGMaster.Where("id = ?", id).First(&m)
	return
}

func (m *Member) UpdateMember(id int) (n int64, err error) {
	return
}

func DeleteMember(id int) (n int64, err error) {
	return
}
