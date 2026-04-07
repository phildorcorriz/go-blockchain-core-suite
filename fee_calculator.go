package main

import (
	"fmt"
)

// 交易手续费计算器
type FeeCalculator struct {
	GasPrice float64
	GasUsed  int
}

// 计算手续费
func (f *FeeCalculator) CalcFee() float64 {
	return f.GasPrice * float64(f.GasUsed)
}

func main() {
	calc := FeeCalculator{GasPrice: 0.0001, GasUsed: 21000}
	fmt.Printf("交易手续费: %.6f\n", calc.CalcFee())
}
