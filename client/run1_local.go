package client

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	sleepTo(30.0 + 5*rand.Float64())
	log.Println("<1> 查询本账号的最新期数 >>> ")

	issue, total, result, err := qIssueGold()
	if err != nil {
		log.Printf("【ERR-11】: %s \n", err)
		return
	}

	log.Printf("  最新开奖期数【%d】，资金池【%d】，开奖结果【%02d】 ... \n", issue, total, result)

	// 第三步 查询本账户的权重值
	sleepTo(51.0)
	log.Println("<3> 查询本账户的权重值 >>> ")

	rds, dev, err := qRiddle(fmt.Sprintf("%d", issue+1))
	if err != nil {
		log.Printf("【ERR-31】: %s \n", err)
		return
	}

	//_ = dev
	if dev < 0.025 {
		log.Printf("//********************  赔率系数的标准方差没有达到设定值【%.3f】，不进行投注  ********************// ... \n", 0.025) // 16,777,216
		return
	}

	// 第四步 委托账户投注
	log.Println("<4> 执行托管账户投注 >>> ")

	m1Gold := conf.Base
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	sigma, bets, nums := 0.975, make(map[int32]int32), make([]string, 0)
	for _, n := range SN28 {
		rd := rds[n]
		if rd <= sigma {
			continue
		}

		var sig float64
		if rd > 1.0 {
			sig = rd
		} else {
			sig = (rd - sigma) / (1.0 - sigma)
		}

		fGold := sig * float64(m1Gold) * float64(STDS1000[n]) / 1000

		// 转换可投注额
		iGold := int32(fGold)
		if int64(m1Gold) > 1<<19 {
			iGold = ofGold(fGold) // 524,288
		}

		if iGold > 0 {
			bets[n] = iGold
			nums = append(nums, fmt.Sprintf("%02d", n))
		}
	}

	log.Printf("  所选的投注数字 %q  >>> \n", strings.Join(nums, ", "))

	// 最后一步 执行投注数字
	if err := qBetting(fmt.Sprintf("%d", issue+1), bets); err != nil {
		log.Printf("【ERR-X9】: %s \n", err)
	}

	log.Println("<9> 全部执行结束 ...")
}
