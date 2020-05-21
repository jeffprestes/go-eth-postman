package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jeffprestes/goethereumhelper"
)

//SignAndPrepareTransaction signs the transaction and updates gas value and nonce
func SignAndPrepareTransaction(ksw *goethereumhelper.KeystoreWallet, passphrase string, client *ethclient.Client, unsignedTxJSON []byte) (signedTxJSON []byte, err error) {
	tx := new(types.Transaction)
	err = tx.UnmarshalJSON(unsignedTxJSON)
	if err != nil {
		return
	}
	transactor, err := ksw.NewKeyStoreTransactor(passphrase)
	if err != nil {
		return
	}
	log.Println("Atualizando o nonce...")
	err = ksw.UpdateKeyedTransactor(transactor, client, 0, 0)
	if err != nil {
		return
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return
	}

	tx = types.NewTransaction(transactor.Nonce.Uint64(),
		*tx.To(),
		tx.Value(),
		tx.Gas(),
		transactor.GasPrice,
		tx.Data())
	txSigned, err := ksw.SignTxWithPassphrase(passphrase, tx, chainID)
	if err != nil {
		return
	}
	signedTxJSON, err = txSigned.MarshalJSON()
	return
}
