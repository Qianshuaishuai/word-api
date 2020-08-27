package controllers

import (
	"github.com/astaxie/beego"
	"gitlab.dreamdev.cn/ebag/knowtech-api/models"
)

type UserController struct {
	BaseController
}

type RegisterParam struct {
	Phone    string `form:"phone" valid:"Required"`
	Password string `form:"password" valid:"Required"`
	Code     string `form:"code" valid:"Required"`
}

type ForgetParam struct {
	Phone    string `form:"phone" valid:"Required"`
	Password string `form:"password" valid:"Required"`
	Code     string `form:"code" valid:"Required"`
}

type CodeParam struct {
	Phone  string `form:"phone" valid:"Required"`
	TypeID int    `form:"typeID"`
}

type GetParam struct {
	Phone string `form:"phone" valid:"Required"`
}

type VipParam struct {
	Phone string `form:"phone" valid:"Required"`
}

type ChangePasswordParam struct {
	Phone    string `form:"phone" valid:"Required"`
	Password string `form:"password" valid:"Required"`
	Code     string `form:"code" valid:"Required"`
}

type LoginParam struct {
	Phone    string `form:"phone" valid:"Required"`
	Password string `form:"password" valid:"Required"`
}

type LogoutParam struct {
	Phone string `form:"phone" valid:"Required"`
}

// @Title 用户注册
// @Description 用户注册
// @Param phone form int true 用户手机
// @Router /v1/user/register [post]
func (u *UserController) Register() {
	datas := u.GetResponseData()
	//get params
	params := &RegisterParam{}
	if u.CheckParams(datas, params) {
		err := models.Register(params.Phone, params.Password, params.Code)
		beego.Debug(err)
		u.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	u.jsonEcho(datas)
}

// @Title 用户登出
// @Description 用户登出
// @Param phone form int true 用户手机
// @Router /v1/user/logout [post]
func (u *UserController) Logout() {
	datas := u.GetResponseData()
	//get params
	params := &LogoutParam{}
	if u.CheckParams(datas, params) {
		err := models.LoginOut(params.Phone)

		u.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	u.jsonEcho(datas)
}

// @Title 获取用户资料
// @Description 获取用户资料
// @Router /v1/user/get [get]
func (u *UserController) Get() {
	datas := u.GetResponseData()
	params := &GetParam{}
	if u.CheckParams(datas, params) {
		data, err := models.GetUser(params.Phone)
		datas["F_data"] = data
		u.IfErr(datas, models.RESP_NO_RESULT, err)
	}
	u.jsonEcho(datas)
}

// @Title 获取验证玛
// @Description 获取验证玛
// @Param phone form int true 用户手机
// @Router /v1/user/code [post]
func (u *UserController) Code() {
	datas := u.GetResponseData()
	//get params
	params := &CodeParam{}
	if u.CheckParams(datas, params) {
		err := models.SendSmsCode(params.Phone)

		u.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	u.jsonEcho(datas)
}

// @Title 用户登陆
// @Description 用户登陆
// @Param phone query int true 用户手机
// @Param password query int true 用户密码
// @Router /v1/user/login [get]
func (u *UserController) Login() {
	datas := u.GetResponseData()
	params := &LoginParam{}
	if u.CheckParams(datas, params) {
		data, err := models.CommonLogin(params.Phone, params.Password)
		datas["F_data"] = data
		u.IfErr(datas, models.RESP_NO_RESULT, err)
	}
	u.jsonEcho(datas)
}

// @Title 重设密码
// @Description 重设密码
// @Param phone query string true 用户手机
// @Router /v1/user/forget [post]
func (u *UserController) Forget() {
	datas := u.GetResponseData()
	params := &ForgetParam{}
	if u.CheckParams(datas, params) {
		err := models.ChangePassword(params.Phone, params.Password, params.Code)
		u.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	u.jsonEcho(datas)
}
