package main

import (
	"flag"
	"log"
	"pc28/client"
	"pc28/server"
)

var sFlag bool
var cFlag bool

func init() {
	flag.BoolVar(&sFlag, "s", false, "服务端模式")
	flag.BoolVar(&cFlag, "c", false, "客户端模式")

	flag.Parse()
}

func main() {
	if sFlag == cFlag {
		flag.Usage()
		return
	}

	if sFlag {
		log.Printf("以【服务端】模式运行 ...\n")
		if err := server.Run(); err != nil {
			log.Fatalf("server.Run() Failure : %s \n", err)
		}
		return
	}

	log.Printf("以【客户端】模式运行 ...\n")
	if err := client.Run(); err != nil {
		log.Fatalf("client.Run() Failure : %s \n", err)
	}
}
