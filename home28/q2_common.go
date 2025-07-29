package home28

import (
	"fmt"
	"tty28/base"
)

type Common struct {
	NextId     int
	LastId     int
	LastResult int
}

func getCommon() (*Common, error) {
	var resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Current struct {
				Id int `json:"id"`
			} `json:"current"`
			Last struct {
				Id   int `json:"id"`
				Code int `json:"code"`
			} `json:"last"`
		} `json:"data"`
	}

	if err := Exec("templates/common.tpl", struct{ Token string }{Token: base.Config.Token}, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Code, resp.Msg)
	}

	return &Common{
		NextId:     resp.Data.Current.Id,
		LastId:     resp.Data.Last.Id,
		LastResult: resp.Data.Last.Code,
	}, nil
}
