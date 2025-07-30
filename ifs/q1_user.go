package ifs

import (
	"fmt"
	"pc28/base"
)

type User struct {
	Id     int
	Mobile string
	Gold   int
}

func getUser() (*User, error) {
	var resp struct {
		Status int `json:"status"`
		Data   struct {
			UserId         int           `json:"userid"`
			GoldEggs       int           `json:"goldeggs"`
			Cashbook       int           `json:"cashbox"`
			Auto28         int           `json:"auto28"`
			Auto28Count    int           `json:"auto28count"`
			Mobile         string        `json:"mobileno"`
			MobileStatus   int           `json:"mobilestatus"`
			IssueCashpoint int           `json:"issetcashpsw"`
			VisitorType    int           `json:"visitortype"`
			HeadImg        string        `json:"headimg"`
			BannerList     []interface{} `json:"bannerList"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	if err := Exec("templates/user.tpl",
		struct {
			UserId   string
			DeviceId string
			Unix     string
			KeyCode  string
			Token    string
		}{
			UserId:   base.Config.UserId,
			DeviceId: base.Config.DeviceId,
			Unix:     base.Config.SUnix,
			KeyCode:  base.Config.SKeyCode,
			Token:    base.Config.Token,
		},
		&resp,
	); err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Status, resp.Msg)
	}

	return &User{Id: resp.Data.UserId, Mobile: resp.Data.Mobile, Gold: resp.Data.GoldEggs}, nil
}
