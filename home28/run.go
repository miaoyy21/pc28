package home28

import (
	"fmt"
	"log"
	"math"
	"strings"
	"tty28/base"
)

var skips = 0
var lastBets = make([]int, 0, len(base.SN28))

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

	log.Printf("用户ID【%d】，用户名【%s】，当前余额【%d】...\n", user.Id, user.Name, user.Gold)

	// 获取最新的已开奖及即将开奖信息
	common, err := getCommon()
	if err != nil {
		log.Printf("ERROR : %s", err.Error())
		return
	}

	log.Printf("已开奖期数【%d】，开奖结果【%02d】，即将开奖期数【%d】...\n", common.LastId, common.LastResult, common.NextId)

	// 已开奖结果赔率
	exactly := false
	if len(lastBets) > 0 {
		lastInfo, err := getInfo(common.LastId)
		if err != nil {
			log.Printf("ERROR : %s", err.Error())
			return
		}

		odd := lastInfo.Values[common.LastResult]
		delta := odd / lastInfo.Values[27-common.LastResult]
		if delta < 1 && lastBets[common.LastResult] > 0 {
			exactly = true
		}
	}

	// 即将开奖赔率
	NextInfo, err := getInfo(common.NextId)
	if err != nil {
		log.Printf("ERROR : %s", err.Error())
		return
	}

	log.Printf("即将开奖期数【%d】，累计投注额【%s】...\n", common.NextId, NextInfo.TotalBet)

	bets, total := make([]int, 0, len(base.SN28)), 0
	for _, no := range base.SN28 {
		odd := NextInfo.Values[no]
		delta := odd / NextInfo.Values[27-no]

		if delta >= 1 {
			bets = append(bets, 0)
			log.Printf("  【   】数字【%02d】，赔率【%-8.2f】，赔率系数【%4.2f】...\n", no, odd, delta)
		} else {
			bet := base.Config.Base * base.STDS1000[no] / 1000
			total = total + bet

			bets = append(bets, bet)
			log.Printf("  【 ✓ 】数字【%02d】，赔率【%-8.2f】，赔率系数【%4.2f】...\n", no, odd, delta)
		}
	}

	if total < 10000 {
		delta := 10000 / float64(total)

		total = 0
		newBets := make([]int, 0, len(bets))
		for _, bet := range bets {
			newBet := int(math.Ceil(float64(bet) * delta))
			total = total + newBet
			newBets = append(newBets, newBet)
		}

		bets = append(bets[0:0], newBets...)
		log.Printf("《投注金额小于最小值，提高倍率【%4.2f】已达到投注要求》...\n", delta)
	}

	sBets := make([]string, 0, len(bets))
	for _, bet := range bets {
		sBets = append(sBets, fmt.Sprintf("%d", bet))
	}
	if len(sBets) != 28 {
		log.Panicf("<<< unreachable >>>")
	}

	// 判定是否执行
	if len(lastBets) < 1 {
		log.Printf("/********************************** <<< 首次运行，不投注 >>> **********************************/\n")
	} else if !exactly {
		newBets, newLatestBets := make([]string, 0, len(sBets)), make([]int, 0, len(sBets))
		for no := range base.SN28 {
			newBets = append(newBets, sBets[27-no])
			newLatestBets = append(newLatestBets, bets[27-no])
		}

		if err := doBet(common.NextId, strings.Join(newBets, ",")); err != nil {
			log.Printf("ERROR : %s", err.Error())
			return
		}

		log.Printf("/********************************** 执行反向投注，投注金额【%-7d】 **********************************/\n", total)
	} else {
		if err := doBet(common.NextId, strings.Join(sBets, ",")); err != nil {
			log.Printf("ERROR : %s", err.Error())
			return
		}

		log.Printf("/********************************** 投注已完成，投注金额【%-7d】 **********************************/\n", total)
	}
	lastBets = bets
}
