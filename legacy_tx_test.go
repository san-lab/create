package main

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestRawTx(t *testing.T) {
	rpcUrl := "http://localhost:8449"
	privateKey, err := crypto.HexToECDSA("c522c068090d4e888dadbab9967fd81a79a451aff84dce2040df59ad5a6ce1e8")
	if err != nil {
		t.Fatal(err)
	}
	nonce := uint64(1)
	//fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	toAddress := common.HexToAddress("0x1c0e8FC9DEcC4Ae5C4947156aC87D5538bC124fb")

	value := big.NewInt(100000000000000000) // in wei (0.1 eth)

	gasLimit := uint64(210000) // in units
	gasPrice := big.NewInt(100000000)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	legacySigner := new(types.FrontierSigner)

	signedTx, err := types.SignTx(tx, legacySigner, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	//t.Log(signedTx.RawSignatureValues())
	traw, err := signedTx.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Raw tx: %x\n", traw)
	pre := "curl " + rpcUrl + `  -X POST   -H "Content-Type: application/json"   --data '{"jsonrpc":"2.0", "method":"eth_sendRawTransaction","params":["0x`
	post := `"],"id":1}'`
	fmt.Printf("%s%x%s\n", pre, traw, post)
}
