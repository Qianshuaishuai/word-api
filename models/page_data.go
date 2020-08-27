package models

import "time"

type Page struct {
	ID         int       `gorm:"column:id" json:"id"`
	UserID     int       `gorm:"column:user_id" json:"userId"`
	BookID     int       `gorm:"column:book_id" json:"bookId"`
	Cover      string    `gorm:"column:cover" json:"cover"`
	CreateTime time.Time `gorm:"column:create_time" json:"-"`

	Frames []PageFrame `gorm:"-" json:"frames"`
}
