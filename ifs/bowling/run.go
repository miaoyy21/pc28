package bowling

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

	log.Println("启动定时器完成 ...")

	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			s3 := time.Now()
			if s3.Second() < 30 || s3.Second() > 33 {
				continue
			}

			// 重新加载配置
			if err := base.InitConfig(); err != nil {
				log.Printf("重新加载配置文件错误：%s \n", err.Error())
				continue
			}
			log.Printf("重载配置文件成功 ...\n")

			// 执行投注
			go run1()

			// 暂停等待5秒
			time.Sleep(45 * time.Second)
		}
	}
}
