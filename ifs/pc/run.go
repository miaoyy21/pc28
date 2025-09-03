package pc

import (
	"log"
	"math/rand"
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
			if s3.Second() >= 3 {
				continue
			} else if s3.Minute()%5 == 1 {
				base.Sleep(rand.Float64() * 30)
				user, err := getUser()
				if err != nil {
					log.Printf("getUser() ERROR : %s", err.Error())
					continue
				}

				log.Printf("用户ID【%d】，手机号码【%s】，当前余额【%d】...\n", user.Id, user.Mobile, user.Gold)
				continue
			} else if s3.Minute()%5 == 0 || s3.Minute()%5 == 2 || s3.Minute()%5 == 4 {
				continue
			}

			// 重新加载配置
			if err := base.InitConfig(); err != nil {
				log.Printf("重新加载配置文件错误：%s \n", err.Error())
				continue
			}
			log.Printf("重载配置文件成功 ...\n")

			// 执行投注
			go run2()

			// 暂停等待5秒
			time.Sleep(60 * time.Second)
		}
	}
}
