package client

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

var latest map[int]struct{}
var stop int

func run2() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Println("//*********************************** 定时任务开始执行 ***********************************//")
	
	// 是否暂停
	if stop > 0 {
		log.Printf("<0> 暂停执行中，请等待【%02d】期 >>>\n ", stop)

		stop--
		return
	}

	sleepTo(30.0 + 5*rand.Float64())

	// 第一步 查询本账号的最新期数
	log.Println("<1> 查询本账号的最新期数 >>> ")

	issue, total, result, err := qIssueGold()
	if err != nil {
		log.Printf("【ERR-X1】: %s \n", err)
		return
	}
	log.Printf("  最新开奖期数【%d】，资金池【%d】，开奖结果【%02d】 ... \n", issue, total, result)
	time.Sleep(time.Second * time.Duration(5*rand.Float64()))

	if len(latest) > 0 {
		if _, ok := latest[result]; !ok {
			stop = 10
			log.Printf("<0> 暂停执行中，请等待【%02d】期 >>>\n ", stop)

			stop--
			latest = make(map[int]struct{})
			return
		}
	}

	// 第二步 查询开奖结果间隔
	log.Println("<2> 查询开奖结果间隔 >>> ")

	rds, err := qSpace()
	if err != nil {
		log.Printf("【ERR-X2】: %s \n", err)
		return
	}

	// 计算投注数字
	latest = make(map[int]struct{})
	m1Gold := 200000
	bets, nums, summery := make(map[int32]int32), make([]string, 0), int32(0)
	for _, n := range SN28 {
		var rx float64

		if rds[n] < 0.50 {
			rx = 1.2
		} else if rds[n] < 0.75 {
			rx = 1.0
		} else if rds[n] < 1.0 {
			rx = 0.8
		} else if rds[n] < 1.25 {
			rx = 0.5
		} else if rds[n] < 1.50 {
			rx = 0.2
		} else {
			log.Printf("  竞猜数字【%02d】：当前间隔/标准间隔【%.3f】，投注系数【 - 】； \n", n, rds[n])
			continue
		}

		log.Printf("  竞猜数字【%02d】：当前间隔/标准间隔【%.3f】，投注系数【%.2f】； \n", n, rds[n], rx)
		iGold := int32(rx * float64(m1Gold) * float64(STDS1000[n]) / 1000)

		bets[n] = iGold
		summery = summery + iGold
		nums = append(nums, fmt.Sprintf("%02d", n))

		if rx >= 0.75 {
			latest[int(n)] = struct{}{}
		}
	}

	log.Printf("【 按最热结果 】所选的投注数字【%q】，总金额【%d】  >>> \n", strings.Join(nums, ", "), summery)
	time.Sleep(time.Second * time.Duration(5*rand.Float64()))

	// 最后一步 执行投注数字
	if err := qBetting(fmt.Sprintf("%d", issue+1), bets); err != nil {
		log.Printf("【ERR-X9】: %s \n", err)
	}

	log.Println("<9> 全部执行结束 ...")
}
