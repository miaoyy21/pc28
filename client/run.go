package client

import (
	"database/sql"
	"fmt"
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

	sleepTo(30)
	go run()

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

			go run()
		}
	}

	//resp, err := Search()
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Println(resp.GetResponse())
}

func run() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("【存在异常】: %s \n", err)
		}
	}()

	sleepTo(30.0 + 10*rand.Float64())
	issue, err := qIssue()
	if err != nil {
		log.Printf("ERR-01 : %s \n", err)
		return
	}

	log.Println("【TODO】 执行查询所有托管账户金额 >>> ")

	sleepTo(54.0)
	log.Println("【TODO】 执行查询本账户权重，并执行所有托管账户投注 >>> ")

	time.Sleep(time.Duration(20*rand.Float64()) * time.Second)
}

func sleepTo(s0 float64) {
	d0 := time.Now().Sub(time.Now().Truncate(time.Minute))
	if s0-d0.Seconds() < 0 {
		panic(fmt.Sprintf("目标第%.2f秒小于当前第%.2f秒", s0, d0.Seconds()))
	}

	log.Printf("等待%.2f秒后继续执行 ... \n", s0-d0.Seconds())
	time.Sleep(time.Second * time.Duration(s0-d0.Seconds()))
}
