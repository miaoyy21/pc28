package ifs

import (
	"fmt"
	"pc28/base"
)

func doRecord() error {
	var resp struct {
		Status int `json:"status"`
		Data   struct {
		} `json:"data"`
		Msg string `json:"msg"`
	}

	if err := Exec("templates/record.tpl", struct {
		UserId   string
		DeviceId string
		Unix     string
		KeyCode  string
		Token    string
	}{
		UserId:   base.Config.UserId,
		DeviceId: base.Config.DeviceId,
		Unix:     base.Config.UnixBetting,
		KeyCode:  base.Config.KeyCodeBetting,
		Token:    base.Config.Token,
	}, &resp); err != nil {
		return err
	}

	if resp.Status != 0 {
		return fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Status, resp.Msg)
	}

	return nil
}
