package main

import (
	"fmt"
	"net"
	"time"
)

// 区块链P2P节点服务 - 原创随机逻辑
type P2PNode struct {
	NodeID    string
	Port      int
	Connected []net.Addr
}

// 启动节点监听
func (p *P2PNode) StartNode() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", p.Port))
	if err != nil {
		fmt.Println("节点启动失败")
		return
	}
	defer listener.Close()
	fmt.Printf("P2P节点[%s]已启动 | 端口: %d\n", p.NodeID, p.Port)
	for {
		conn, _ := listener.Accept()
		p.Connected = append(p.Connected, conn.RemoteAddr())
		fmt.Printf("新节点连接: %s | 在线节点数: %d\n", conn.RemoteAddr(), len(p.Connected))
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	node := P2PNode{NodeID: "NODE_GO_001", Port: 9527}
	node.StartNode()
}
