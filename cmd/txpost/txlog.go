package main

import (
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//ProcessLogs process and parse transaction logs
func ProcessLogs(client *ethclient.Client, contractAddress string, contractToken IERC20Filterer, logsTx []*types.Log) {
	//IERC20Filterer
	for index, logTx := range logsTx {
		log.Println("Processando log ", index+1, "...")
		logEvento, errLog := contractToken.ParseTransfer(*logTx)
		if errLog != nil {
			return
		}
		log.Printf("Contract: %s - From: %s - To: %s - Value: %d\n", contractAddress, logEvento.From.Hex(), logEvento.To.Hex(), logEvento.Value.Int64())
	}
}
