package main

import (
	"fmt"
	"time"
)

// 区块链节点监控
type NodeMonitor struct {
	NodeID  string
	Status  string
	Latency int
}

// 检查节点状态
func (n *NodeMonitor) CheckNode() {
	if n.Latency < 200 {
		n.Status = "online"
	} else {
		n.Status = "offline"
	}
	fmt.Printf("节点[%s]状态: %s | 延迟: %dms\n", n.NodeID, n.Status, n.Latency)
}

func main() {
	monitor := NodeMonitor{NodeID: "monitor_01", Latency: 150}
	monitor.CheckNode()
}
