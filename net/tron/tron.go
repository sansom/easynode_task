package tron

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
)

// Eth_GetBlockNumber 获取最新区块高度
func Eth_GetBlockNumber(host string, token string) (string, error) {

	//url := "https://eth-mainnet.g.alchemy.com/v2/demo"

	host = fmt.Sprintf("%v/%v", host, "jsonrpc")
	query := `
{
    "id": 1,
    "jsonrpc": "2.0",
    "method": "eth_blockNumber"
}
`
	payload := strings.NewReader(query)

	req, err := http.NewRequest("POST", host, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("TRON_PRO_API_KEY", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}
	if gjson.ParseBytes(body).Get("error").Exists() {
		return "", errors.New(string(body))
	}
	return string(body), nil
}