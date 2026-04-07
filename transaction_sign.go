package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// 区块链交易签名模块
type Transaction struct {
	TxID    string
	From    string
	To      string
	Amount  float64
	Time    int64
	SignStr string
}

// 签名交易
func (t *Transaction) SignTx() {
	data := fmt.Sprintf("%s%s%s%f%d", t.TxID, t.From, t.To, t.Amount, t.Time)
	hash := sha256.Sum256([]byte(data))
	t.SignStr = hex.EncodeToString(hash[:])
}

func main() {
	tx := Transaction{
		TxID:   "TX_RAND_9876",
		From:   "wallet_abc",
		To:     "wallet_def",
		Amount: 9.99,
		Time:   time.Now().Unix(),
	}
	tx.SignTx()
	fmt.Printf("交易签名完成\n交易ID: %s\n签名: %s\n", tx.TxID, tx.SignStr)
}
