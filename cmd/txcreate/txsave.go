package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//GenerateTransferTxAsJSON Generates a trx and serialized it in JSON Format
func GenerateTransferTxAsJSON(
	parsedABI abi.ABI,
	smartContractAddress, accountTo common.Address,
	chainID, valueToken uint64) (txJSON []byte, err error) {

	contractMethodParameters, err := parsedABI.Pack("transfer", accountTo, big.NewInt(int64(valueToken)))
	if err != nil {
		log.Fatalln("Erro ao fazer o encoding da trx", err.Error())
		return
	}

	tx := types.NewTransaction(0, smartContractAddress, big.NewInt(0), 50000, big.NewInt(0), contractMethodParameters)

	txJSON, err = tx.MarshalJSON()
	if err != nil {
		log.Fatalln("Error serializing the transaction", err.Error())
		return
	}
	return
}
