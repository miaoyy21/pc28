package ifs

import (
	"log"
	"pc28/base"
	"time"
)

func Run() error {
	// 配置文件
	if err := base.InitConfig(); err != nil {
		return err
	}

	t := time.NewTicker(time.Second)
	defer t.Stop()

	base.SleepTo(float64(time.Now().Add(time.Second).Second()))
	log.Println("启动定时器完成 ...")
	for {
		select {
		case <-t.C:
			if time.Now().Second() != 23 && time.Now().Second() != 53 {
				if time.Now().Second() == 0 {
					base.SleepTo(1)
				}

				continue
			}

			// 重新加载配置
			if err := base.InitConfig(); err != nil {
				log.Printf("重新加载配置文件错误：%s \n", err.Error())
				continue
			}
			log.Printf("重载配置文件成功 ...\n")

			go run()
		}
	}
}
