#!/bin/bash
go mod tidy
go build -o ./build/bin/walletopen ./cmd/walletopen