package bowling

import (
	"fmt"
	"math"
	"pc28/base"
	"pc28/ifs/exec"
	"strconv"
)

type Detail struct {
	Sqrt   float64
	Values map[int]float64
}

func getDetail(issueId string) (*Detail, error) {
	var resp struct {
		Status int `json:"status"`
		Data   struct {
			MyRiddle []struct {
				Num  string `json:"num"`
				Rate string `json:"rate"`
			} `json:"myriddle"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	if err := exec.Exec("templates/bowling_detail.tpl", struct {
		IssueId string

		UserId   string
		DeviceId string
		Unix     string
		KeyCode  string
		Token    string
	}{
		IssueId: issueId,

		UserId:   base.Config.UserId,
		DeviceId: base.Config.DeviceId,
		Unix:     base.Config.UnixOdds,
		KeyCode:  base.Config.KeyCodeOdds,
		Token:    base.Config.Token,
	}, &resp); err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Status, resp.Msg)
	}

	var sqrt2 float64

	values := make(map[int]float64)
	for _, riddle := range resp.Data.MyRiddle {
		num, err := strconv.Atoi(riddle.Num)
		if err != nil {
			return nil, err
		}

		rate, err := strconv.ParseFloat(riddle.Rate, 64)
		if err != nil {
			return nil, err
		}

		values[num] = rate
		sqrt2 = sqrt2 + (float64(base.STDS1000[num])/1000)*math.Pow(rate-1.0, 2)
	}

	return &Detail{Sqrt: sqrt2, Values: values}, nil
}
