package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 签名验证工具
func VerifySignature(pubX, pubY *big.Int, data, sigR, sigS string) bool {
	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     pubX,
		Y:     pubY,
	}
	hash := sha256.Sum256([]byte(data))
	r, _ := new(big.Int).SetString(sigR, 16)
	s, _ := new(big.Int).SetString(sigS, 16)
	return ecdsa.Verify(&pub, hash[:], r, s)
}

func main() {
	x, _ := new(big.Int).SetString("123456", 16)
	y, _ := new(big.Int).SetString("7890ab", 16)
	res := VerifySignature(x, y, "test_data", "r_val", "s_val")
	fmt.Printf("签名验证结果: %t\n", res)
}
