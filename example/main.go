package main

import (
	"fmt"
	"github.com/luxfi/erc20-go/erc20"
	"github.com/luxfi/geth/common"
	"github.com/luxfi/geth/ethclient"
)

type Config struct {
	Network         string `yaml:"network"`
	ContractAddress string `yaml:"contract_address"`
}

func main() {
	conf := Config{
		Network:         "https://goerli.infura.io/v3/f81902bab6204bb88b03a1507225fe0a",
		ContractAddress: "0x2ac3c1d3e24b45c6c310534bc2dd84b5ed576335",
	}
	client, err := ethclient.Dial(conf.Network)
	if err != nil {
		fmt.Printf("Failed to connect to eth: %v", err)
		return
	}
	token, err := erc20.NewGGToken(common.HexToAddress(conf.ContractAddress), client)
	if err != nil {
		fmt.Printf("Failed to instantiate a Token contract: %v", err)
		return
	}
	name, err := token.Name(nil)
	if err != nil {
		fmt.Printf("Failed to get name: %v", err)
		return
	}
	fmt.Printf("name: %v\n", name)

	totalSupply, err := token.TotalSupply(nil)
	if err != nil {
		fmt.Printf("Failed to get name: %v", err)
		return
	}
	fmt.Printf("totalSupply: %v\n", totalSupply.String())
}
