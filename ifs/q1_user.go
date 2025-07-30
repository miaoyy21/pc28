package ifs

import (
	"fmt"
	"pc28/base"
)

type User struct {
	Id   int
	Gold int
}

func getUser() (*User, error) {
	var resp struct {
		Status int `json:"status"`
		Data   struct {
			Id   int `json:"userid"`
			Gold int `json:"goldeggs"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	if err := Exec("templates/user.tpl",
		struct {
			UserId   string
			DeviceId string
			Unix     string
			Token    string
			KeyCode  string
		}{
			UserId:   base.Config.UserId,
			DeviceId: base.Config.DeviceId,
			Unix:     base.Config.Unix,
			Token:    base.Config.Token,
			KeyCode:  base.Config.KeyCode,
		},
		&resp,
	); err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Status, resp.Msg)
	}

	return &User{Id: resp.Data.Id, Gold: resp.Data.Gold}, nil
}
