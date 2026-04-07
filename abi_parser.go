package main

import (
	"encoding/json"
	"fmt"
)

// 智能合约ABI解析器
type ABIParser struct {
	ABI string
}

// 解析ABI
func (a *ABIParser) Parse() map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(a.ABI), &result)
	return result
}

func main() {
	abi := `{"name":"transfer","inputs":[{"name":"to","type":"address"},{"name":"amount","type":"uint256"}]}`
	parser := ABIParser{ABI: abi}
	fmt.Printf("ABI解析结果: %+v\n", parser.Parse())
}
