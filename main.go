package main

import (
	loglib "github.com/HYY-yu/LogLib"
	"github.com/astaxie/beego"
	"gitlab.dreamdev.cn/ebag/knowtech-api/controllers"
	"gitlab.dreamdev.cn/ebag/knowtech-api/models"
	_ "gitlab.dreamdev.cn/ebag/knowtech-api/routers"
)

func main() {
	//连接数据库
	models.InitGorm()
	db := models.GetDb()
	defer db.Close()

	//就问能不能ping通
	errPing := db.DB().Ping()
	if errPing != nil {
		loglib.GetLogger().LogErr(errPing, "can't connect db")
		return
	}
	models.Test()
	//如果服务器Panic ，返回500错，而不是错误信息。并且记录
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
