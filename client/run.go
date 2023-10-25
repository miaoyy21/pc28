package client

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"time"
)

func Run(targetGold, targetBetting string) error {
	// 配置文件
	if err := InitConfig(); err != nil {
		return err
	}

	// 数据库连接对象
	db, err := sql.Open("mysql", conf.Dsn)
	if err != nil {
		return err
	}

	// Ping
	if err := db.Ping(); err != nil {
		return err
	}
	log.Println("连接数据库成功 ...")

	// 如果超过30秒，那么等待30秒后运行
	if time.Now().Second() >= 30 {
		log.Printf("当前时间为%q，等待30秒 ... \n", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(30 * time.Second)
	}

	sleepTo(30)
	go run(db, targetGold, targetBetting)

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

			go run(db, targetGold, targetBetting)
		}
	}
}

func run(db *sql.DB, targetGold, targetBetting string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Exception : %s \n", err)
		}
	}()

	// 第一步 查询本账号的最新期数
	sleepTo(30.0 + 5*rand.Float64())
	log.Println("【1】执行查询本账号的最新期数 ... ")
	issue, total, err := qIssueGold()
	if err != nil {
		log.Printf("ERR-11 : %s \n", err)
		return
	}
	log.Printf("【1】本账号的最新期数为 %d ... \n", issue)

	mrx := 1.0
	if total < 1<<27 {
		mrx = float64(total) / float64(1<<27) // 134,217,728
	}

	// 第二步 查询托管账户的金额
	sleepTo(40.0 + 5*rand.Float64())
	log.Println("【2】执行查询托管账户的金额 ... ")

	users, err := dQueryUsers(db)
	if err != nil {
		log.Printf("ERR-21 : %s \n", err)
		return
	}

	for _, user := range users {
		gold, err := gGold(targetGold, user.Cookie, user.UserAgent, user.Unix, user.KeyCode, user.DeviceId, user.UserId, user.Token)
		if err != nil {
			log.Printf("ERR-22 : [%s] %s \n", user.UserId, err)
			return
		}

		user.Gold = gold
		if _, err := db.Exec("UPDATE users SET gold = ? WHERE user_id = ?", gold, user.UserId); err != nil {
			log.Printf("ERR-23 : [%s] %s \n", user.UserId, err)
			return
		}
	}
	log.Printf("【2】TODO 查询托管账户的金额 %#v ... \n", users)

	// 第三步 查询本账户下期权重值
	sleepTo(54.0)
	log.Println("【3】执行查询本账户下期权重值 ... ")
	rds, err := qRiddle(fmt.Sprintf("%d", issue+1))
	if err != nil {
		log.Printf("ERR-31 : %s \n", err)
		return
	}
	log.Printf("【3】TODO 查询本账户下期权重值 %#v ... \n", rds)

	// 第四步 委托账户投注
	log.Println("【4】执行委托账户投注 ... ")
	for _, user := range users {
		m1Gold := ofM1Gold(user.Gold)

		bets := make(map[int32]int32)
		for _, i := range SN28 {
			if rds[i] <= user.Sigma {
				continue
			}

			fGold := mrx * ((rds[i] - user.Sigma) / (1.0 - user.Sigma)) * float64(2*m1Gold) * float64(STDS1000[i]) / 1000
			iGold := ofGold(fGold)
			if iGold > 0 {
				bets[i] = iGold
			}
		}

		if err := gBetting(targetBetting, fmt.Sprintf("%d", issue+1), bets,
			user.Cookie, user.UserAgent, user.Unix, user.KeyCode, user.DeviceId, user.UserId, user.Token); err != nil {
			log.Printf("ERR-41 : %s \n", err)
			return
		}
	}
	log.Println("【4】 执行委托账户投注完成 ... ")
}

func sleepTo(s0 float64) {
	d0 := time.Now().Sub(time.Now().Truncate(time.Minute))
	if s0-d0.Seconds() < 0 {
		panic(fmt.Sprintf("目标第%.2f秒小于当前第%.2f秒", s0, d0.Seconds()))
	}

	log.Printf("等待%.2f秒后继续执行 ... \n", s0-d0.Seconds())
	time.Sleep(time.Second * time.Duration(s0-d0.Seconds()))
}
