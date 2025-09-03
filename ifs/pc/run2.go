package pc

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/url"
	"pc28/base"
	"strings"
)

func run2() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Printf("/********************************** 开始执行定时任务 **********************************/")

	// 获取用户信息
	//base.Sleep(rand.Float64() * 30)
	//user, err := getUser()
	//if err != nil {
	//	log.Printf("getUser() ERROR : %s", err.Error())
	//	return
	//}
	//
	//log.Printf("用户ID【%d】，手机号码【%s】，当前余额【%d】...\n", user.Id, user.Mobile, user.Gold)

	// 获取最新的已开奖及即将开奖信息
	base.Sleep(rand.Float64() * 30)
	common, err := getCommon()
	if err != nil {
		log.Printf("getCommon() ERROR : %s", err.Error())
		return
	}

	log.Printf("已开奖期数【%d | %s】，开奖结果【%02d】，即将开奖期数【%d | %s】...\n", common.ThisIssueId, common.ThisIssueNumber, common.ThisResult, common.NextIssueId, common.NextIssueNumber)

	// 即将开奖赔率
	base.Sleep(rand.Float64() * 30)
	nextIssue, err := getInfo(common.NextIssueId)
	if err != nil {
		log.Printf("getInfo(%d) ERROR : %s", common.NextIssueId, err.Error())
		return
	}

	log.Printf("即将开奖期数【%d | %s】，波动率【%6.4f】，累计投注额【%d】...\n", common.NextIssueId, common.NextIssueNumber, nextIssue.Sqrt, nextIssue.Total)

	bets, total := make([]int, 0, len(base.SN28)), 0
	for _, no := range base.SN28 {
		sigma := nextIssue.Values[no] / (1000 / float64(base.STDS1000[no]))

		var delta float64
		if sigma < nextIssue.Avg {
			delta = math.Pow((sigma-nextIssue.Min)/(nextIssue.Avg-nextIssue.Min), 2)
		} else {
			delta = sigma * math.Pow(base.Config.Enigma, sigma-nextIssue.Avg)
		}

		//if sigma > base.Config.Sigma {
		//	if sigma <= 1.0 {
		//		delta = (sigma - base.Config.Sigma) / (1.0 - base.Config.Sigma)
		//	} else {
		//		delta = sigma * math.Pow(base.Config.Enigma, sigma-1.0)
		//	}
		//}

		//if sigma < nextIssue.Avg {
		//	delta = math.Pow((sigma-nextIssue.Min)/(nextIssue.Avg-nextIssue.Min), 2)
		//} else {
		//	delta = math.Pow((nextIssue.Max-sigma)/(nextIssue.Max-nextIssue.Avg), 2)
		//}

		bet := int(delta * float64(base.Config.Base) * float64(base.STDS1000[no]) / 1000)
		if bet <= 0 {
			bet = 0
			log.Printf("  【   】数字【%02d】，赔率【%-8.2f %6.4f】...\n", no, nextIssue.Values[no], sigma)
		} else {
			log.Printf("  【 ✓ 】数字【%02d】，赔率【%-8.2f %6.4f】，投注系数【%6.4f】...\n", no, nextIssue.Values[no], sigma, delta)
		}

		total = total + bet
		bets = append(bets, bet)
	}

	if nextIssue.Sqrt < base.Config.Sqrt {
		log.Printf("/********************************** 开奖期数【%d | %s】的波动率【%6.4f】小于设定值【%6.4f】，本期不进行投注 **********************************/\n", common.NextIssueId, common.NextIssueNumber, nextIssue.Sqrt, base.Config.Sqrt)
		return
	}

	sBets := make([]string, 0, len(bets))
	for _, bet := range bets {
		sBets = append(sBets, fmt.Sprintf("%d", bet))
	}

	// 执行投注
	//base.SleepTo(50.0 + rand.Float64()*10)
	//name, sBetEscape := fmt.Sprintf("%d%s", user.Gold/10000%1000, time.Now().Format("1504")), url.QueryEscape(strings.Join(sBets, ","))
	//if err := doMode(name, sBetEscape); err != nil {
	//	log.Printf("doBet() ERROR : %s", err.Error())
	//	return
	//}
	//log.Printf("/********************************** 保存投注模式【%s】，设置的投注金额【%-7d】，波动率【%6.4f】 **********************************/\n", name, total, nextIssue.Sqrt)

	base.Sleep(rand.Float64() * 30)
	sBetEscape := url.QueryEscape(strings.Join(sBets, ","))
	if err := doBet(common.NextIssueNumber, sBetEscape, total); err != nil {
		log.Printf("doBet() ERROR : %s", err.Error())
		return
	}

	log.Printf("/********************************** 投注已完成，投注金额【%-7d】 **********************************/\n", total)
}
