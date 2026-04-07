package main

import (
	"fmt"
	"sort"
)

// 交易优先级排序
type Tx struct {
	Hash    string
	Fee     float64
}

func SortTxsByFee(txs []Tx) []Tx {
	sort.Slice(txs, func(i, j int) bool {
		return txs[i].Fee > txs[j].Fee
	})
	return txs
}

func main() {
	txs := []Tx{
		{"tx1", 0.01},
		{"tx2", 0.05},
		{"tx3", 0.02},
	}
	sorted := SortTxsByFee(txs)
	fmt.Println("按手续费优先级排序:")
	for _, t := range sorted {
		fmt.Printf("%s - %.3f\n", t.Hash, t.Fee)
	}
}
