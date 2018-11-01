package models

import (
	db "projects/cantlost/dbs"
)

type Area struct {
	ID         int
	Name       string
	ParentName string
	ParentID   string
	ShortName  string
	Code       int
	Zipcode    int
	Pinyin     string
	Lng        string
	Lat        string
	Level      int
	Position   string
	Sort       int
}

func (m *Area) CreateTable() {
	db.PGMaster.CreateTable(m)
}

func (m *Area) AddArea() (id int64, err error) {
	db.PGMaster.Create(m)
	return
}

func ListArea(page, pageSize int, filters ...interface{}) (lists []Area, count int64, err error) {
	return
}

func OneArea(id int) (m Area, err error) {
	db.PGMaster.Where("id = ?", id).First(m)
	return
}

func (m *Area) UpdateArea(id int) (n int64, err error) {
	return
}

func DeleteArea(id int) (n int64, err error) {
	return
}
