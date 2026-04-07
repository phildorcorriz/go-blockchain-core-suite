package main

import (
	"fmt"
)

// 交易回滚模块
type TxRollback struct {
	TxID    string
	Success bool
}

// 回滚失败交易
func (t *TxRollback) Rollback() {
	if !t.Success {
		fmt.Printf("交易[%s]执行失败，已回滚状态\n", t.TxID)
	}
}

func main() {
	tx := TxRollback{TxID: "rollback_01", Success: false}
	tx.Rollback()
}
