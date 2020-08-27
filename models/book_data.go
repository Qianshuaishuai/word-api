package models

import "time"

type Book struct {
	ID         int       `gorm:"column:id" json:"id"`
	UserID     int       `gorm:"column:user_id" json:"userId"`
	Name       string    `gorm:"column:name" json:"name"`
	Press      string    `gorm:"column:press" json:"press"`
	Grade      int       `gorm:"column:grade" json:"grade"`
	PubDate    string    `gorm:"column:pub_date" json:"pubDate"`
	Cover      string    `gorm:"column:cover" json:"cover"`
	ImageAi    int       `gorm:"column:image_ai" json:"imageAi"`
	CreateTime time.Time `gorm:"column:create_time" json:"-"`

	Pages []Page `gorm:"-" json:"pages"`
}
