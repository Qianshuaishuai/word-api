package helper

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

//检查sign是否对等
func Check(request *http.Request) bool {
	signkeyPre := ""
	signkeySuffix := ""
	signVersion := ""
	//获取待加密的请求参数
	request.ParseForm()
	sign := request.Form.Get("F_sign")
	if len(sign) <= 2 {
		return false
	}
	request.Form.Del("F_sign")
	//获取加密版本
	signVersion = sign[0:2]
	signVersionErr := false
	switch signVersion {
	case "01":
	case "02":
		method := strings.ToUpper(request.Method)
		signkeyPre = method + "&" + url.QueryEscape("/") + "&"
	default:
		signVersionErr = true
	}
	if signVersionErr {
		return false
	}
	//获取待加密的请求参数
	paramesEncode := ""
	requestParameKey := make([]string, len(request.Form))
	if len(requestParameKey) <= 0 {
		return false
	}
	i := 0
	for k, _ := range request.Form {
		requestParameKey[i] = k
		i++
	}
	sort.Strings(requestParameKey)
	for _, k := range requestParameKey {
		paramesEncode += url.QueryEscape(k) + "=" + url.QueryEscape(request.Form[k][0]) + "&"
	}
	if len(paramesEncode) > 0 {
		paramesEncode = paramesEncode[0 : len(paramesEncode)-1]
		paramesEncode = strings.Replace(paramesEncode, "+", "%20", -1)
	}

	//获取加密的key
	signKey := request.Form.Get("F_accesstoken")
	if len(signKey) <= 0 {
		return false
	}
	signKey = signkeyPre + signKey + signkeySuffix
	//hmac加密
	mac := hmac.New(sha1.New, []byte(signKey))
	mac.Write([]byte(paramesEncode))
	expectedMAC := mac.Sum(nil)
	//base64加密
	sign2 := signVersion + base64.URLEncoding.EncodeToString(expectedMAC)
	if sign2 == sign {
		return true
	}
	return false
}
