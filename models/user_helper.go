package models

import (
	"fmt"
	"math/rand"
	"time"

	"gitlab.dreamdev.cn/ebag/word-api/helper"
)

//获取验证码（随机6位）
func GetRandomCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

//发送短信-腾讯云
func SendSms(phone string) (code string, err error) {
	phones := make([]string, 1)
	params := make([]string, 1)
	randomCode := GetRandomCode()
	phones[0] = "+86" + phone
	params[0] = randomCode

	TencentSms(phones, params)
	return randomCode, nil
}

//生成用户密码加密（用salt）
func TranslateUserPassword(oldPwd string) (newPwd, salt string) {
	salt = helper.GetRandomString(4) //生成一个4位随机字符串的盐
	newPwd = helper.Md5Crypt(oldPwd, salt)

	return newPwd, salt
}

//判断密码是否正确
func JdugeUserPasswordCorrect(md5Pwd, password, salt string) bool {
	correctPwd := helper.Md5Crypt(password, salt)

	if correctPwd == md5Pwd {
		return true
	}

	return false
}
