package main

import "tty28/home28"

func main() {
	// 加载配置模版
	//if err := conf.LoadTemplates(); err != nil {
	//	log.Panicf("config.LoadTemplates() Failure :: %s", err.Error())
	//}

	// 幸运28
	//if err := _luck.Run(); err != nil {
	//	log.Printf("_luck.Run() Failure :: %s \n", err.Error())
	//}

	// 加拿大28
	//if err := _canada.Run(); err != nil {
	//	log.Printf("_canada.Run() Failure :: %s \n", err.Error())
	//}

	home28.Run()
}
