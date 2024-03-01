 curl https://sepolia-rollup.arbitrum.io/rpc \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_getRawTransactionByHash","params":['$1'],"id":1,"jsonrpc":"2.0"}'
