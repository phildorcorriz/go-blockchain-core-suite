package main

import (
	"fmt"
)

// 区块奖励计算
func CalcBlockReward(height uint64) float64 {
	if height < 100000 {
		return 10.0
	} else if height < 1000000 {
		return 5.0
	}
	return 2.5
}

func main() {
	fmt.Printf("高度1000奖励: %.1f\n", CalcBlockReward(1000))
	fmt.Printf("高度500000奖励: %.1f\n", CalcBlockReward(500000))
}
