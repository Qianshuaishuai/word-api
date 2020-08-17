// @APIVersion 1.0.0
// @Title 初中古诗词阅读Api
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
}
