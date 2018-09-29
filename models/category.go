package models

import (
	db "projects/cantlost/dbs"
	"time"
)

type Category struct {
	ID        int        `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Title     string     `json:"title"`
}

func (m *Category) CreateTable() {
	db.PGMaster.CreateTable(m)
}

func (m *Category) AddCategory() (id int64, err error) {
	db.PGMaster.Create(m)
	return
}

func ListCategory(page, pageSize int, filters ...interface{}) (lists []Category, count int64, err error) {
	return
}

func OneCategory(id int) (m Category, err error) {
	db.PGMaster.Where("id = ?", id).First(m)
	return
}

func (m *Category) UpdateCategory(id int) (n int64, err error) {
	return
}

func DeleteCategory(id int) (n int64, err error) {
	return
}
