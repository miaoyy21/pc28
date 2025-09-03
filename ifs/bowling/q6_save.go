package bowling

import (
	"fmt"
	"pc28/base"
	"pc28/ifs/exec"
)

func doSave(issueId string, bets map[int]int) error {
	var resp struct {
		Status int      `json:"status"`
		Data   struct{} `json:"data"`
		Msg    string   `json:"msg"`
	}

	if err := exec.Exec("templates/bowling_save.tpl", struct {
		IssueId string

		Key0  int
		Key1  int
		Key2  int
		Key3  int
		Key4  int
		Key5  int
		Key6  int
		Key7  int
		Key8  int
		Key9  int
		Key10 int
		Key11 int
		Key12 int
		Key13 int
		Key14 int
		Key15 int
		Key16 int
		Key17 int
		Key18 int
		Key19 int
		Key20 int
		Key21 int
		Key22 int
		Key23 int
		Key24 int
		Key25 int
		Key26 int
		Key27 int

		UserId   string
		DeviceId string
		Unix     string
		KeyCode  string
		Token    string
	}{
		IssueId: issueId,

		Key0:  bets[0],
		Key1:  bets[1],
		Key2:  bets[2],
		Key3:  bets[3],
		Key4:  bets[4],
		Key5:  bets[5],
		Key6:  bets[6],
		Key7:  bets[7],
		Key8:  bets[8],
		Key9:  bets[9],
		Key10: bets[10],
		Key11: bets[11],
		Key12: bets[12],
		Key13: bets[13],
		Key14: bets[14],
		Key15: bets[15],
		Key16: bets[16],
		Key17: bets[17],
		Key18: bets[18],
		Key19: bets[19],
		Key20: bets[20],
		Key21: bets[21],
		Key22: bets[22],
		Key23: bets[23],
		Key24: bets[24],
		Key25: bets[25],
		Key26: bets[26],
		Key27: bets[27],

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
