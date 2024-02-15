package wechat

import "time"

type ApiAccessToken struct {
	Err
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ExpiresAt   int64  `json:"expires_at"`
}

// Expired 判断token是否过期
func (a *ApiAccessToken) Expired() bool {
	return a.ExpiresAt < time.Now().Unix()
}

type UserInfo struct {
	Openid string
}

type AppInfo struct {
	AppId  string `json:"appid,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type JsSessionReq struct {
	AppInfo
	Code      string `json:"js_code,omitempty"`
	GrantType string `json:"grant_type,omitempty"`
}

// JsSession 微信小程序js登录返回的会话信息
type JsSession struct {
	SessionKey string `json:"session_key,omitempty"`
	OpenInfo
	Err
}

type OpenInfo struct {
	Openid  string `json:"openid,omitempty"`
	UnionId string `json:"unionid,omitempty"`
}

type Err struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

type PhoneNumResp struct {
	Err
	PhoneInfo PhoneInfo `json:"phone_info,omitempty"` // 用户手机号信息
}

type PhoneWaterMarker struct {
	Timestamp int64  `json:"timestamp,omitempty"` // 用户获取手机号操作的时间戳
	Appid     string `json:"appid,omitempty"`     //小程序appid
}

type PhoneInfo struct {
	PhoneNumber     string           `json:"phoneNumber,omitempty"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string           `json:"purePhoneNumber,omitempty"` // 没有区号的手机号
	CountryCode     string           `json:"countryCode,omitempty"`     // 区号
	Watermark       PhoneWaterMarker `json:"watermark,omitempty"`       // 数据水印
}
