package pc28

import (
	"fmt"
	"pc28/base"
	"pc28/ifs/exec"
)

type Index struct {
	UserEggs int

	ThisIssueId string
	ThisResult  string

	NextIssueId    string
	NextIssueMoney int
}

func getIndex() (*Index, error) {
	var resp struct {
		Status int `json:"status"`
		Data   struct {
			Info []struct {
				UserEggs int `json:"usereggs"`
			} `json:"info"`
			KList []struct {
				Issue  string `json:"issue"`
				TMoney int    `json:"tmoney"`
			} `json:"klist"`
			Plist []struct {
				Issue   string `json:"issue"`
				LResult string `json:"lresult"`
			} `json:"plist"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	if err := exec.Exec("templates/bowling_index.tpl",
		struct {
			UserId   string
			DeviceId string
			Unix     string
			KeyCode  string
			Token    string
		}{
			UserId:   base.Config.UserId,
			DeviceId: base.Config.DeviceId,
			Unix:     base.Config.UnixNormal,
			KeyCode:  base.Config.KeyCodeNormal,
			Token:    base.Config.Token,
		},
		&resp,
	); err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Status, resp.Msg)
	}

	return &Index{
		UserEggs:       resp.Data.Info[0].UserEggs,
		ThisIssueId:    resp.Data.Plist[0].Issue,
		ThisResult:     resp.Data.Plist[0].LResult,
		NextIssueId:    resp.Data.KList[0].Issue,
		NextIssueMoney: resp.Data.KList[0].TMoney,
	}, nil
}
