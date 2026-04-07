package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// 区块链数字签名工具 - 原创实现
func CreateBlockSignature(data string) (string, *ecdsa.PrivateKey, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", nil, err
	}
	hash := sha256.Sum256([]byte(data))
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		return "", nil, err
	}
	return fmt.Sprintf("%x%x", r, s), privKey, nil
}

func main() {
	sig, key, err := CreateBlockSignature("block_transaction_1001")
	if err != nil {
		fmt.Println("签名失败:", err)
		return
	}
	fmt.Printf("区块签名成功\n私钥: %x\n签名值: %s\n", key.D, sig)
}
