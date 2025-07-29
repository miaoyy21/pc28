package main

import (
	"log"
	"pc28/ifs"
)

func main() {
	if err := ifs.Run(); err != nil {
		log.Fatalf("ifs.Run() Failure %s \n", err.Error())
	}
}
