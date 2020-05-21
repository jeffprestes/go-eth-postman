package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jeffprestes/goethereumhelper"
)

//SubmitTransferTransactionsWithParsingLogs submits transfer transactions and parses his logs
func SubmitTransferTransactionsWithParsingLogs(client *ethclient.Client, contractToken IERC20Filterer, txsJSON ...[]byte) (logsGerais string, err error) {
	for _, txJSON := range txsJSON {
		receipt, errSubmit := SubmitTransaction(client, txJSON)
		if errSubmit != nil {
			err = errSubmit
			return
		}
		for _, log := range receipt.Logs {
			logEvento, errLog := contractToken.ParseTransfer(*log)
			if errLog != nil {
				return
			}
			logsGerais += fmt.Sprintf("\nFrom: %s - To: %s - Value: %d\n", logEvento.From.Hex(), logEvento.To.Hex(), logEvento.Value.Int64())
			logsGerais += fmt.Sprintf("Tx Serializada Processada:\nStatus: %d - Hash: %s\n", receipt.Status, receipt.TxHash.Hex())
		}
	}
	return
}

//SubmitTransaction submit a signed transaction stored at JSON text
func SubmitTransaction(client *ethclient.Client, txJSON []byte) (txReceipt *types.Receipt, err error) {
	txSerialized := new(types.Transaction)
	txSerialized.UnmarshalJSON(txJSON)

	client.SendTransaction(context.Background(), txSerialized)

	txReceipt, errResult := goethereumhelper.GetTransactionResult(client, txSerialized.Hash(), 20, 1)
	if errResult != nil {
		err = errResult
		return
	}
	return
}
