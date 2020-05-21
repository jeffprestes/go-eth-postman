#!/bin/bash
go mod tidy
go build -o ./build/bin/walletcreate ./cmd/walletcreate