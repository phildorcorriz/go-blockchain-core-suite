package main

import (
	"fmt"
)

// 交易池管理模块
type TxPool struct {
	PendingTxs []string
	MaxSize    int
}

// 添加交易到池
func (t *TxPool) AddTx(tx string) bool {
	if len(t.PendingTxs) < t.MaxSize {
		t.PendingTxs = append(t.PendingTxs, tx)
		return true
	}
	return false
}

// 清空交易池
func (t *TxPool) ClearPool() {
	t.PendingTxs = []string{}
}

func main() {
	pool := TxPool{MaxSize: 10}
	pool.AddTx("tx_001")
	pool.AddTx("tx_002")
	fmt.Printf("交易池待打包交易: %v\n", pool.PendingTxs)
	pool.ClearPool()
	fmt.Printf("清空后交易池: %v\n", pool.PendingTxs)
}
