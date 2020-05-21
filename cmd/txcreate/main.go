package main

import (
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	contractAddress := common.HexToAddress(os.Getenv("TF_CONTRACT_ADDRESS"))
	accountTo := common.HexToAddress("0x263C3Ab7E4832eDF623fBdD66ACee71c028Ff591")

	parsedABI, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		log.Fatalln("Erro ao fazer o parser do ABI", err.Error())
	}

	//txJSON, err := GenerateTransferSignedTxAsText(ksw, txOpts, parsedABI, os.Getenv("PALAVRA_PASSE"), contractAddress, accountTo, nonce, chainID.Uint64(), 100)
	txJSON, err := GenerateTransferTxAsJSON(parsedABI, contractAddress, accountTo, 4, 100)
	if err != nil {
		log.Fatalln("Erro ao serializar a tx1", err.Error())
	}
	log.Printf("\nTx1:\n%s\n", string(txJSON))

	accountTo = common.HexToAddress("0x4e063FAc22663e02693E22065A239E49Bc1056dC")

	//txJSON2, err := GenerateTransferSignedTxAsText(ksw, txOpts, parsedABI, os.Getenv("PALAVRA_PASSE"), contractAddress, accountTo, nonce, chainID.Uint64(), 100)
	txJSON2, err := GenerateTransferTxAsJSON(parsedABI, contractAddress, accountTo, 4, 100)
	if err != nil {
		log.Fatalln("Erro ao serializar a tx2", err.Error())
	}
	log.Printf("\nTx2:\n%s\n", string(txJSON2))

}
