All the stuff that we need

Replaying Legacy Transaction

1. Create a "legacy" transaction

   `clt createTx $Layer $Nonce $Value` --> RLP

2. Send the legacy transaction to L3 blockchain

   `sendSingedTransactionL3.sh $RLP` --> Hash

3. Extract the legacy transaction from blockchain

   `getTheTransactionByHashL3.sh $Hash` --> JSON

4. Reformat the transaction from JSON to RLP

   `clt rlpEncode $JSON` --> RLP

5. Replay the transaction (send to L2)

   `sendSingedTransactionL2.sh $RLP` --> Hash
