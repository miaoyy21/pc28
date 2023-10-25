package client

import "math"

var SN28 = []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}

var STDS1000 = map[int32]int32{
	0:  1,
	1:  3,
	2:  6,
	3:  10,
	4:  15,
	5:  21,
	6:  28,
	7:  36,
	8:  45,
	9:  55,
	10: 63,
	11: 69,
	12: 73,
	13: 75,
	14: 75,
	15: 73,
	16: 69,
	17: 63,
	18: 55,
	19: 45,
	20: 36,
	21: 28,
	22: 21,
	23: 15,
	24: 10,
	25: 6,
	26: 3,
	27: 1,
}

func ofM1Gold(g int64) int64 {
	if g < 1<<22 {
		// 4194304
		return g / 75
	} else if g < 1<<23 {
		// 8388608
		return g / 100
	} else if g < 1<<24 {
		// 16777216
		return g / 125
	} else if g < 1<<25 {
		// 33554432
		return g / 150
	} else if g < 1<<26 {
		// 67108864
		return g / 175
	} else if g < 1<<27 {
		// 134217728
		return g / 200
	} else if g < 1<<28 {
		// 268435456
		return g / 225
	} else {
		return g / 250
	}
}

func ofGold(fGold float64) int32 {
	var iGold int32
	if fGold >= 1<<16 {
		iGold = int32(math.Round(fGold/2000.0) * 2000)
	} else if fGold >= 1<<15 {
		iGold = int32(math.Round(fGold/1500.0) * 1500)
	} else if fGold >= 1<<14 {
		iGold = int32(math.Round(fGold/1000.0) * 1000)
	} else {
		iGold = int32(math.Round(fGold/500.0) * 500)
	}

	return iGold
}