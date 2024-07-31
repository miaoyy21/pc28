package client

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math"
	"math/rand"
	"strings"
)

func run1Local() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Println("//*********************************** 定时任务开始执行 模式1 本地 ***********************************//")

	// 第一步 查询本账号的最新期数
	sleepTo(30 + 5*rand.Float64())
	log.Println("<1> 查询本账号的最新期数 >>> ")

	issue, total, result, err := qHistory()
	if err != nil {
		log.Printf("【ERR-11】: %s \n", err)
		return
	}
	log.Printf("  最新开奖期数【%d】，资金池【%d】，开奖结果【%02d】 ... \n", issue, total, result)

	// 第二步 查询当前期的投注金额和人数
	sleepTo(48 + 1*rand.Float64())
	log.Println("<2> 查询本期的投注信息 >>> ")
	issueTotal, issueMembers, err := qIssue(fmt.Sprintf("%d", issue+1))
	if err != nil {
		log.Printf("【ERR-21】: %s \n", err)
		return
	}
	log.Printf("  本期开奖期数【%s】，总投注额【%d】，投注人数【%d】 ... \n", fmt.Sprintf("%d", issue+1), issueTotal, issueMembers)

	if issueTotal < 100000 {
		log.Printf("//********************  本期总投注额没有打到设定值【%d】，不进行投注  ********************// ... \n", 100000)
		return
	}

	// 第三步 查询本账户的权重值
	log.Println("<3> 查询本账户的权重值 >>> ")

	rds, exp, dev, err := qRiddle(fmt.Sprintf("%d", issue+1))
	if err != nil {
		log.Printf("【ERR-31】: %s \n", err)
		return
	}

	//if dev < 0.025 {
	//	log.Printf("//********************  赔率系数的标准方差没有达到设定值【%.3f】，不进行投注  ********************// ... \n", 0.025) // 16,777,216
	//	return
	//}

	rx, rn := 1.0, 0
	for _, rd := range rds {
		if rd > 1 {
			rn++
		}
	}
	if rn >= 25 {
		rx = 2.5
	}

	// 第四步 委托账户投注
	log.Println("<4> 执行托管账户投注 >>> ")

	maxSig := float64(issueTotal) * 0.3333 / float64(conf.Base)
	sigma, bets, nums, summery := 0.975, make(map[int32]int32), make([]string, 0), int32(0)
	for _, n := range SN28 {
		rd := rds[n]
		if rd <= sigma {
			continue
		}

		var sig float64
		if rd > 1.0 {
			sig = rd
			if sig > maxSig {
				sig = math.Min(dev, maxSig) * math.Pow(1.25, sig/exp)
			} else {
				if dev > 4 && sig > 4 {
					sig = math.Min(dev, sig)
				}
			}

			for sig > 50 {
				sig = 0.5 * sig * math.Pow(1.125, maxSig/exp)
			}
		} else {
			sig = (rd - sigma) / (1.0 - sigma)
		}

		fGold := rx * sig * float64(conf.Base) * float64(STDS1000[n]) / 1000

		// 转换可投注额
		iGold := ofGold(fGold)

		if iGold > 0 {
			bets[n] = iGold
			summery = summery + iGold
			nums = append(nums, fmt.Sprintf("🍌  %02d 【%-7d  %-5.2f  %8d】", n, iGold, sig, int(math.Ceil(float64(iGold)*sig*1000.0/float64(STDS1000[n])))))
		}
	}

	log.Printf("  最大系数【%.3f】，扩大倍率【%.3f】，投注基数【%d】，投注金额【🍎 %d】，投注数字：\n\t  %s  \n  >>> \n", maxSig, rx, conf.Base, summery, strings.Join(nums, "\n\t  "))

	// 最后一步 执行投注数字
	if err := qBetting(fmt.Sprintf("%d", issue+1), bets); err != nil {
		log.Printf("【ERR-X9】: %s \n", err)
	}

	log.Println("<9> 全部执行结束 ...")
}
