package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 简易零知识证明实现
func ZKProof(secret string) string {
	hash := sha256.Sum256([]byte(secret))
	return hex.EncodeToString(hash[:])
}

func main() {
	proof := ZKProof("zk_secret_001")
	fmt.Printf("零知识证明值: %s\n", proof)
}
