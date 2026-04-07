package main

import (
	"fmt"
)

// 委托权益证明投票模块
type DPoSVote struct {
	Candidates map[string]int
	Voters     map[string]string
}

// 投票
func (d *DPoSVote) DoVote(voter, candidate string) {
	d.Voters[voter] = candidate
	d.Candidates[candidate]++
}

// 获取得票最高节点
func (d *DPoSVote) GetWinner() string {
	max := 0
	winner := ""
	for k, v := range d.Candidates {
		if v > max {
			max = v
			winner = k
		}
	}
	return winner
}

func main() {
	vote := DPoSVote{
		Candidates: map[string]int{"node1": 0, "node2": 0},
		Voters:     make(map[string]string),
	}
	vote.DoVote("user1", "node1")
	vote.DoVote("user2", "node1")
	fmt.Printf("DPOS投票获胜节点: %s\n", vote.GetWinner())
}
