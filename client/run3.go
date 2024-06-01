package client

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type SpaceResult struct {
	Num  int32
	Rate float64
}

func run3() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Println("//*********************************** 定时任务开始执行【模式3】 ***********************************//")

	// 是否暂停
	if stop > 0 {
		log.Printf("<0> 暂停执行中，请等待【%d】期 >>>\n ", stop)

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

	if len(latest) > 0 && conf.Stop > 0 {
		if _, ok := latest[result]; !ok {
			stop = conf.Stop + rand.Intn(3)
			log.Printf("<0> 暂停执行中，请等待【%d】期 >>>\n ", stop)

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

	// 排序
	spaces := make([]SpaceResult, 0, len(rds))
	for num, rate := range rds {
		spaces = append(spaces, SpaceResult{Num: num, Rate: rate})
	}
	sort.Slice(spaces, func(i, j int) bool {
		return spaces[i].Rate >= spaces[j].Rate
	})

	// 重新添加
	nds := make(map[int32]float64)
	removed05, removed69, removed1013 := 6, 3, 2
	for _, space := range spaces {
		if removed05 > 0 && ((space.Num >= 0 && space.Num <= 5) || (space.Num >= 22 && space.Num <= 27)) {
			removed05--
			continue
		}

		if removed69 > 0 && ((space.Num >= 6 && space.Num <= 9) || (space.Num >= 18 && space.Num <= 21)) {
			removed69--
			continue
		}

		if removed1013 > 0 && ((space.Num >= 10 && space.Num <= 13) || (space.Num >= 14 && space.Num <= 17)) {
			removed1013--
			continue
		}

		nds[space.Num] = space.Rate
	}

	// 计算投注数字
	latest = make(map[int]struct{})
	bets, nums, summery := make(map[int32]int32), make([]string, 0), int32(0)
	for _, n := range SN28 {
		var rx float64

		if _, ok := nds[n]; !ok {
			log.Printf("  竞猜数字【%02d】：当前间隔/标准间隔【%.3f】，自动忽略【 ✘ 】； \n", n, rds[n])
			continue
		}

		if nds[n] < 0.50 {
			rx = 1.2
		} else if nds[n] < 0.8 {
			rx = 1.0
		} else if nds[n] < 1.2 {
			rx = 0.8
		} else if nds[n] < 1.6 {
			rx = 0.5
		} else if nds[n] < 2.0 {
			rx = 0.2
		} else {
			log.Printf("  竞猜数字【%02d】：当前间隔/标准间隔【%.3f】，投注系数【 - 】； \n", n, nds[n])
			continue
		}

		log.Printf("  竞猜数字【%02d】：当前间隔/标准间隔【%.3f】，投注系数【%.2f】； \n", n, nds[n], rx)
		iGold := int32(rx * float64(conf.Base) * float64(STDS1000[n]) / 1000)

		bets[n] = iGold
		summery = summery + iGold
		nums = append(nums, fmt.Sprintf("%02d", n))

		if rx >= 0.75 {
			latest[int(n)] = struct{}{}
		}
	}

	log.Printf("【 按最热结果 】投注基数【%d】，投注数字【%q】，投注金额【%d】  >>> \n", conf.Base, strings.Join(nums, ", "), summery)
	time.Sleep(time.Second * time.Duration(5*rand.Float64()))

	// 最后一步 执行投注数字
	if err := qBetting(fmt.Sprintf("%d", issue+1), bets); err != nil {
		log.Printf("【ERR-X9】: %s \n", err)
	}

	log.Println("<9> 全部执行结束 ...")
}
