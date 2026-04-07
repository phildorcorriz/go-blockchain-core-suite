package main

import (
	"fmt"
	"time"
)

// 区块链数据同步模块
type ChainSync struct {
	LocalHeight  uint64
	RemoteHeight uint64
	SyncStatus   bool
}

// 同步区块数据
func (c *ChainSync) StartSync() {
	if c.RemoteHeight > c.LocalHeight {
		fmt.Printf("开始同步区块 | 本地高度: %d 远程高度: %d\n", c.LocalHeight, c.RemoteHeight)
		for c.LocalHeight < c.RemoteHeight {
			c.LocalHeight++
			fmt.Printf("同步完成高度: %d\n", c.LocalHeight)
			time.Sleep(50 * time.Millisecond)
		}
		c.SyncStatus = true
		fmt.Println("区块链数据同步完成！")
	}
}

func main() {
	sync := ChainSync{LocalHeight: 100, RemoteHeight: 105}
	sync.StartSync()
}
