package home28

import (
	"fmt"
	"tty28/base"
)

func doBet(id int, bets string) error {
	var resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	if err := Exec("templates/bet.tpl", struct {
		Token string
		Id    int
		Bets  string
	}{Token: base.Config.Token, Id: id, Bets: bets}, &resp); err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Code, resp.Msg)
	}

	return nil
}
