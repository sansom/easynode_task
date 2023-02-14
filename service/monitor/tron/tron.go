package tron

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/uduncloud/easynode_task/net/tron"
	"github.com/uduncloud/easynode_task/util"
)

func GetBlockTransactionCountByNumberFromChain(number int64, host, key string) int64 {
	h, _ := util.IntToHex(fmt.Sprintf("%v", number))
	res, err := tron.Eth_GetBlockTransactionCountByNumber(host, key, h)
	if err != nil {
		//fmt.Printf("Eth_GetBlockTransactionCountByNumber|number=%v|error=%v", number, err)
		return 0
	}

	r := gjson.Parse(res).Get("result").String()
	c, err := util.HexToInt(r)
	if err != nil {
		//fmt.Printf("Eth_GetBlockTransactionCountByNumber|number=%v|error=%v", v, err)
		return 0
	}

	return c
}
