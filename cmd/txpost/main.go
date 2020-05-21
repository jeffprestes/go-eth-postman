package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jeffprestes/goethereumhelper"
)

func main() {
	log.Println("Abrindo keystore...")
	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Erro ao pegar o diretorio padrao do usuario", err.Error())
	}
	ks := keystore.NewKeyStore(filepath.Join(usrHomeDir, os.Getenv("PROJECT_DATA_DIR_NAME"), "keys"), keystore.StandardScryptN, keystore.StandardScryptP)
	ksw, err := goethereumhelper.NewKeystoreWallet(ks, os.Getenv("SENDER_ACCOUNT"))
	if err != nil {
		log.Fatalln("Conta não existente na keystore", err.Error())
	}

	log.Println("Conectando no Infura...")
	client, err := GetEthClient()
	if err != nil {
		log.Fatalln("Erro ao se conectar com o Ethereum via Infura", err.Error())
	}

	tx1 := `{"nonce":"0x0","gasPrice":"0x0","gas":"0xc350","to":"0x1f63fbf4e00824dc00481c2a3f8bd5cc9d055762","value":"0x0","input":"0xa9059cbb000000000000000000000000263c3ab7e4832edf623fbdd66acee71c028ff5910000000000000000000000000000000000000000000000000000000000000064","v":"0x0","r":"0x0","s":"0x0","hash":"0xa267046cf6f24aa9193fd8f982d27b963908531500e397e6498f627edf7e31c3"}`
	tx2 := `{"nonce":"0x0","gasPrice":"0x0","gas":"0xc350","to":"0x1f63fbf4e00824dc00481c2a3f8bd5cc9d055762","value":"0x0","input":"0xa9059cbb0000000000000000000000004e063fac22663e02693e22065a239e49bc1056dc0000000000000000000000000000000000000000000000000000000000000064","v":"0x0","r":"0x0","s":"0x0","hash":"0x9725b11d7590942dd46dcc52573eda96bb6eba79ec857a27974301570abf01f5"}`

	var trxToProcess []string
	trxToProcess = append(trxToProcess, tx1)
	trxToProcess = append(trxToProcess, tx2)

	logs, err := ProcessTransactions(trxToProcess, ksw, os.Getenv("PALAVRA_PASSE"), client)
	if err != nil {
		log.Fatalln("Erro no processamento das transações", err.Error())
	}

	contractTokenAddress := common.HexToAddress(os.Getenv("TF_CONTRACT_ADDRESS"))
	contractToken, err := GetContract(client, contractTokenAddress)
	if err != nil {
		log.Fatalln("Erro ao instanciar o contrato do token", err.Error())
	}

	ProcessLogs(client, os.Getenv("TF_CONTRACT_ADDRESS"), contractToken.IERC20Filterer, logs)
}
