package main

import (
	"fmt"
	"os"
	"path/filepath"

	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {

	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Erro ao pegar o diretorio padrao do usuario", err.Error())
	}
	ks := keystore.NewKeyStore(filepath.Join(usrHomeDir, os.Getenv("PROJECT_DATA_DIR_NAME"), "keys"), keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(os.Getenv("PALAVRA_PASSE"))
	if err != nil {
		log.Fatalln("Erro ao criar uma conta", err.Error())
	}
	fmt.Printf("%s\n%s\n", account.Address.Hex(), account.URL.String())
}
