package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Usage: main <cmd> (<layer> <nonce> <value> || <txdata>)
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: clt (createTx || rlpEncode) (<layer> <nonce> <value> || <txdata>)")
		return
	}
	switch os.Args[1] {
	case "createTx":
		layer := os.Args[2]
		nonce, ok := new(big.Int).SetString(os.Args[3], 10)
		if !ok {
			fmt.Println("Error parsing nonce")
			return
		}
		//parse value as a float
		var value float64
		_, err := fmt.Sscanf(os.Args[4], "%f", &value)
		if err != nil {
			fmt.Println("Error parsing value")
			return
		}

		ethvalue := big.NewInt(int64(value * 1e18))

		fmt.Println(GenerateSendRawTxCall(layer, nonce, ethvalue))
	case "rlpEncode":
		rlp, err := RLPEncodeTransaction(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println()
		fmt.Println("Rlp encoded transaction:")
		fmt.Println(`\"` + rlp + `\"`)
	default:
		fmt.Println("Unknown command")
		fmt.Println("Known commands: createTx, rlpEncode")
	}
}

// Returns the Hex of the RLP encoded tx, its hash, and, hopefully, no error
func CreateRawTxRLP(privateKeyString string, nonce uint64, toAddressString string, value *big.Int, gasLimit uint64, gasPrice *big.Int) (string, string, error) {
	privateKey, err := crypto.HexToECDSA("c522c068090d4e888dadbab9967fd81a79a451aff84dce2040df59ad5a6ce1e8")
	if err != nil {
		return "", "", err
	}
	toAddress := common.HexToAddress(toAddressString)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	legacySigner := new(types.FrontierSigner)

	signedTx, err := types.SignTx(tx, legacySigner, privateKey)
	if err != nil {
		return "", "", err
	}

	js, _ := json.MarshalIndent(signedTx, " ", "  ")
	fmt.Println(string(js))

	//t.Log(signedTx.RawSignatureValues())
	traw, err := signedTx.MarshalBinary()
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("0x%x", traw), signedTx.Hash().Hex(), nil
}

func GenerateSendRawTxCall(layer string, nonce, value *big.Int) string {
	var rpcurl string
	switch layer {
	case "L2", "arb-sepolia":
		rpcurl = "https://sepolia-rollup.arbitrum.io/rpc"
	case "L3":
		rpcurl = "http://localhost:8449"
	}
	rlps, txhs, err := CreateRawTxRLP("c522c068090d4e888dadbab9967fd81a79a451aff84dce2040df59ad5a6ce1e8", nonce.Uint64(), "0x5F18bD40CF6cBbf034ff3d2003576B95E73D32e3", value, 210000, big.NewInt(100000000))
	if err != nil {
		return err.Error()
	}
	fmt.Println("New transaction hash: ", txhs)
	return fmt.Sprintf("curl %s  -X POST   -H \"Content-Type: application/json\"   --data '{\"jsonrpc\":\"2.0\", \"method\":\"eth_sendRawTransaction\",\"params\":[\"%s\"],\"id\":1}'", rpcurl, rlps)

}

func RLPEncodeTransaction(txjson string) (string, error) {
	tx := new(types.Transaction)
	err := tx.UnmarshalJSON([]byte(txjson))
	if err != nil {
		return "", err
	}
	rlpb, err := tx.MarshalBinary()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0x%x", rlpb), nil
}
