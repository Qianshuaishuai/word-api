package models

import "time"

type Question struct {
	ID            int             `gorm:"column:id" json:"id"`
	UserID        int             `gorm:"column:user_id" json:"userId"`
	Subject       int             `gorm:"column:subject" json:"subject"`
	Grade         int             `gorm:"column:grade" json:"grade"`
	Name          string          `gorm:"column:name" json:"name"`
	ImageAi       int             `gorm:"column:image_ai" json:"imageAi"`
	Image         string          `gorm:"column:image" json:"image"`
	CreateTime    time.Time       `gorm:"column:create_time" json:"-"`
	Frames        []QuestionFrame `gorm:"-" json:"frames"`
	CreateTimeStr string          `gorm:"-" json:"createTimeStr"`
}
