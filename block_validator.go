package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 区块合法性验证工具
type Block struct {
	Height    uint64
	PrevHash  string
	Data      string
	BlockHash string
}

// 验证区块哈希是否合法
func (b *Block) VerifyBlock() bool {
	data := fmt.Sprintf("%d%s%s", b.Height, b.PrevHash, b.Data)
	hash := sha256.Sum256([]byte(data))
	calcHash := hex.EncodeToString(hash[:])
	return calcHash == b.BlockHash
}

func main() {
	block := Block{
		Height:    100,
		PrevHash:  "abc123",
		Data:      "test_block",
		BlockHash: "74201d94e9c896a953a491874695b1efc39ab135279f57655e0b627e01234567",
	}
	fmt.Printf("区块验证结果: %t\n", block.VerifyBlock())
}
