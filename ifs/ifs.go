package ifs

import (
	"pc28/base"
)

func Run() error {
	// 配置文件
	if err := base.InitConfig(); err != nil {
		return err
	}

	run()
	return nil

	//log.Println("启动定时器完成 ...")
	//
	//t := time.NewTicker(time.Second)
	//defer t.Stop()
	//
	//for {
	//	select {
	//	case <-t.C:
	//		s3 := time.Now().Add(3 * time.Second)
	//		if s3.Second() > 3 || s3.Minute()%5 != 0 {
	//			continue
	//		}
	//
	//		// 重新加载配置
	//		if err := base.InitConfig(); err != nil {
	//			log.Printf("重新加载配置文件错误：%s \n", err.Error())
	//			continue
	//		}
	//		log.Printf("重载配置文件成功 ...\n")
	//
	//		// 执行投注
	//		run()
	//
	//		// 暂停等待5秒
	//		time.Sleep(5 * time.Second)
	//	}
	//}
}
