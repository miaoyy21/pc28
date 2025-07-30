package ifs

import (
	"fmt"
	"log"
	"math"
	"pc28/base"
	"strings"
)

func run() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Printf("/********************************** 开始执行定时任务 **********************************/")

	// 获取用户信息
	user, err := getUser()
	if err != nil {
		log.Printf("ERROR : %s", err.Error())
		return
	}

	log.Printf("用户ID【%d】，手机号码【%s】，当前余额【%d】...\n", user.Id, user.Mobile, user.Gold)

	// 获取最新的已开奖及即将开奖信息
	common, err := getCommon()
	if err != nil {
		log.Printf("ERROR : %s", err.Error())
		return
	}

	log.Printf("已开奖期数【%d | %s】，开奖结果【%02d】，即将开奖期数【%d | %s】...\n", common.ThisIssueId, common.ThisIssueNumber, common.ThisResult, common.NextIssueId, common.NextIssueNumber)

	// 即将开奖赔率
	nextIssue, err := getInfo(common.NextIssueId)
	if err != nil {
		log.Printf("ERROR : %s", err.Error())
		return
	}

	log.Printf("即将开奖期数【%d | %s】，累计投注额【%d】...\n", common.NextIssueId, common.NextIssueNumber, nextIssue.Total)

	bets, total := make([]int, 0, len(base.SN28)), 0
	for _, no := range base.SN28 {
		sigma := nextIssue.Values[no] / (1000 / float64(base.STDS1000[no]))

		var delta float64
		if sigma > base.Config.Sigma {
			if sigma <= 1.0 {
				delta = (sigma - base.Config.Sigma) / (1.0 - base.Config.Sigma)
			} else {
				delta = math.Pow(base.Config.Enigma, sigma)
			}
		}

		bet := int(delta * float64(base.Config.Base) * float64(base.STDS1000[no]) / 1000)
		if bet <= 0 {
			log.Printf("  【   】数字【%02d】，赔率【%-8.2f】，赔率系数【%4.2f】...\n", no, nextIssue.Values[no], sigma)
		} else {
			log.Printf("  【 ✓ 】数字【%02d】，赔率【%-8.2f】，赔率系数【%4.2f】...\n", no, nextIssue.Values[no], delta)
		}

		total = total + bet
		bets = append(bets, bet)
	}

	sBets := make([]string, 0, len(bets))
	for _, bet := range bets {
		sBets = append(sBets, fmt.Sprintf("%d", bet))
	}

	// 执行投注
	log.Printf("sBets is %s\n", strings.Join(sBets, ","))
	log.Println()
	//if err := doBet(common.NextIssueNumber, strings.Join(sBets, ","), total); err != nil {
	//	log.Printf("ERROR : %s", err.Error())
	//	return
	//}

	log.Printf("/********************************** 投注已完成，投注金额【%-7d】 **********************************/\n", total)
}
