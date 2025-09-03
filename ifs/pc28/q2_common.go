package pc28

import (
	"fmt"
	"pc28/base"
	"pc28/ifs/exec"
	"strconv"
)

type Common struct {
	ThisIssueId     int
	ThisIssueNumber string
	ThisResult      int
	NextIssueId     int
	NextIssueNumber string
}

func getCommon() (*Common, error) {
	var resp struct {
		Status int `json:"status"`
		Data   struct {
			Items []struct {
				Id               int    `json:"id"`
				Number           string `json:"number"`
				NowTime          string `json:"nowtime"`
				OpenTime         string `json:"opentime"`
				Countdown        int    `json:"countdown"`
				OpenDown         int    `json:"opendown"`
				Invest           int    `json:"isinvest"`
				IsOpen           int    `json:"isopen"`
				OpenResultDetail string `json:"openresultdetail"`
				Detail1          string `json:"detail1"`
				Detail2          string `json:"detail2"`
				Detail3          string `json:"detail3"`
				OpenResult       string `json:"openresult"`
				TMoney           int    `json:"tmoney"`
				WMember          int    `json:"wmember"`
				TMoneyAll        int    `json:"tmoney_all"`
				UtMoney          int    `json:"utmoney"`
				Rn               int    `json:"rn"`
			} `json:"items"`
			Issue []struct {
				LastIssue  string `json:"lastissue"`
				TotalCount int    `json:"totalcount"`
			} `json:"issue"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	if err := exec.Exec("templates/pc28_common.tpl", struct {
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
	}, &resp); err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, fmt.Errorf("错误代码 [%d] ，错误信息[%s]", resp.Status, resp.Msg)
	}

	var thisIssueId int
	var thisIssueNumber string
	var thisResult int
	var thisIndex int

	items := resp.Data.Items
	for index, item := range items {
		if item.IsOpen == 1 {
			thisIndex = index
			thisIssueId = item.Id
			thisIssueNumber = item.Number

			iResult, err := strconv.Atoi(item.OpenResult)
			if err != nil {
				return nil, fmt.Errorf("查询历史异常，[%s]不是一个整数类型的开奖结果", item.OpenResult)
			}
			thisResult = iResult

			break
		}
	}

	if thisIndex < 1 {
		return nil, fmt.Errorf("查询历史异常，当前期数的索引为[%d]", thisIndex)
	}

	return &Common{
		ThisIssueId:     thisIssueId,
		ThisIssueNumber: thisIssueNumber,
		ThisResult:      thisResult,
		NextIssueId:     items[thisIndex-1].Id,
		NextIssueNumber: items[thisIndex-1].Number,
	}, nil
}
