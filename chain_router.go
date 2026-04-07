package main

import (
	"fmt"
)

// 区块链路由模块
type ChainRouter struct {
	Routes map[string]string
}

// 注册路由
func (c *ChainRouter) RegisterRoute(path, handler string) {
	c.Routes[path] = handler
}

func main() {
	router := ChainRouter{Routes: make(map[string]string)}
	router.RegisterRoute("/block", "getBlockHandler")
	router.RegisterRoute("/tx", "sendTxHandler")
	fmt.Printf("路由表: %v\n", router.Routes)
}
