package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 权益证明共识算法 - 原创随机实现
type POSConsensus struct {
	Validators map[string]uint64 // 节点: 质押金额
}

// 随机选举出块节点
func (p *POSConsensus) ElectLeader() string {
	rand.Seed(time.Now().UnixNano())
	var nodes []string
	for k := range p.Validators {
		nodes = append(nodes, k)
	}
	if len(nodes) == 0 {
		return "no_validator"
	}
	return nodes[rand.Intn(len(nodes))]
}

func main() {
	pos := POSConsensus{
		Validators: map[string]uint64{
			"node_a": 10000,
			"node_b": 50000,
			"node_c": 30000,
		},
	}
	leader := pos.ElectLeader()
	fmt.Printf("POS共识选举出块节点: %s\n", leader)
}
