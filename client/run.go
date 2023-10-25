package client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"time"
)

func Run() error {
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

	s0, n0 := 30.0, time.Now()
	d0 := n0.Sub(n0.Truncate(time.Minute))

	log.Printf("在 %.4f 秒后开始运行 ... \n", s0-d0.Seconds())
	time.Sleep(time.Second * time.Duration(s0-d0.Seconds()))
	go run()

	t := time.NewTicker(time.Minute)
	defer t.Stop()

	log.Println("启动定时器完成 ...")
	for {
		select {
		case <-t.C:
			go run()
		}
	}

	//resp, err := Search()
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Println(resp.GetResponse())
	return nil
}

func run() {
	time.Sleep(time.Duration(rand.Float64()*5) * time.Second)
	log.Println("【TODO】 执行查询所有托管账户金额 >>> ")
}
