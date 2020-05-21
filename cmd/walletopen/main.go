package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jeffprestes/goethereumhelper"
)

func main() {
	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Erro ao pegar o diretorio padrao do usuario", err.Error())
	}
	ks := keystore.NewKeyStore(filepath.Join(usrHomeDir, os.Getenv("PROJECT_DATA_DIR_NAME"), "keys"), keystore.StandardScryptN, keystore.StandardScryptP)
	ksw, err := goethereumhelper.NewKeystoreWallet(ks, "0x621A4BAB5f27fCCa37dC2BB8A32B7546Bb0cB9AE")
	if err != nil {
		log.Fatalln("Conta não existente na keystore", err.Error())
	}
	textoASerAssinado := []byte("Jefferson Prestes")
	bytesAssinados, err := ksw.SignTextWithPassphrase(os.Getenv("PALAVRA_PASSE"), textoASerAssinado)
	if err != nil {
		log.Fatalln("Não foi possivel assinar o texto", err.Error())
	}
	fmt.Printf("%s\n%s\n", string(textoASerAssinado), common.Bytes2Hex(bytesAssinados))
}
