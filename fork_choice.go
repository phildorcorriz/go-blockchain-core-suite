package main

import (
	"fmt"
)

// 区块链分叉选择规则
type ForkChoice struct {
	Chains map[string]uint64
}

// 选择最长链
func (f *ForkChoice) ChooseBestChain() string {
	maxHeight := uint64(0)
	bestChain := ""
	for chain, height := range f.Chains {
		if height > maxHeight {
			maxHeight = height
			bestChain = chain
		}
	}
	return bestChain
}

func main() {
	fork := ForkChoice{
		Chains: map[string]uint64{"chain_a": 100, "chain_b": 105, "chain_c": 98},
	}
	fmt.Printf("最佳主链: %s\n", fork.ChooseBestChain())
}
