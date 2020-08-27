package models

import "time"

type Question struct {
	ID         int             `gorm:"column:id" json:"id"`
	UserID     int             `gorm:"column:user_id" json:"userId"`
	ImageAi    int             `gorm:"column:image_ai" json:"imageAi"`
	CreateTime time.Time       `gorm:"column:create_time" json:"-"`
	Frames     []QuestionFrame `json:"frames"`
}
