package ifs

import (
	"fmt"
	"pc28/base"
)

func doMode(name string, sBets string) error {
	var resp struct {
		Status int      `json:"status"`
		Data   struct{} `json:"data"`
		Msg    string   `json:"msg"`
	}

	if err := Exec("templates/mode.tpl", struct {
		SBets string
		Name  string

		UserId   string
		DeviceId string
		Unix     string
		KeyCode  string
		Token    string
	}{
		SBets: sBets,
		Name:  name,

		UserId:   base.Config.UserId,
		DeviceId: base.Config.DeviceId,
		Unix:     base.Config.UnixMode,
		KeyCode:  base.Config.KeyCodeMode,
		Token:    base.Config.Token,
	}, &resp); err != nil {
		return err
	}

	if resp.Status != 0 {
		return fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Status, resp.Msg)
	}

	return nil
}
