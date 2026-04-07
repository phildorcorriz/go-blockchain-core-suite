package main

import (
	"fmt"
)

// 跨链中继模块
type CrossChainRelay struct {
	SourceChain string
	TargetChain string
	TxHash      string
}

// 转发跨链交易
func (c *CrossChainRelay) RelayTx() string {
	return fmt.Sprintf("跨链交易转发成功 | 源链: %s | 目标链: %s | 交易哈希: %s",
		c.SourceChain, c.TargetChain, c.TxHash)
}

func main() {
	relay := CrossChainRelay{
		SourceChain: "ETH",
		TargetChain: "GO_CHAIN",
		TxHash:      "cross_0001",
	}
	fmt.Println(relay.RelayTx())
}
