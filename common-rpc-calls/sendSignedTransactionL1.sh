curl https://rpc.sepolia.org \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":['$1'],"id":1}'

