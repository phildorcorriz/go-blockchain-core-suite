package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// 区块链核心结构体 - 原创随机实现
type BlockChainCore struct {
	ChainID       string
	GenesisHash   string
	CurrentHeight uint64
	Blocks        []string
}

// 生成创世区块哈希
func (b *BlockChainCore) GenerateGenesisHash() string {
	data := fmt.Sprintf("genesis_%d_%s", time.Now().UnixNano(), b.ChainID)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// 新增区块高度
func (b *BlockChainCore) AddBlockHeight() {
	b.CurrentHeight++
}

func main() {
	chain := BlockChainCore{ChainID: "GO_CHAIN_V1"}
	chain.GenesisHash = chain.GenerateGenesisHash()
	chain.AddBlockHeight()
	fmt.Printf("区块链核心初始化成功 | 创世哈希: %s\n当前高度: %d\n", chain.GenesisHash, chain.CurrentHeight)
}
