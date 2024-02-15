package wechat

import (
	"errors"
	"fmt"
	"github.com/lishimeng/x/rest"
	"net/url"
)

type Client struct {
	AppId       string
	Secret      string
	Host        string
	AccessToken ApiAccessToken
}

func New(appId, secret string) (c *Client) {
	c = &Client{
		AppId:  appId,
		Secret: secret,
		Host:   defaultHost,
	}
	return
}

const defaultHost = "https://api.weixin.qq.com"

const (
	jscodeFormat = "%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	phoneNumTpl  = `%s?access_token=%s`
)

func (c *Client) JsCode2Session(code string) (session JsSession, err error) {
	// 调用微信API获取session_key和openid
	// ...
	apiUrl, err := url.JoinPath(c.Host, "/sns/jscode2session")
	if err != nil {
		return
	}

	apiUrl = fmt.Sprintf(jscodeFormat, apiUrl, c.AppId, c.Secret, code)

	httpCode, err := rest.New().GetJson(apiUrl, nil, &session)
	if err != nil {
		return
	}
	if httpCode != 200 {
		err = errors.New("http code is not 200")
		return
	}

	return
}

func (c *Client) GetPhoneNumber(code string) (resp PhoneNumResp, err error) {
	if c.AccessToken.Expired() {
		_, err = c.getToken()
		if err != nil {
			return
		}
	}
	apiUrl, err := url.JoinPath(c.Host, "wxa/business/getuserphonenumber")
	if err != nil {
		return
	}
	apiUrl = fmt.Sprintf(phoneNumTpl, apiUrl, c.AccessToken.AccessToken)

	req := map[string]string{"code": code}
	httpCode, err := rest.New().PostJson(apiUrl, req, &resp)
	if err != nil {
		return
	}
	if httpCode != 200 {
		err = errors.New("http code is not 200")
		return
	}
	return
}

func GenerateToken(openid string) (string, error) {
	// 使用openid生成JWT token
	// ...

	// 返回JWT token
	return "token", nil
}

func VerifyToken(token string) (string, error) {
	// 验证JWT token的有效性
	// ...

	// 返回openid
	return "openid", nil
}

func GetUserInfo(openid string) (*UserInfo, error) {
	// 根据openid从数据库中获取用户信息
	// ...

	// 返回用户信息
	return &UserInfo{
		Openid: openid,
		// 其他字段...
	}, nil
}

func GetUserInfoByToken(token string) (*UserInfo, error) {
	// 根据token从数据库中获取用户信息
	// ...

	// 返回用户信息
	return &UserInfo{
		Openid: "openid",
		// 其他字段...
	}, nil
}
