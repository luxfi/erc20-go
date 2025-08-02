package main

import (
	"fmt"
	"github.com/luxfi/geth/common"
)

func main() {
	addr := common.HexToAddress("0x0000000000000000000000000000000000000000")
	fmt.Println("Address:", addr)
}