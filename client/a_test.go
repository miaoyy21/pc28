package client

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	tc := time.NewTicker(time.Second * 10)

	t.Logf("%s Start ... \n", time.Now().Format("2006-01-02 15:04:05"))
	for {
		select {
		case <-tc.C:
			t.Logf("%s Reset ... \n", time.Now().Format("2006-01-02 15:04:05"))
			tc.Reset(2 * time.Second)
		}
	}
}
