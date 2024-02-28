curl http://localhost:8449 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_getTransactionCount","params":["0xBC0E189507D624De860E1e5b3B06DEE3475A70e9","latest"],"id":1,"jsonrpc":"2.0"}'

