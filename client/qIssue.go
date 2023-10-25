package client

import (
	"errors"
	"fmt"
	"pc28/hdo"
	"strconv"
)

type QIssueItem struct {
	Issue  string `json:"issue"`
	Result string `json:"lresult"`
	Money  string `json:"tmoney"`
	Member int    `json:"tmember"`
}

type QIssueData struct {
	Items []QIssueItem `json:"items"`
}

type QIssueResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`

	Data QIssueData `json:"data"`
}

type QIssueRequest struct {
	PageSize  int    `json:"pagesize"`
	Unix      string `json:"unix"`
	KeyCode   string `json:"keycode"`
	PType     string `json:"ptype"`
	DeviceId  string `json:"deviceid"`
	ChannelId string `json:"channelid"`
	UserId    string `json:"userid"`
	Token     string `json:"token"`
}

func qIssue() (int, error) {
	req := QIssueRequest{
		PageSize:  200,
		PType:     conf.PType,
		Unix:      conf.Unix,
		KeyCode:   conf.KeyCode,
		DeviceId:  conf.DeviceId,
		ChannelId: conf.ChannelId,
		UserId:    conf.UserId,
		Token:     conf.Token,
	}

	var resp QIssueResponse

	err := hdo.Do(conf.Origin, conf.Cookie, conf.UserAgent, conf.IssueURL, req, &resp)
	if err != nil {
		return 0, err
	}

	if resp.Status != 0 {
		return 0, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Status, resp.Msg)
	}

	if len(resp.Data.Items) < 1 {
		return 0, errors.New("没有收到返回结果")
	}

	return strconv.Atoi(resp.Data.Items[0].Issue)
}
