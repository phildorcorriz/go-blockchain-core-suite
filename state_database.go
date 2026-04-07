package main

import (
	"fmt"
)

// 区块链状态数据库
type StateDB struct {
	State map[string]string
}

// 设置状态
func (s *StateDB) SetState(key, value string) {
	s.State[key] = value
}

// 获取状态
func (s *StateDB) GetState(key string) string {
	return s.State[key]
}

func main() {
	db := StateDB{State: make(map[string]string)}
	db.SetState("balance_wallet1", "10000")
	fmt.Printf("钱包余额: %s\n", db.GetState("balance_wallet1"))
}
