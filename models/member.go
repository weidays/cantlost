package models

import (
	db "projects/cantlost/dbs"
	"time"
)

type Member struct {
	ID          int `json:"id" form:"id"`
	RegistedAt  *time.Time
	UpdatedAt   *time.Time
	PicURL      string
	NickName    string
	RealName    string
	PhoneNumber string
	Gender      string
	BirthDay    *time.Time
	LoginName   string `json:"login_name" form:"login_name"`
	Password    string `json:"password" form:"password"`
}

func (m *Member) CreateTable() {
	db.PGMaster.CreateTable(m)
}

func (m *Member) AddMember() (id int64, err error) {
	db.PGMaster.Create(m)
	return
}

func ListMember(page, pageSize int, filters ...interface{}) (lists []Member, count int64, err error) {

	return
}

func OneMember(id int) (m Member, err error) {
	return
}

func (m *Member) UpdateMember(id int) (n int64, err error) {
	return
}

func DeleteMember(id int) (n int64, err error) {
	return
}
