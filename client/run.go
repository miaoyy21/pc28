package client

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

	resp, err := Search()
	if err != nil {
		return err
	}

	fmt.Println(resp.GetResponse())
	return nil
}
