package models

import "time"

//用户
type User struct {
	ID         int       `gorm:"column:id" json:"id"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	Enable     int       `gorm:"column:enable" json:"enable"`
	Password   string    `gorm:"column:password" json:"-"`
	Salt       string    `gorm:"column:salt" json:"-"`
	Username   string    `gorm:"column:username" json:"username"`
	Grade      int       `gorm:"column:grade" json:"grade"`
	Subjects   string    `gorm:"column:subjects" json:"subjects"`
	Organ      string    `gorm:"column:organ" json:"organ"`
	Address    string    `gorm:"column:address" json:"address"`
	Introduce  string    `gorm:"column:introduce" json:"introduce"`
	CreateTime time.Time `gorm:"column:create_time" json:"-"`
}

//验证码
type Code struct {
	Phone string    `gorm:"column:phone"`
	Code  string    `gorm:"column:code"`
	Time  time.Time `gorm:"column:time"`
}

//登录状态
type LoginStatus struct {
	Phone string    `gorm:"column:phone" json:"phone"`
	Time  time.Time `gorm:"column:time" json:"time"`
}
