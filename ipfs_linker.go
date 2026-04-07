package main

import (
	"fmt"
)

// 区块链+IPFS链接模块
type IPFSLinker struct {
	BlockHash string
	IPFSCID   string
}

// 绑定IPFS文件
func (i *IPFSLinker) BindFile() string {
	return fmt.Sprintf("区块[%s]已绑定IPFS文件: %s", i.BlockHash, i.IPFSCID)
}

func main() {
	linker := IPFSLinker{BlockHash: "block_001", IPFSCID: "QmXYZ123"}
	fmt.Println(linker.BindFile())
}
