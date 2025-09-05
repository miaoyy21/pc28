package bowling

import (
	"log"
	"math"
	"math/rand"
	"pc28/base"
)

var preBets map[int]int
var totalBets int
var winBets int

func run1() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Printf("/********************************** 开始执行定时任务 **********************************/")

	// 获取用户信息
	base.Sleep(rand.Float64() * 5)
	value, err := getIndex()
	if err != nil {
		log.Printf("getIndex() ERROR : %s", err.Error())
		return
	}

	var rateBets float64
	if len(preBets) > 0 {
		vBets, ok := preBets[value.ThisResult]
		if ok && vBets > 0 {
			winBets++
		}
	}

	// 胜率
	if totalBets > 0 {
		rateBets = float64(winBets) / float64(totalBets)
	}

	log.Printf("已开奖期数【%s】，开奖结果【%02d】，当前余额【%d】，胜率【%d/%d  %4.2f%%】...\n", value.ThisIssueId, value.ThisResult, value.UserEggs, winBets, totalBets, rateBets)

	// 即将开奖赔率
	base.Sleep(rand.Float64() * 7.5)
	detail, err := getDetail(value.NextIssueId)
	if err != nil {
		log.Printf("getDetail(%s) ERROR : %s", value.NextIssueId, err.Error())
		return
	}

	bets, total := make(map[int]int), 0
	for _, no := range base.SN28 {
		sigma := detail.Values[no] / (1000 / float64(base.STDS1000[no]))

		var delta float64
		if sigma < 1.0 {
			delta = 0
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
		bets[no] = bet
	}

	if detail.Sqrt < base.Config.Sqrt {
		preBets = make(map[int]int)
		log.Printf("/********************************** 开奖期数【%s】的波动率【%6.4f】小于设定值【%6.4f】，本期不进行投注 **********************************/\n", value.NextIssueId, detail.Sqrt, base.Config.Sqrt)
		return
	}

	base.Sleep(rand.Float64() * 10)
	if err := doSave(value.NextIssueId, bets); err != nil {
		log.Printf("doSave() ERROR : %s", err.Error())
		return
	}

	preBets = bets
	totalBets++
	log.Printf("/********************************** 投注已完成，投注金额【%-7d】 **********************************/\n", total)
}
