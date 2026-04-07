package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// 区块头生成工具
type BlockHeader struct {
	Version   int
	PrevHash  string
	MerkleRoot string
	Timestamp int64
	Difficulty int
	Nonce     int
}

// 生成区块头哈希
func (b *BlockHeader) GenHeaderHash() string {
	data := fmt.Sprintf("%d%s%s%d%d%d",
		b.Version, b.PrevHash, b.MerkleRoot, b.Timestamp, b.Difficulty, b.Nonce)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func main() {
	header := BlockHeader{
		Version:    1,
		PrevHash:   "prev_hash",
		MerkleRoot: "merkle_root",
		Timestamp:  time.Now().Unix(),
		Difficulty: 3,
		Nonce:      12345,
	}
	fmt.Printf("区块头哈希: %s\n", header.GenHeaderHash())
}
