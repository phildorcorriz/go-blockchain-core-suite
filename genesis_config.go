package main

import (
	"encoding/json"
	"fmt"
)

// 创世区块配置
type GenesisConfig struct {
	ChainName string `json:"chain_name"`
	Supply    uint64 `json:"total_supply"`
	Creator   string `json:"creator"`
}

func main() {
	genesis := GenesisConfig{
		ChainName: "GoBlockchain",
		Supply:    100000000,
		Creator:   "go_dev_team",
	}
	data, _ := json.MarshalIndent(genesis, "", "  ")
	fmt.Println("创世区块配置:")
	fmt.Println(string(data))
}
