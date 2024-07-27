package client

import (
	"errors"
	"fmt"
	"pc28/hdo"
	"strconv"
	"strings"
)

type QHistoryResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`

	Data struct {
		Items []struct {
			Result string `json:"lresult"`
			Issue  string `json:"issue"`
			Money  string `json:"tmoney"`
		} `json:"items"`
	} `json:"data"`
}

type QHistoryRequest struct {
	PageSize  int    `json:"pagesize"`
	Unix      string `json:"unix"`
	KeyCode   string `json:"keycode"`
	PType     string `json:"ptype"`
	DeviceId  string `json:"deviceid"`
	ChannelId string `json:"channelid"`
	UserId    string `json:"userid"`
	Token     string `json:"token"`
}

func qHistory() (int, int, int, error) {
	req := QHistoryRequest{
		PageSize:  200,
		PType:     conf.PType,
		Unix:      conf.Unix,
		KeyCode:   conf.KeyCode,
		DeviceId:  conf.DeviceId,
		ChannelId: conf.ChannelId,
		UserId:    conf.UserId,
		Token:     conf.Token,
	}

	var resp QHistoryResponse

	err := hdo.Do(conf.Origin, conf.Cookie, conf.UserAgent, conf.HistoryURL, req, &resp)
	if err != nil {
		return 0, 0, 0, err
	}

	if resp.Status != 0 {
		return 0, 0, 0, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Status, resp.Msg)
	}

	if len(resp.Data.Items) < 1 {
		return 0, 0, 0, errors.New("没有收到返回结果")
	}

	issue, err := strconv.Atoi(resp.Data.Items[0].Issue)
	if err != nil {
		return 0, 0, 0, err
	}

	gold, err := strconv.Atoi(strings.ReplaceAll(resp.Data.Items[0].Money, ",", ""))
	if err != nil {
		return 0, 0, 0, err
	}

	result, err := strconv.Atoi(resp.Data.Items[0].Result)
	if err != nil {
		return 0, 0, 0, err
	}

	return issue, gold, result, nil
}
