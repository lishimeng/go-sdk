package tianditu

import (
	"errors"
	"fmt"
	"github.com/lishimeng/go-sdk/tianditu/rest"
	"net/url"
)

type Geo2AddressResponse struct {
	Status  string           `json:"status,omitempty"`
	Message string           `json:"msg,omitempty"`
	Result  GeoAddressResult `json:"result,omitempty"`
}

type GeoAddressResult struct {
	FormattedAddress string           `json:"formatted_address,omitempty"`
	Location         Location         `json:"location,omitempty"` // 详细地址
	AddressComponent AddressComponent `json:"addressComponent,omitempty"`
}

// AddressComponent 此点的具体信息（分类）
type AddressComponent struct {
	Address         string `json:"address,omitempty"`          // 此点最近地点信息
	AddressDistance int    `json:"address_distance,omitempty"` // 此点距离最近地点信息距离
	AddressPosition string `json:"address_position,omitempty"` // 此点在最近地点信息方向
	City            string `json:"city,omitempty"`             //此点所在国家或城市或区县
	Poi             string `json:"poi,omitempty"`              // 距离此点最近poi点
	PoiDistance     int    `json:"poi_distance	,omitempty"`    // 距离此点最近poi点的距离
	PoiPosition     string `json:"poi_position,omitempty"`     // 此点在最近poi点的方向
	Road            string `json:"road,omitempty"`             //距离此点最近的路
	RoadDistance    int    `json:"road_distance,omitempty"`    // 此点距离此路的距离
}

type Location struct {
	Lon float64 `json:"lon,omitempty"` // 此点坐标x值
	Lat float64 `json:"lat,omitempty"` // 此点坐标y值
}

func (c *SimpleClient) Geo2Address(longitude, latitude float64) (resp Geo2AddressResponse, err error) {
	const queryTpl = `{'lon':%f,'lat':%f,'ver':1}`
	geoParam := fmt.Sprintf(queryTpl, longitude, latitude)
	geoParam = url.QueryEscape(geoParam)

	urlPath := fmt.Sprintf(geoCoderTpl, c.Host, geoParam, c.Key)
	code, err := rest.Post(urlPath, &resp,
		rest.Header{
			Name:  "User-Agent",
			Value: `Apifox/1.0.0 (https://apifox.com)`,
		},
		rest.Header{
			Name:  "Accept",
			Value: "*/*",
		},
		rest.Header{
			Name:  "Host",
			Value: "api.tianditu.gov.cn",
		})
	if err != nil {
		return
	}
	if resp.Status != "0" {
		err = errors.New(resp.Message)
		return
	}
	if code != 200 {
		err = ErrHttpCode
		return
	}
	return
}
