package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

var testTxHash = "0x68dc1ca2c9e971fa9d0918ec38c8a2ec3800ddcbc22bef4da4cf8a6b5d5908ac"
var testTxData = `{"blockHash":"0x4116015755a1094322a91cfcc03d9b020458fe05ad367f05e3fc9ccad231b8d1","blockNumber":"0x65","from":"0xbc0e189507d624de860e1e5b3b06dee3475a70e9","gas":"0x33450","gasPrice":"0x5f5e100","hash":"0x5d8be6184c80e0f589a3d7b9609c37a5b75ee7119410447c78a4d26db2fde3af","input":"0x","nonce":"0x2","to":"0x1c0e8fc9decc4ae5c4947156ac87d5538bc124fb","transactionIndex":"0x1","value":"0x2386f26fc10000","type":"0x0","v":"0x1b","r":"0x8628cf6b85a9b4b8ab32be002c54f37dac5e2ef226c0b26c30eca331fe6ce8f0","s":"0x281c0ead6b40ad7c1f2eb6aa482243adc84b8b93bfcd3f58d6c8d2f3b3db7f4c"}`

func TestRLPEncodeTransaction(t *testing.T) {

	tx := new(types.Transaction)
	err := tx.UnmarshalJSON([]byte(testTxData))
	if err != nil {
		t.Fatal(err)
	}
	trb, err := json.MarshalIndent(tx, " ", "  ")
	fmt.Println(string(trb))
	rlpb, err := tx.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("RLP tx: %x\n", rlpb)
}
