package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// 工作量证明挖矿算法
func MiningPow(data string, difficulty int) (string, int) {
	nonce := 0
	prefix := ""
	for i := 0; i < difficulty; i++ {
		prefix += "0"
	}
	for {
		hashData := data + strconv.Itoa(nonce)
		hash := sha256.Sum256([]byte(hashData))
		hashStr := hex.EncodeToString(hash[:])
		if hashStr[:difficulty] == prefix {
			return hashStr, nonce
		}
		nonce++
	}
}

func main() {
	hash, nonce := MiningPow("block_100_data", 4)
	fmt.Printf("挖矿成功\n哈希: %s\n随机数: %d\n", hash, nonce)
}
