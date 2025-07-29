package home28

import (
	"fmt"
	"strconv"
	"strings"
	"tty28/base"
)

type User struct {
	Id   int
	Name string
	Gold int
}

func getUser() (*User, error) {
	var resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Id   int    `json:"id"`
			Name string `json:"uname"`
			Gold string `json:"gold"`
		} `json:"data"`
	}

	if err := Exec("templates/user.tpl", struct{ Token string }{Token: base.Config.Token}, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Code, resp.Msg)
	}

	gold, err := strconv.Atoi(strings.Replace(resp.Data.Gold, ",", "", -1))
	if err != nil {
		return nil, err
	}

	return &User{
		Id:   resp.Data.Id,
		Name: resp.Data.Name,
		Gold: gold,
	}, nil
}
