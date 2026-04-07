package main

import (
	"fmt"
)

// P2P消息协议
type P2PMessage struct {
	Type    string
	Content string
	From    string
}

// 打包消息
func PackMessage(msgType, content, from string) P2PMessage {
	return P2PMessage{
		Type:    msgType,
		Content: content,
		From:    from,
	}
}

func main() {
	msg := PackMessage("BLOCK", "new_block_100", "node_01")
	fmt.Printf("P2P消息: %+v\n", msg)
}
