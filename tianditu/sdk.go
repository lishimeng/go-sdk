package tianditu

import (
	"errors"
	"fmt"
)

const (
	defaultHost = "api.tianditu.gov.cn"
	geoCoderTpl = `%s/geocoder?postStr=%s&type=geocode&tk=%s`
)

var (
	ErrHttpCode = errors.New("http err")
)

type Client interface {
	Geo2Address(longitude, latitude float64) (resp Geo2AddressResponse, err error) // 地址逆解析
}

type SimpleClient struct {
	Host   string // 服务器, 默认 defaultHost
	UseSSL bool   // 使用https
	Key    string // 验证key
}

type Option func(*SimpleClient)

// WithHost 自定义服务器地址,  默认 defaultHost, host不带schema
func WithHost(host string) Option {
	return func(c *SimpleClient) {
		c.Host = host
	}
}

// WithSSL 是否使用https, 默认不使用
func WithSSL(useSSL bool) Option {
	return func(c *SimpleClient) {
		c.UseSSL = useSSL
	}
}

// WithKey 设置验证key
func WithKey(key string) Option {
	return func(c *SimpleClient) {
		c.Key = key
	}
}

func NewClient(opts ...Option) Client {
	c := &SimpleClient{Host: defaultHost, UseSSL: false}
	for _, opt := range opts {
		opt(c)
	}
	c.buildHost()
	var client Client = c
	return client
}

func (c *SimpleClient) buildHost() {
	schema := "http"
	if c.UseSSL {
		schema = "https"
	}
	c.Host = fmt.Sprintf("%s://%s", schema, c.Host)
}
