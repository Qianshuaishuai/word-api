// @APIVersion 1.0.0
// @Title 亲知科技Api 1.0.0版本
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"gitlab.dreamdev.cn/ebag/knowtech-api/controllers"
)

func init() {

	beego.Router("/", &controllers.HealthCheckController{}, "get:HealthCheck")

	//用户系统相关
	beego.Router("/v1/user/register", &controllers.UserController{}, "post:Register")
	beego.Router("/v1/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/v1/user/forget", &controllers.UserController{}, "post:Forget")
	beego.Router("/v1/user/get", &controllers.UserController{}, "get:Get")
	beego.Router("/v1/user/code", &controllers.UserController{}, "post:Code")
	beego.Router("/v1/user/logout", &controllers.UserController{}, "post:Logout")
	beego.Router("/v1/user/setting", &controllers.UserController{}, "post:Setting")

	//同步书本相关
	beego.Router("/v1/book/add", &controllers.BookController{}, "post:AddBook")
	beego.Router("/v1/book/edit", &controllers.BookController{}, "post:EditBook")
	beego.Router("/v1/book/delete", &controllers.BookController{}, "post:DeleteBook")
	beego.Router("/v1/book/list", &controllers.BookController{}, "get:GetBookList")
	beego.Router("/v1/book/detail", &controllers.BookController{}, "get:GetBookDetail")

	//单道试题相关
	beego.Router("/v1/question/add", &controllers.QuestionController{}, "post:AddQuestion")
	beego.Router("/v1/question/edit", &controllers.QuestionController{}, "post:EditQuestion")
	beego.Router("/v1/question/delete", &controllers.QuestionController{}, "post:DeleteQuestion")
	beego.Router("/v1/question/list", &controllers.QuestionController{}, "get:GetQuestionList")
	beego.Router("/v1/question/detail", &controllers.QuestionController{}, "get:GetQuestionDetail")
}
