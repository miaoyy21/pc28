package client

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"
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
	sleepTo(48)
	log.Println("<2> 查询本期的投注信息 >>> ")
	issueTotal, issueMembers, err := qIssue(fmt.Sprintf("%d", issue+1))
	if err != nil {
		log.Printf("【ERR-21】: %s \n", err)
		return
	}
	log.Printf("  本期开奖期数【%s】，总投注额【%d】，投注人数【%d】 ... \n", fmt.Sprintf("%d", issue+1), issueTotal, issueMembers)

	if issueTotal < 1<<18 {
		log.Printf("//********************  本期总投注额没有打到设定值【%d】，不进行投注  ********************// ... \n", 1<<18) // 16,777,216
		return
	}

	// 第三步 查询本账户的权重值
	sleepTo(50)
	log.Println("<3> 查询本账户的权重值 >>> ")

	rds, _, dev, err := qRiddle(fmt.Sprintf("%d", issue+1))
	if err != nil {
		log.Printf("【ERR-31】: %s \n", err)
		return
	}

	if dev < 0.025 {
		log.Printf("//********************  赔率系数的标准方差没有达到设定值【%.3f】，不进行投注  ********************// ... \n", 0.025) // 16,777,216
		return
	}

	// 第四步 委托账户投注
	log.Println("<4> 执行托管账户投注 >>> ")
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	maxSig := float64(issueTotal) * 0.25 / float64(conf.Base)
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
				sig = math.Min(dev, maxSig*math.Pow(1.01, (sig-maxSig)/4.0))
			} else {
				sig = math.Min(dev, sig)
			}
		} else {
			sig = (rd - sigma) / (1.0 - sigma)
		}

		fGold := sig * float64(conf.Base) * float64(STDS1000[n]) / 1000

		// 转换可投注额
		iGold := ofGold(fGold)

		if iGold > 0 {
			bets[n] = iGold
			summery = summery + iGold
			nums = append(nums, fmt.Sprintf("%02d [%.2f]", n, sig))
		}
	}

	log.Printf("  最大系数【%.3f】，投注基数【%d】，投注数字 %q，投注金额【%d】  >>> \n", maxSig, conf.Base, strings.Join(nums, ", "), summery)

	// 最后一步 执行投注数字
	if err := qBetting(fmt.Sprintf("%d", issue+1), bets); err != nil {
		log.Printf("【ERR-X9】: %s \n", err)
	}

	log.Println("<9> 全部执行结束 ...")
}
