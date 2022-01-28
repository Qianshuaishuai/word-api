package models

import "time"

type PageFrame struct {
	ID          int       `gorm:"column:id" json:"id"`
	UserID      int       `gorm:"column:user_id" json:"userId"`
	BookID      int       `gorm:"column:book_id" json:"bookId"`
	PageID      int       `gorm:"column:page_id" json:"pageId"`
	ResourceURL string    `gorm:"column:resource_url" json:"resourceUrl"`
	Position    string    `gorm:"column:position" json:"position"`
	CreateTime  time.Time `gorm:"column:create_time" json:"-"`
}

type QuestionFrame struct {
	ID          int       `gorm:"column:id" json:"id"`
	UserID      int       `gorm:"column:user_id" json:"userId"`
	QuestionID  int       `gorm:"column:question_id" json:"questionId"`
	ResourceURL string    `gorm:"column:resource_url" json:"resourceUrl"`
	Position    string    `gorm:"column:position" json:"position"`
	CreateTime  time.Time `gorm:"column:create_time" json:"-"`
}

type UploadFrame struct {
	Position string `json:"position"`
	URL      string `json:"url"`
}
