package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestRawTx(t *testing.T) {
	rpcUrl := "http://localhost:8449"
	privateKey := "c522c068090d4e888dadbab9967fd81a79a451aff84dce2040df59ad5a6ce1e8"
	nonce := uint64(2)
	toAddress := "0x1c0e8FC9DEcC4Ae5C4947156aC87D5538bC124fb"
	value := big.NewInt(10000000000000000) // in wei (0.01 eth)

	gasLimit := uint64(210000) // in units
	gasPrice := big.NewInt(100000000)

	rlps, txhs, err := CreateRawTxRLP(privateKey, nonce, toAddress, value, gasLimit, gasPrice)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Hash: ", txhs)
	fmt.Printf("Raw tx: %s\n", rlps)
	pre := "curl " + rpcUrl + `  -X POST   -H "Content-Type: application/json"   --data '{"jsonrpc":"2.0", "method":"eth_sendRawTransaction","params":["`
	post := `"],"id":1}'`
	fmt.Printf("%s%s%s\n", pre, rlps, post)
}
