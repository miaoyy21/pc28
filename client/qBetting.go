package client

import (
	"fmt"
	"math/rand"
	"pc28/hdo"
	"time"
)

type QBettingResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type QBettingRequest struct {
	Issue     string `json:"issue"`
	GoldEggs  int32  `json:"totalgoldeggs"`
	Numbers   int32  `json:"numbers"`
	Unix      string `json:"unix"`
	Keycode   string `json:"keycode"`
	PType     string `json:"ptype"`
	DeviceId  string `json:"deviceid"`
	ChannelId string `json:"channelid"`
	Userid    string `json:"userid"`
	Token     string `json:"token"`
}

func qBetting(issue string, bets map[int32]int32) error {
	req := QBettingRequest{
		Issue:     issue,
		Unix:      conf.Unix,
		Keycode:   conf.KeyCode,
		PType:     conf.PType,
		DeviceId:  conf.DeviceId,
		ChannelId: conf.ChannelId,
		Userid:    conf.UserId,
		Token:     conf.Token,
	}

	for i, g := range bets {
		var resp QBettingResponse

		req.Numbers = i
		req.GoldEggs = g

		time.Sleep(time.Second * time.Duration(0.25*rand.Float64()))
		err := hdo.Do(conf.Origin, conf.Cookie, conf.UserAgent, conf.BettingURL, req, &resp)
		if err != nil {
			return err
		}

		if resp.Status != 0 {
			return fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Status, resp.Msg)
		}
	}

	return nil
}
