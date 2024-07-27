package client

import (
	"fmt"
	"pc28/hdo"
)

type QIssueResponse struct {
	Status int `json:"status"`
	Data   struct {
		TMoney  int `json:"tmoney"`
		TMember int `json:"tmember"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type QIssueRequest struct {
	Issue     string `json:"issue"`
	Unix      string `json:"unix"`
	KeyCode   string `json:"keycode"`
	PType     string `json:"ptype"`
	DeviceId  string `json:"deviceid"`
	UserId    string `json:"userid"`
	Token     string `json:"token"`
	ChannelId string `json:"channelid"`
}

func qIssue(issue string) (int, int, error) {
	req := QIssueRequest{
		Issue:     issue,
		Unix:      conf.Unix,
		KeyCode:   conf.KeyCode,
		PType:     conf.PType,
		DeviceId:  conf.DeviceId,
		UserId:    conf.UserId,
		Token:     conf.Token,
		ChannelId: conf.ChannelId,
	}

	var resp QIssueResponse

	err := hdo.Do(conf.Origin, conf.Cookie, conf.UserAgent, conf.IssueURL, req, &resp)
	if err != nil {
		return 0, 0, err
	}

	if resp.Status != 0 {
		return 0, 0, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Status, resp.Msg)
	}

	return resp.Data.TMoney, resp.Data.TMember, nil
}
