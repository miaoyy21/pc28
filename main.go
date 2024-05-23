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
	flag.BoolVar(&sFlag, "s", true, "服务端模式")
	flag.BoolVar(&cFlag, "c", false, "客户端模式")

	flag.Parse()
}

const portGold = "9001"
const portBetting = "9002"

func main() {
	if sFlag == cFlag {
		flag.Usage()
		return
	}

	if sFlag {
		log.Printf("以【服务端】模式运行 ...\n")
		server.Run(portGold, portBetting)

		return
	}

	log.Printf("以【客户端】模式运行 ...\n")
	if err := client.Run(portGold, portBetting); err != nil {
		log.Fatalf("client.Run() Failure : %s \n", err)
	}
}
