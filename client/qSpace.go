package client

import (
	"errors"
	"fmt"
	"pc28/hdo"
)

//curl 'http://manorapp.pceggs.com/IFS/Manor28/Manor28_SpaceList.ashx' \
//-H 'Accept: application/json, text/plain, */*' \
//-H 'Accept-Language: zh-CN,zh;q=0.9' \
//-H 'Cache-Control: no-cache' \
//-H 'Content-Type: application/json;charset=UTF-8' \
//-H 'Cookie: ' \
//-H 'Origin: http://manorapp.pceggs.com' \
//-H 'Pragma: no-cache' \
//-H 'Proxy-Connection: keep-alive' \
//-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36' \
//--data-raw '{"unix":"1716518900","keycode":"7233ecd412c146116da02b2a1275dcea","ptype":"3","deviceid":"0E6EE3CC-8184-4CD7-B163-50AE8AD4516F","userid":"31591499","token":"ahrfdtc2wagitbggeca7ajtjhlw7l7sqz7qfqjp8","channelid":"0"}' \
//--insecure

type QSpaceResponse struct {
	Status int `json:"status"`
	Data   struct {
		SpeacList []struct {
			Nums     int `json:"nums"`
			Nowspeac int `json:"nowspeac"`
			Defspeac int `json:"defspeac"`
		} `json:"SpeacList"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type QSpaceRequest struct {
	Unix      string `json:"unix"`
	KeyCode   string `json:"keycode"`
	PType     string `json:"ptype"`
	DeviceId  string `json:"deviceid"`
	ChannelId string `json:"channelid"`
	UserId    string `json:"userid"`
	Token     string `json:"token"`
}

func qSpace() (map[int32]float64, error) {
	req := QSpaceRequest{
		PType:     conf.PType,
		Unix:      conf.Unix,
		KeyCode:   conf.KeyCode,
		DeviceId:  conf.DeviceId,
		ChannelId: conf.ChannelId,
		UserId:    conf.UserId,
		Token:     conf.Token,
	}

	var resp QSpaceResponse

	err := hdo.Do(conf.Origin, conf.Cookie, conf.UserAgent, conf.SpaceURL, req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Status, resp.Msg)
	}

	if len(resp.Data.SpeacList) < 1 {
		return nil, errors.New("没有收到返回结果")
	}

	rts := make(map[int32]float64)
	for _, d := range resp.Data.SpeacList {
		rts[int32(d.Nums)] = float64(d.Nowspeac) / float64(d.Defspeac)
	}

	return rts, nil
}
