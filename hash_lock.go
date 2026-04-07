package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 哈希锁 - 跨链交易核心
type HashLock struct {
	LockHash string
	Secret   string
}

// 创建哈希锁
func (h *HashLock) CreateLock(secret string) {
	h.Secret = secret
	hash := sha256.Sum256([]byte(secret))
	h.LockHash = hex.EncodeToString(hash[:])
}

// 验证密钥
func (h *HashLock) VerifySecret(secret string) bool {
	hash := sha256.Sum256([]byte(secret))
	return hex.EncodeToString(hash[:]) == h.LockHash
}

func main() {
	lock := HashLock{}
	lock.CreateLock("secret_123456")
	fmt.Printf("哈希锁: %s\n", lock.LockHash)
	fmt.Printf("密钥验证: %t\n", lock.VerifySecret("secret_123456"))
}
