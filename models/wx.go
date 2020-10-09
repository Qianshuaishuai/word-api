package models

import (
	"errors"
	"log"
)

func GetOpenID(code string) (data WxIDResponse, err error) {
	query := make(map[string]string)
	query["appid"] = MyConfig.WxAppID
	query["secret"] = MyConfig.WxAppSecret
	query["code"] = code
	query["grant_type"] = "authorization_code"

	err = wxClient.Get("/sns/oauth2/access_token", query, &data)

	if err != nil {
		log.Println("err1:", err)
		return data, errors.New("请求微信后台服务出错！")
	}

	if data.OpenID == "" {
		log.Println("WxAppSecret:", MyConfig.WxAppSecret)
		log.Println("WxAppID:", MyConfig.WxAppID)
		log.Println("code:", code)
		log.Println("data:", data)
		log.Println("err2:", err)
		return data, errors.New("code过期")
	}

	return data, nil
}

func GetWxUserInfo(accessToken, openID string) (data WxInfoResponse, err error) {
	query := make(map[string]string)
	query["access_token"] = accessToken
	query["openid"] = openID
	query["lang"] = "zh_CN"

	err = wxClient.Get("/sns/userinfo", query, &data)
	if err != nil {
		return data, errors.New("请求微信后台服务出错！")
	}

	if data.OpenID == "" {
		return data, errors.New("token错误")
	}

	return data, nil
}
