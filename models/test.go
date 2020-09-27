package models

func Test() {
	GetDb().Table("t_users").Where("phone = ?", "15602335027").Update("username", "Dashuai")
}
