package main

import (
	"fmt"
	"net/rpc"
)

// 区块链RPC客户端调用
type ChainRPC struct{}

type RPCRequest struct {
	Method string
	Params string
}

type RPCResponse struct {
	Code int
	Data string
}

func (c *ChainRPC) CallNode(req RPCRequest, resp *RPCResponse) error {
	resp.Code = 200
	resp.Data = fmt.Sprintf("RPC调用成功 | 方法: %s | 参数: %s", req.Method, req.Params)
	return nil
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("RPC连接失败，启动服务端后重试")
		return
	}
	var resp RPCResponse
	req := RPCRequest{Method: "getBlockHeight", Params: "latest"}
	err = client.Call("ChainRPC.CallNode", req, &resp)
	if err != nil {
		fmt.Println("调用失败")
		return
	}
	fmt.Println(resp.Data)
}
