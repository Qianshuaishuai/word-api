package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	loglib "github.com/HYY-yu/LogLib"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"gitlab.dreamdev.cn/ebag/knowtech-api/helper"
	"gitlab.dreamdev.cn/ebag/knowtech-api/models"
)

//公共controller
type BaseController struct {
	beego.Controller
	UniqueLogFlag string
	startTime     time.Time
	AppCtx        context.Context
}

type ResponseData map[string]interface{}

func (self *BaseController) Prepare() {
	//生成用户记录日志的唯一id
	self.UniqueLogFlag = helper.GetGuid()
	self.startTime = time.Now()
	//log请求
	self.logRequest()
}

func (self *BaseController) GetResponseData() ResponseData {
	return ResponseData{"F_responseNo": models.RESP_OK}
}

func (self *BaseController) CheckParams(datas ResponseData, params interface{}) bool {
	if err := self.ParseForm(params); err != nil {
		datas["F_responseNo"] = models.RESP_PARAM_ERR
		return false
	}

	//验证参数
	valid := validation.Validation{}
	if ok, _ := valid.Valid(params); ok {
		return true
	}

	datas["F_responseNo"] = models.RESP_PARAM_ERR
	datas["F_responseMsg"] = fmt.Sprint(valid.ErrorsMap)
	return false
}

func (self *BaseController) IfErr(datas ResponseData, no int, err error) bool {
	if err != nil {
		datas["F_responseNo"] = no
		datas["F_responseMsg"] = err.Error()
		return true
	}
	return false
}

//json echo
func (self *BaseController) jsonEcho(datas ResponseData) {
	var responseNoInt int
	var responseMsgStr string

	if responseNo, ok := datas["F_responseNo"]; ok {
		responseNoInt = responseNo.(int)
	} else {
		responseNoInt = models.RESP_ERR
	}

	if responseMsg, ok := datas["F_responseMsg"]; ok && len(responseMsg.(string)) > 0 {
		responseMsgStr = responseMsg.(string)
	} else {
		datas["F_responseMsg"] = ""
		if msg, ok := models.MyConfig.ConfigMyResponse[datas["F_responseNo"].(int)]; ok {
			datas["F_responseMsg"] = msg
			responseMsgStr = msg
		}
	}

	self.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	//跨域支持
	self.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")

	self.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	switch responseNoInt {
	case models.RESP_ERR:
		self.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	case models.RESP_NO_ACCESS:
		self.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	case models.RESP_TOKEN_ERR:
		self.Ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
	}

	self.Data["json"] = datas

	//Log
	self.logEcho(self.Ctx.ResponseWriter.Status,
		responseNoInt, responseMsgStr,
		time.Since(self.startTime).String())

	//输出
	self.ServeJSON()
}

//记录请求
func (self *BaseController) logRequest() {
	loglib.GetLogger().LogRequest(self.Ctx, self.UniqueLogFlag)
}

//记录输出
func (self *BaseController) logEcho(statusCode int, responseNo int, responseMsg string, apiTime string) {
	loglib.GetLogger().LogResponse(statusCode, responseNo, responseMsg, apiTime, self.UniqueLogFlag)
}

func checkBaseParams(datas ResponseData, F_teacher_id string, F_accesstoken string) {
	if len(strings.TrimSpace(F_teacher_id)) == 0 {
		datas["F_responseNo"] = models.RESP_PARAM_ERR
		datas["F_responseMsg"] = "F_teacher_id不能为空"
		return
	}
	if len(strings.TrimSpace(F_accesstoken)) == 0 {
		datas["F_responseNo"] = models.RESP_PARAM_ERR
		datas["F_responseMsg"] = "F_accessToken不能为空"
		return
	}
}
