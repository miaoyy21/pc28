package client

import (
	"fmt"
	"log"
	"math"
	"pc28/hdo"
	"strconv"
)

type QRiddleRequest struct {
	Issue     string `json:"issue"`
	Unix      string `json:"unix"`
	Keycode   string `json:"keycode"`
	PType     string `json:"ptype"`
	DeviceId  string `json:"deviceid"`
	ChannelId string `json:"channelid"`
	Userid    string `json:"userid"`
	Token     string `json:"token"`
}

type QRiddleResponse struct {
	Status int `json:"status"`
	Data   struct {
		Riddle []struct {
			Num  string `json:"num"`
			Rate string `json:"rate"`
		} `json:"myriddle"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func qRiddle(issue string) (map[int32]float64, float64, float64, error) {
	req := &QRiddleRequest{
		Issue:     issue,
		Unix:      conf.Unix,
		Keycode:   conf.KeyCode,
		PType:     conf.PType,
		DeviceId:  conf.DeviceId,
		ChannelId: conf.ChannelId,
		Userid:    conf.UserId,
		Token:     conf.Token,
	}

	var resp QRiddleResponse

	err := hdo.Do(conf.Origin, conf.Cookie, conf.UserAgent, conf.RiddleURL, req, &resp)
	if err != nil {
		return nil, 0, 0, err
	}

	if resp.Status != 0 {
		return nil, 0, 0, fmt.Errorf("接收到状态错误吗 : [%d] %s", resp.Status, resp.Msg)
	}

	var exp float64
	rts := make(map[int32]float64)
	for _, riddle := range resp.Data.Riddle {
		n, err := strconv.Atoi(riddle.Num)
		if err != nil {
			return nil, 0, 0, err
		}

		r, err := strconv.ParseFloat(riddle.Rate, 64)
		if err != nil {
			return nil, 0, 0, err
		}

		rx := r / (1000.0 / float64(STDS1000[int32(n)]))
		log.Printf("  竞猜数字【%02d】：实际赔率【%7.2f】，赔率系数【%.3f】； \n", n, r, rx)

		rts[int32(n)] = rx
		exp = exp + rx/28
	}

	var dev float64
	for _, rx := range rts {
		dev = dev + (rx-exp)*(rx-exp)/28
	}

	log.Printf("  赔率系数的数学期望为【%.4f】，标准方差为【%.4f】 \n", exp, math.Sqrt(dev))
	return rts, exp, math.Sqrt(dev), nil
}
