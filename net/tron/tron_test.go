package tron

import (
	"log"
	"testing"
)

func TestEth_GetBlockNumber(t *testing.T) {
	log.Println(Eth_GetBlockNumber("https://api.trongrid.io", "244f918d-56b5-4a16-9665-9637598b1223"))
}

func TestEth_GetBlockTransactionCountByNumber(t *testing.T) {

	log.Println(Eth_GetBlockTransactionCountByNumber("https://api.trongrid.io", "244f918d-56b5-4a16-9665-9637598b1223", "0xF9CC56"))
}
