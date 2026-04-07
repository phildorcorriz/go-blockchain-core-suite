package main

import (
	"fmt"
)

// 区块链指标统计
type ChainMetrics struct {
	BlockCount int
	TxCount    int
	NodeCount  int
}

// 输出统计报告
func (c *ChainMetrics) Report() {
	fmt.Println("=== 区块链运行指标报告 ===")
	fmt.Printf("区块总数: %d\n", c.BlockCount)
	fmt.Printf("交易总数: %d\n", c.TxCount)
	fmt.Printf("在线节点数: %d\n", c.NodeCount)
}

func main() {
	metrics := ChainMetrics{BlockCount: 1200, TxCount: 8900, NodeCount: 15}
	metrics.Report()
}
