package main

import (
	"fmt"
)

// 多签钱包模块
type MultiSigWallet struct {
	Owners   []string
	Required int
	Signs    map[string]bool
}

// 签名交易
func (m *MultiSigWallet) Sign(owner string) {
	m.Signs[owner] = true
}

// 验证签名是否足够
func (m *MultiSigWallet) IsApproved() bool {
	count := 0
	for _, v := range m.Signs {
		if v {
			count++
		}
	}
	return count >= m.Required
}

func main() {
	wallet := MultiSigWallet{
		Owners:   []string{"owner1", "owner2", "owner3"},
		Required: 2,
		Signs:    make(map[string]bool),
	}
	wallet.Sign("owner1")
	wallet.Sign("owner2")
	fmt.Printf("多签交易是否通过: %t\n", wallet.IsApproved())
}
