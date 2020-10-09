package models

//openid解密后返回的结构体
type XWxIDResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

//accessToken解密后返回的结构体
type WxTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

//openid解密后返回的结构体
type WxIDResponse struct {
	OpenID       string `json:"openid"`
	AccessToken  string `json:"access_token"`
	ExpriesIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type WxInfoResponse struct {
	OpenID   string `json:"openid"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"`
	Province string `json:"province"`
	City     string `json:"city"`
	Country  string `json:"country"`
	ImageURL string `json:"headimgurl"`
	Unionid  string `json:"unionid"`
}
