 curl http://localhost:8449 \
  -X POST \
  -H "Content-Type: application/json" \
  --data '{"method":"eth_getTransactionByHash","params":['$1'],"id":1,"jsonrpc":"2.0"}'
