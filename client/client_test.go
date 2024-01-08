package client

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	for i := 0; i < 10; i++ {
		go do()
	}

	fmt.Println("Done")

	stds := make(map[int]int)
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				stds[i+j+k]++
			}
		}
	}

	fmt.Printf("%#v \n", stds)
}

func do() {
	fmt.Println("Hello")
}
