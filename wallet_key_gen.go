package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

// 区块链钱包密钥生成
func GenerateWalletKeys() (string, string, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}
	pub := priv.PublicKey
	privStr := fmt.Sprintf("%x", priv.D)
	pubStr := fmt.Sprintf("%x%x", pub.X, pub.Y)
	return privStr, pubStr, nil
}

func main() {
	priv, pub, err := GenerateWalletKeys()
	if err != nil {
		fmt.Println("密钥生成失败")
		return
	}
	fmt.Println("=== 区块链钱包密钥 ===")
	fmt.Printf("私钥: %s\n公钥: %s\n", priv, pub)
}
