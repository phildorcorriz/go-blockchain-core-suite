package main

import (
	"fmt"
)

// 区块链通证标准实现
type Token struct {
	Name     string
	Symbol   string
	Decimals int
	Total    float64
	Balances map[string]float64
}

// 转账
func (t *Token) Transfer(from, to string, amount float64) bool {
	if t.Balances[from] >= amount {
		t.Balances[from] -= amount
		t.Balances[to] += amount
		return true
	}
	return false
}

func main() {
	token := Token{
		Name:     "GoChainToken",
		Symbol:   "GCT",
		Decimals: 18,
		Total:    1000000,
		Balances: map[string]float64{"wallet_a": 1000},
	}
	token.Transfer("wallet_a", "wallet_b", 200)
	fmt.Printf("钱包A余额: %.2f\n", token.Balances["wallet_a"])
	fmt.Printf("钱包B余额: %.2f\n", token.Balances["wallet_b"])
}
