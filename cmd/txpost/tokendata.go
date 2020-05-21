package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func getTokenDetails(assetContract *IERC20) (symbol string, err error) {

	symbol, err = assetContract.Symbol(nil)
	if err != nil {
		log.Println("getTokenDetails - assetContract.Symbol ", err)
		return
	}
	return
}

func getTokenBalance(assetContract *IERC20, contractAddress, sender, to common.Address) (contractBalance, senderBalance, receiverBalance *big.Int, err error) {

	//log.Println("Sucesso receipt ", receipt.Status)
	contractBalance, err = assetContract.BalanceOf(nil, contractAddress)
	if err != nil {
		log.Println("getTokenBalance - contractBalance.BalanceOf ", err)
		return
	}
	senderBalance, err = assetContract.BalanceOf(nil, sender)
	if err != nil {
		log.Println("getTokenBalance - senderBalance.BalanceOf ", err)
		return
	}
	receiverBalance, err = assetContract.BalanceOf(nil, to)
	if err != nil {
		log.Println("getTokenBalance - receiverBalance.BalanceOf ", err)
		return
	}
	return
}
