package ether

import (
	"log"
	"testing"
)

func TestEth_GetBlockNumber(t *testing.T) {
	log.Println(Eth_GetBlockNumber("https://eth-mainnet.g.alchemy.com/v2", "demo"))
}
