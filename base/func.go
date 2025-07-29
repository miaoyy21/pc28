package base

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func SleepTo(s0 float64) {
	d0 := time.Now().Sub(time.Now().Truncate(time.Minute))
	if s0-d0.Seconds() < 0 {
		log.Printf(fmt.Sprintf("【网络延迟原因？？？？？？】目标第%.2f秒小于当前第%.2f秒\n", s0, d0.Seconds()))
	}

	log.Printf("等待%.2f秒后继续执行 ... \n", s0-d0.Seconds())
	time.Sleep(time.Second * time.Duration(s0-d0.Seconds()))
}

func OfNextIssue(issue int64) string {
	sDt := time.Now().Format("20060102")
	sIssue := fmt.Sprintf("%d", issue+1)

	if strings.EqualFold(sDt, sIssue[:8]) {
		return sIssue
	}

	return fmt.Sprintf("%s0001", sDt)
}
