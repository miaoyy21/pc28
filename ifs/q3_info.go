package ifs

import (
	"fmt"
	"math"
	"pc28/base"
)

type Info struct {
	Total  int
	Values map[int]float64
	Sqrt   float64
	Min    float64
	Avg    float64
	Max    float64
}

func getInfo(issueId int) (*Info, error) {
	var resp struct {
		Status int `json:"status"`
		Data   struct {
			Items []struct {
				Id      int    `json:"id"`
				C0      int    `json:"c0"`
				C1      int    `json:"c1"`
				C2      int    `json:"c2"`
				C3      int    `json:"c3"`
				C4      int    `json:"c4"`
				C5      int    `json:"c5"`
				C6      int    `json:"c6"`
				C7      int    `json:"c7"`
				C8      int    `json:"c8"`
				C9      int    `json:"c9"`
				C10     int    `json:"c10"`
				C11     int    `json:"c11"`
				C12     int    `json:"c12"`
				C13     int    `json:"c13"`
				C14     int    `json:"c14"`
				C15     int    `json:"c15"`
				C16     int    `json:"c16"`
				C17     int    `json:"c17"`
				C18     int    `json:"c18"`
				C19     int    `json:"c19"`
				C20     int    `json:"c20"`
				C21     int    `json:"c21"`
				C22     int    `json:"c22"`
				C23     int    `json:"c23"`
				C24     int    `json:"c24"`
				C25     int    `json:"c25"`
				C26     int    `json:"c26"`
				C27     int    `json:"c27"`
				TMoney  int    `json:"tmoney"`
				Issue   string `json:"issue"`
				RowNum  int    `json:"rownum"`
				RowNum1 int    `json:"rownum1"`
			} `json:"items"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	if err := Exec("templates/info.tpl", struct {
		IssueId int

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

	values := make(map[int]float64)

	items := resp.Data.Items
	if len(items) != 1 {
		return nil, fmt.Errorf("查询明细异常，其长度为[%d]", len(items))
	}

	item := items[0]
	values[0] = float64(item.TMoney) / float64(item.C0)
	values[1] = float64(item.TMoney) / float64(item.C1)
	values[2] = float64(item.TMoney) / float64(item.C2)
	values[3] = float64(item.TMoney) / float64(item.C3)
	values[4] = float64(item.TMoney) / float64(item.C4)
	values[5] = float64(item.TMoney) / float64(item.C5)
	values[6] = float64(item.TMoney) / float64(item.C6)
	values[7] = float64(item.TMoney) / float64(item.C7)
	values[8] = float64(item.TMoney) / float64(item.C8)
	values[9] = float64(item.TMoney) / float64(item.C9)
	values[10] = float64(item.TMoney) / float64(item.C10)
	values[11] = float64(item.TMoney) / float64(item.C11)
	values[12] = float64(item.TMoney) / float64(item.C12)
	values[13] = float64(item.TMoney) / float64(item.C13)
	values[14] = float64(item.TMoney) / float64(item.C14)
	values[15] = float64(item.TMoney) / float64(item.C15)
	values[16] = float64(item.TMoney) / float64(item.C16)
	values[17] = float64(item.TMoney) / float64(item.C17)
	values[18] = float64(item.TMoney) / float64(item.C18)
	values[19] = float64(item.TMoney) / float64(item.C19)
	values[20] = float64(item.TMoney) / float64(item.C20)
	values[21] = float64(item.TMoney) / float64(item.C21)
	values[22] = float64(item.TMoney) / float64(item.C22)
	values[23] = float64(item.TMoney) / float64(item.C23)
	values[24] = float64(item.TMoney) / float64(item.C24)
	values[25] = float64(item.TMoney) / float64(item.C25)
	values[26] = float64(item.TMoney) / float64(item.C26)
	values[27] = float64(item.TMoney) / float64(item.C27)

	var sqrt2 float64
	min, avg, max := 1.0, 0.0, 0.0
	for no, value := range values {
		sigma := value * float64(base.STDS1000[no]) / 1000
		if sigma > max {
			max = sigma
		}

		if sigma < min {
			min = sigma
		}

		avg = avg + (float64(base.STDS1000[no])/1000)*sigma
		sqrt2 = sqrt2 + (float64(base.STDS1000[no])/1000)*math.Pow(sigma-1.0, 2)
	}

	return &Info{Total: item.TMoney, Values: values, Sqrt: math.Sqrt(sqrt2), Min: min, Avg: avg, Max: max}, nil
}
