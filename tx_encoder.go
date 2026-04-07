package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// 交易编码解码模块
type TxEncode struct {
	TxID   string
	Amount float64
}

// 编码交易
func EncodeTx(tx TxEncode) string {
	jsonData, _ := json.Marshal(tx)
	return base64.StdEncoding.EncodeToString(jsonData)
}

// 解码交易
func DecodeTx(str string) (TxEncode, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return TxEncode{}, err
	}
	var tx TxEncode
	json.Unmarshal(data, &tx)
	return tx, nil
}

func main() {
	tx := TxEncode{TxID: "encode_01", Amount: 500}
	encodeStr := EncodeTx(tx)
	fmt.Printf("编码后: %s\n", encodeStr)
	decodeTx, _ := DecodeTx(encodeStr)
	fmt.Printf("解码后: %+v\n", decodeTx)
}
