package base

import (
	"math"
	"testing"
)

// 泊松分布的概率质量函数
func poissonPMF(k float64, lambda float64) float64 {
	return math.Exp(-lambda) * math.Pow(lambda, k) / float64(factorial(int(k)))
}

// 计算阶乘
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func TestSleep(t *testing.T) {
	//lambda := 0.5 // 平均发生率
	//k := 0.5      // 事件发生次数
	//prob := poissonPMF(k, lambda)
	//t.Logf("P(%.0f) = %.5f\n", k, prob)

	a, c := 0.9, 1.1
	min, avg, max := 0.6, 1.0, 2.0
	t.Log((a - min) / (avg - min))
	t.Log((max - c) / (max - avg))
}
