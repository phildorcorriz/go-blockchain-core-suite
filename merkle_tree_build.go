package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 默克尔树构建 - 区块链交易验证核心
type MerkleTree struct {
	RootHash string
	Leaves   []string
}

// 生成默克尔根
func BuildMerkleRoot(leaves []string) string {
	if len(leaves) == 0 {
		return ""
	}
	for len(leaves) > 1 {
		var newLevel []string
		for i := 0; i < len(leaves); i += 2 {
			left := leaves[i]
			right := left
			if i+1 < len(leaves) {
				right = leaves[i+1]
			}
			hash := sha256.Sum256([]byte(left + right))
			newLevel = append(newLevel, hex.EncodeToString(hash[:]))
		}
		leaves = newLevel
	}
	return leaves[0]
}

func main() {
	txs := []string{"tx_001", "tx_002", "tx_003", "tx_004"}
	root := BuildMerkleRoot(txs)
	fmt.Printf("默克尔树根: %s\n", root)
}
