package models

import (
	"errors"
	"time"
)

const (
	DEFAULT_OVERDUE_CODE_TIME = 100.0 //验证码失效
	DEFAULT_LOGIN_TIME        = 120
)

//修改
func Setting(phone string, grade int, subjects, organ, address, introduce string) (user User, err error) {
	GetDb().Table("t_users").Where("phone = ?", phone).Find(&user)

	if user.Phone == "" {
		return user, errors.New("未找到帐号")
	}

	user.Phone = phone
	user.Grade = grade
	user.Subjects = subjects
	user.Organ = organ
	user.Address = address
	user.Introduce = introduce
	user.CreateTime = time.Now()

	err = GetDb().Table("t_users").Where("phone = ?", phone).Update(&user).Error
	if err != nil {
		return user, errors.New("更新失败")
	}

	return user, nil
}

//注册接口
func Register(username, phone, password, codeStr string, grade int, subjects, organ, address, introduce string) (newUser User, err error) {
	var code Code
	GetDb().Table("t_codes").Where("phone = ? AND code = ?", phone, codeStr).Find(&code)

	if code.Phone == "" {
		return newUser, errors.New("验证玛错误")
	}

	nowTime := time.Now()
	offsetTime := nowTime.Sub(code.Time).Minutes()

	if offsetTime > DEFAULT_OVERDUE_CODE_TIME {
		return newUser, errors.New("验证玛过期")
	}

	var user User
	user.Phone = phone
	user.Password, user.Salt = TranslateUserPassword(password)
	user.Username = username
	user.Grade = grade
	user.Subjects = subjects
	user.Organ = organ
	user.Address = address
	user.Introduce = introduce
	user.CreateTime = time.Now()

	var count int
	GetDb().Table("t_users").Where("phone = ?", phone).Count(&count)

	if count > 0 {
		return newUser, errors.New("账号已注册")
	}

	err = GetDb().Table("t_users").Create(&user).Error

	if err != nil {
		return newUser, errors.New("创建失败")
	}

	var status LoginStatus
	status.Phone = phone
	status.Time = time.Now()

	var statusCount int
	GetDb().Table("t_login_status").Where("phone = ?", phone).Count(&statusCount)

	if statusCount <= 0 {
		GetDb().Table("t_login_status").Create(&status)
	} else {
		GetDb().Table("t_login_status").Where("phone = ?", phone).Update(&status)
	}

	GetDb().Table("t_users").Where("phone = ?", phone).Find(&newUser)

	if user.Phone == "" {
		return newUser, errors.New("未找到该用户信息")
	}

	return newUser, nil
}

//登录接口(普通模式)
func CommonLogin(phone, password string) (user User, err error) {
	var checkUser User

	phoneLength := len([]rune(phone))
	if phoneLength == 11 {
		GetDb().Table("t_users").Where("phone = ?", phone).Find(&checkUser)
	} else {
		GetDb().Table("t_users").Where("username = ?", phone).Find(&checkUser)
	}

	if checkUser.Phone == "" {
		return user, errors.New("该账号未注册")
	}

	isPwdCorrect := JdugeUserPasswordCorrect(checkUser.Password, password, checkUser.Salt)

	if !isPwdCorrect {
		return user, errors.New("账号或密码错误")
	}

	var status LoginStatus
	status.Phone = checkUser.Phone
	status.Time = time.Now()

	var statusCount int
	GetDb().Table("t_login_status").Where("phone = ?", checkUser.Phone).Count(&statusCount)

	if statusCount <= 0 {
		GetDb().Table("t_login_status").Create(&status)
	} else {
		GetDb().Table("t_login_status").Where("phone = ?", checkUser.Phone).Update(&status)
	}

	GetDb().Table("t_users").Where("phone = ?", checkUser.Phone).Find(&user)

	if user.Phone == "" {
		return user, errors.New("未找到该用户信息")
	}

	return user, nil
}

//发送验证码接口
func SendSmsCode(phone string) (err error) {
	codeStr, sendErr := SendSms(phone)

	if sendErr != nil {
		return errors.New("发送验证码失败")
	}

	var code Code
	code.Code = codeStr
	code.Phone = phone
	code.Time = time.Now()

	var count int
	GetDb().Table("t_codes").Where("phone = ?", phone).Count(&count)

	if count <= 0 {
		err = GetDb().Table("t_codes").Create(&code).Error
	} else {
		err = GetDb().Table("t_codes").Where("phone = ?", phone).Update(&code).Error
	}

	if err != nil {
		return errors.New("更新数据库信息失败")
	}

	return nil
}

//修改密码
func ChangePassword(phone, newPassword, codeStr string) (err error) {
	var code Code
	GetDb().Table("t_codes").Where("phone = ? AND code = ?", phone, codeStr).Find(&code)

	if code.Phone == "" {
		return errors.New("验证玛错误")
	}

	nowTime := time.Now()
	offsetTime := nowTime.Sub(code.Time).Minutes()

	if offsetTime > DEFAULT_OVERDUE_CODE_TIME {
		return errors.New("验证玛过期")
	}

	var user User
	GetDb().Table("t_users").Where("phone = ?", phone).Find(&user)

	if user.Phone == "" {
		return errors.New("该帐号未注册")
	}

	user.Password, user.Salt = TranslateUserPassword(newPassword)

	err = GetDb().Table("t_users").Where("phone = ?", phone).Update(&user).Error
	if err != nil {
		return errors.New("更新密码出错")
	}

	return nil
}

//用户登出
func LoginOut(phone string) (err error) {
	var count int
	GetDb().Table("t_login_status").Where("phone = ?", phone).Count(&count)

	if count <= 0 {
		return errors.New("账号错误")
	}

	err = GetDb().Table("t_login_status").Where("phone = ?", phone).Delete(LoginStatus{}).Error

	if err != nil {
		return errors.New("登出失败")
	}

	return nil
}

//获取用户个人信息
func GetUser(phone string) (user User, err error) {
	GetDb().Table("t_users").Where("phone = ?", phone).Find(&user)

	if user.Phone == "" {
		return user, errors.New("未找到该用户信息")
	}

	return user, nil
}
