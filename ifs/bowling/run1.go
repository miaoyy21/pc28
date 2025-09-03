package pc28

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/url"
	"pc28/base"
	"strings"
)

func run1() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Printf("/********************************** 开始执行定时任务 **********************************/")

	// 获取用户信息
	base.Sleep(rand.Float64() * 10)
	value, err := getIndex()
	if err != nil {
		log.Printf("getIndex() ERROR : %s", err.Error())
		return
	}

	log.Printf("已开奖期数【%s】，开奖结果【%s】，当前余额【%d】...\n", value.ThisIssueId, value.ThisResult, value.UserEggs)

	// 即将开奖赔率
	base.Sleep(rand.Float64() * 10)
	detail, err := getDetail(value.NextIssueId)
	if err != nil {
		log.Printf("getDetail(%s) ERROR : %s", value.NextIssueId, err.Error())
		return
	}

	bets, total := make([]int, 0, len(base.SN28)), 0
	for _, no := range base.SN28 {
		sigma := detail.Values[no] / (1000 / float64(base.STDS1000[no]))

		var delta float64
		if sigma < 1.0 {
			delta = (sigma - base.Config.Sigma) / (1.0 - base.Config.Sigma)
		} else {
			delta = sigma * math.Pow(base.Config.Enigma, sigma-1.0)
		}

		bet := int(delta * float64(base.Config.Base) * float64(base.STDS1000[no]) / 1000)
		if bet <= 0 {
			bet = 0
			log.Printf("  【   】数字【%02d】，赔率【%-8.2f %6.4f】...\n", no, detail.Values[no], sigma)
		} else {
			log.Printf("  【 ✓ 】数字【%02d】，赔率【%-8.2f %6.4f】，投注系数【%6.4f】...\n", no, detail.Values[no], sigma, delta)
		}

		total = total + bet
		bets = append(bets, bet)
	}

	if detail.Sqrt < base.Config.Sqrt {
		log.Printf("/********************************** 开奖期数【%s】的波动率【%6.4f】小于设定值【%6.4f】，本期不进行投注 **********************************/\n", value.NextIssueId, detail.Sqrt, base.Config.Sqrt)
		return
	}

	sBets := make([]string, 0, len(bets))
	for _, bet := range bets {
		sBets = append(sBets, fmt.Sprintf("%d", bet))
	}

	base.Sleep(rand.Float64() * 10)
	sBetEscape := url.QueryEscape(strings.Join(sBets, ","))
	if err := doBet(common.NextIssueNumber, sBetEscape, total); err != nil {
		log.Printf("doBet() ERROR : %s", err.Error())
		return
	}

	log.Printf("/********************************** 投注已完成，投注金额【%-7d】 **********************************/\n", total)
}
