BINARY_WALLET_CREATE=walletcreate
BINARY_WALLET_OPEN=walletopen
BINARY_TXCREATE=txcreate
BINARY_TXPOST=txpost

PROJECT_DATA_DIR_NAME=.go-eth-postman
PALAVRA_PASSE=Teste1020
TF_CONTRACT_ADDRESS=0x1F63fbf4E00824dC00481C2a3f8bD5Cc9d055762
ETH_CLIENT_URL=https://rinkeby.infura.io/v3/1eae79e9335242369fe9b17b9413d721
SENDER_ADDRESS=0x621A4BAB5f27fCCa37dC2BB8A32B7546Bb0cB9AE

.PHONY: all build buildwalletcreate buildwalletopen clean cleanwalletcreate cleanwalletopen

cleanwalletcreate: 
	rm build/bin/$(BINARY_WALLET_CREATE)

cleanwalletopen: 
	rm build/bin/$(BINARY_WALLET_OPEN)

clean:
	cleanwalletcreate cleanwalletopen

buildwalletcreate:
	scripts/$(BINARY_WALLET_CREATE)/build.sh

buildwalletopen:
	scripts/$(BINARY_WALLET_OPEN)/build.sh

build:
	buildwalletcreate buildwalletopen

runwalletcreate:
	env PALAVRA_PASSE=$(PALAVRA_PASSE) PROJECT_DATA_DIR_NAME=$(PROJECT_DATA_DIR_NAME) go run cmd/$(BINARY_WALLET_CREATE)/*.go

runwalletopen:
	env PALAVRA_PASSE=$(PALAVRA_PASSE) PROJECT_DATA_DIR_NAME=$(PROJECT_DATA_DIR_NAME) go run cmd/$(BINARY_WALLET_OPEN)/*.go

runtxcreate:
	env TF_CONTRACT_ADDRESS=$(TF_CONTRACT_ADDRESS) SENDER_ADDRESS=$(SENDER_ADDRESS) go run cmd/$(BINARY_TXCREATE)/*.go

runtxpost:
	env PALAVRA_PASSE=$(PALAVRA_PASSE) PROJECT_DATA_DIR_NAME=$(PROJECT_DATA_DIR_NAME) TF_CONTRACT_ADDRESS=$(TF_CONTRACT_ADDRESS) ETH_CLIENT_URL=$(ETH_CLIENT_URL) SENDER_ADDRESS=$(SENDER_ADDRESS) go run cmd/$(BINARY_TXPOST)/*.go


runall: runwalletcreate runwalletopen

all: clean build
