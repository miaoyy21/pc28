package main

import (
	"log"
	"pc28/ifs"
)

func main() {
	if err := ifs.Run(); err != nil {
		log.Fatalf("服务启动失败：%s \n", err.Error())
	}
}
