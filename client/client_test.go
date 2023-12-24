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
}

func do() {
	fmt.Println("Hello")
}
