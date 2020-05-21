package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jeffprestes/goethereumhelper"
)

//GetEthClient gets ethereum client connection via a provider like Infura
func GetEthClient() (client *ethclient.Client, err error) {
	client, err = goethereumhelper.GetCustomNetworkClient(os.Getenv("ETH_CLIENT_URL"))
	if err != nil {
		log.Println("GetEthClient - Client ", err)
		return
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("GetEthClient - ChainID ", err)
		return
	}
	if chainID.Int64() < 1 {
		err = errors.New("It was not able to connect to an Ethereum node and get ChainID")
		log.Println("GetEthClient - ChainID ", err)
		return
	}
	return
}

//SignAndSubmitTransaction signs, prepare and submit a transaction
func SignAndSubmitTransaction(ksw *goethereumhelper.KeystoreWallet, passphrase string, client *ethclient.Client, tx []byte) (receipt *types.Receipt, err error) {
	log.Println("Assinando e preparando transação...")
	txSigned, err := SignAndPrepareTransaction(ksw,
		passphrase,
		client, tx)
	if err != nil {
		return
	}
	log.Printf("\nTx Signed:\n%s\n", string(txSigned))

	log.Println("Submetendo transação...")
	receipt, err = SubmitTransaction(client, txSigned)
	if err != nil {
		return
	}
	return
}

//ProcessTransactions process a batch of transactions
func ProcessTransactions(transactions []string,
	ksw *goethereumhelper.KeystoreWallet,
	passphrase string,
	client *ethclient.Client) (logs []*types.Log, err error) {
	for index, tx := range transactions {
		log.Println("Processando transação ", index+1, "...")
		receipt, errTx := SignAndSubmitTransaction(ksw, passphrase, client, []byte(tx))
		if errTx != nil {
			err = errTx
			return
		}
		logs = append(logs, receipt.Logs...)
		log.Printf("Tx Serializada Processada:\nStatus: %d - Hash: %s\n\n", receipt.Status, receipt.TxHash.Hex())
	}
	return
}

//GetContract returns contract instance connected to Smart Contract at Ethereum
func GetContract(client *ethclient.Client, contractAddress common.Address) (assetContract *IERC20, err error) {
	assetContract, err = NewIERC20(contractAddress, client)
	if err != nil {
		log.Println("GetContract - NewIERC20 ", err)
		return
	}
	return
}
