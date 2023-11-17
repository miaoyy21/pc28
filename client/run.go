package client

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

func run(db *sql.DB, portGold, portBetting string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【Exception】: %s \n", err)
		}
	}()

	log.Println("//*********************************** 定时任务开始执行 ***********************************//")

	// 第一步 查询本账号的最新期数
	sleepTo(30.0 + 5*rand.Float64())
	log.Println("查询本账号的最新期数 >>> ")

	issue, total, err := qIssueGold()
	if err != nil {
		log.Printf("【ERR-11】: %s \n", err)
		return
	}
	log.Printf("最新开奖期数【%d】 ... \n", issue)

	mrx := 1.0
	if total < 1<<26 {
		// TODO 调整参数
		mrx = float64(total) / float64(1<<26) // 134,217,728 / 2
	}

	// 第二步 查询托管账户的金额
	sleepTo(40.0 + 5*rand.Float64())
	log.Println("查询托管账户的资金余额 >>> ")

	users, err := dQueryUsers(db)
	if err != nil {
		log.Printf("【ERR-21】: %s \n", err)
		return
	}

	for _, user := range users {
		gold, err := gGold(net.JoinHostPort(user.Host, portGold), user.Cookie, user.UserAgent, user.Unix, user.KeyCode, user.DeviceId, user.UserId, user.Token)
		if err != nil {
			log.Printf("【ERR-22】: [%s] %s \n", user.UserId, err)
			return
		}

		user.Gold = gold

		// Update User's Gold
		if _, err := db.Exec("UPDATE user SET gold = ?, update_at = ? WHERE user_id = ?", gold, time.Now().Format("2006-01-02 15:04"), user.UserId); err != nil {
			log.Printf("【ERR-23】: [%s] %s \n", user.UserId, err)
			return
		}

		// Insert User's Log
		if _, err := db.Exec("INSERT INTO user_log(user_id, time_at, gold) VALUES (?,?,?)", user.UserId, time.Now().Format("2006-01-02 15:04"), gold); err != nil {
			log.Printf("【ERR-24】: [%s] %s \n", user.UserId, err)
			return
		}

		log.Printf("托管账户【%-10s】的资金余额 %d ... \n", user.UserName, user.Gold)
	}

	// 第三步 查询本账户的权重值
	sleepTo(54.0)
	log.Println("查询本账户的权重值 >>> ")

	rds, err := qRiddle(fmt.Sprintf("%d", issue+1))
	if err != nil {
		log.Printf("【ERR-31】: %s \n", err)
		return
	}

	// 第四步 委托账户投注
	var wg sync.WaitGroup

	wg.Add(len(users))
	log.Println("执行托管账户投注 >>> ")

	for _, user := range users {
		go func(user *User) {
			m1Gold := ofM1Gold(user.Gold)
			log.Printf("托管账户【%-10s】 ：活跃系数【%.4f】，原投注基数【%d】，实际投注基数【%d】 >>> \n", user.UserName, mrx, m1Gold, int64(mrx*float64(m1Gold)))

			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

			bets := make(map[int32]int32)
			for _, n := range SN28 {
				if rds[n] <= user.Sigma {
					continue
				}

				var sig float64
				if rds[n] > 1.0 {
					sig = rds[n]
				} else {
					sig = (rds[n] - user.Sigma) / (1.0 - user.Sigma)
				}

				fGold := mrx * sig * float64(m1Gold) * float64(STDS1000[n]) / 1000

				// 转换可投注额
				iGold := int32(fGold)
				if int64(mrx*float64(m1Gold)) > 1<<19 {
					iGold = ofGold(fGold) // 524,288
				}

				if iGold > 0 {
					bets[n] = iGold
				}
			}

			if err := gBetting(net.JoinHostPort(user.Host, portBetting), fmt.Sprintf("%d", issue+1), bets,
				user.Cookie, user.UserAgent, user.Unix, user.KeyCode, user.DeviceId, user.UserId, user.Token); err != nil {
				log.Printf("【ERR-41】: %s \n", err)

				if _, err := db.Exec("UPDATE user SET msg = ? WHERE user_id = ?", err.Error(), user.UserId); err != nil {
					log.Printf("【ERR-42】: [%s] %s \n", user.UserId, err)
					return
				}

				return
			}

			if _, err := db.Exec("UPDATE user SET msg = ? WHERE user_id = ?", "OK", user.UserId); err != nil {
				log.Printf("【ERR-43】: [%s] %s \n", user.UserId, err)
				return
			}

			wg.Done()
		}(user)
	}

	wg.Wait()
	log.Println("全部执行结束 ...")
}
