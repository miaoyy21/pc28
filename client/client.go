package client

import (
	"fmt"
	"log"
	"math"
	"time"
)

var SN28 = []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}

var STDS1000 = map[int32]int32{
	0:  1,
	1:  3,
	2:  6,
	3:  10,
	4:  15,
	5:  21,
	6:  28,
	7:  36,
	8:  45,
	9:  55,
	10: 63,
	11: 69,
	12: 73,
	13: 75,
	14: 75,
	15: 73,
	16: 69,
	17: 63,
	18: 55,
	19: 45,
	20: 36,
	21: 28,
	22: 21,
	23: 15,
	24: 10,
	25: 6,
	26: 3,
	27: 1,
}

func ofM1Gold(user *User) int64 {
	if user.M1Gold > 0 {
		return int64(user.M1Gold)
	}

	g := user.Gold
	if g < 1<<22 {
		// 4194304
		return g / 40
	} else if g < 1<<23 {
		// 8388608
		return g / 50
	} else if g < 1<<24 {
		// 16777216
		return g / 60
	} else if g < 1<<25 {
		// 33554432
		return g / 75
	} else if g < 1<<26 {
		// 67108864
		return g / 95
	} else if g < 1<<27 {
		// 134217728
		return g / 120
	} else if g < 1<<28 {
		// 268435456
		return g / 150
	} else {
		return g / 200
	}
}

func ofGold(fGold float64) int32 {
	var iGold int32
	if fGold >= 1<<16 {
		iGold = int32(math.Round(fGold/2000.0) * 2000)
	} else if fGold >= 1<<15 {
		iGold = int32(math.Round(fGold/1500.0) * 1500)
	} else if fGold >= 1<<14 {
		iGold = int32(math.Round(fGold/1000.0) * 1000)
	} else {
		iGold = int32(math.Round(fGold/500.0) * 500)
	}

	return iGold
}

func sleepTo(s0 float64) {
	d0 := time.Now().Sub(time.Now().Truncate(time.Minute))
	if s0-d0.Seconds() < 0 {
		log.Printf(fmt.Sprintf("【网络延迟原因？？？？？？】目标第%.2f秒小于当前第%.2f秒\n", s0, d0.Seconds()))
	}

	log.Printf("等待%.2f秒后继续执行 ... \n", s0-d0.Seconds())
	time.Sleep(time.Second * time.Duration(s0-d0.Seconds()))
}

func Run(portGold, portBetting string) error {
	// 配置文件
	if err := InitConfig(); err != nil {
		return err
	}

	// 数据库连接对象
	//db, err := sql.Open("mysql", conf.Dsn)
	//if err != nil {
	//	return err
	//}

	// Ping
	//if err := db.Ping(); err != nil {
	//	return err
	//}
	//log.Println("连接数据库成功 ...")

	// 如果超过30秒，那么等待30秒后运行
	if time.Now().Second() >= 30 {
		log.Printf("当前时间为%q，等待30秒 ... \n", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(30 * time.Second)
	}

	sleepTo(30)
	//go run1(db, portGold, portBetting)
	go run1Local()

	t := time.NewTicker(time.Minute)
	defer t.Stop()

	log.Println("启动定时器完成 ...")
	for {
		select {
		case <-t.C:
			// 长时间运行时，可能会产生时间偏移，自动调整
			d0 := time.Now().Sub(time.Now().Truncate(time.Minute))
			t.Reset(time.Duration(90-d0.Seconds()) * time.Second)
			log.Printf("【重置时钟】偏移量%.2f秒 ...\n", 30-d0.Seconds())

			// 重新加载配置
			if err := InitConfig(); err != nil {
				log.Printf("重新加载配置文件错误：%s \n", err.Error())
				continue
			}
			log.Printf("重载配置文件成功 ...\n")

			//go run1(db, portGold, portBetting)
			go run1Local()
		}
	}
}
