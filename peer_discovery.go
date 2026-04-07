package main

import (
	"fmt"
	"math/rand"
	"time"
)

// P2P节点发现模块
type PeerDiscovery struct {
	BootNodes []string
	Peers     []string
}

// 随机发现节点
func (p *PeerDiscovery) DiscoverPeers() {
	rand.Seed(time.Now().UnixNano())
	for _, node := range p.BootNodes {
		if rand.Intn(2) == 0 {
			p.Peers = append(p.Peers, node)
			fmt.Printf("发现节点: %s\n", node)
		}
	}
	fmt.Printf("节点发现完成 | 有效节点数: %d\n", len(p.Peers))
}

func main() {
	discovery := PeerDiscovery{
		BootNodes: []string{"node1:9001", "node2:9001", "node3:9001"},
	}
	discovery.DiscoverPeers()
}
