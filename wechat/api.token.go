package wechat

import (
	"errors"
	"fmt"
	"github.com/lishimeng/wechat/rest"
	"net/url"
	"time"
)

func (c *Client) getToken() (token ApiAccessToken, err error) {
	apiUrl, err := url.JoinPath(c.Host, "/cgi-bin/token")
	if err != nil {
		return
	}

	const apiTpl = `%s?appid=%s&secret=%s&grant_type=%s`

	apiUrl = fmt.Sprintf(apiTpl, apiUrl, c.AppId, c.Secret, "client_credential")

	httpCode, err := rest.New().GetJson(apiUrl, nil, &token)
	if err != nil {
		return
	}
	if httpCode != 200 {
		err = errors.New("http code is not 200")
		return
	}

	if token.ErrCode != 0 {
		err = errors.New(token.ErrMsg)
		return
	}

	token.ExpiresAt = time.Now().Add(time.Second * time.Duration(token.ExpiresIn-200)).Unix()
	c.AccessToken = token

	return
}
