#!/bin/bash

## Run All Tests
# godotenv -f ./.env go test -cover ./...

## Run Specific Test
# godotenv -f ./.env go test -cover ./... ./robinhood/account_test
godotenv -f ./.env go test -run TestGetAccounts ./robinhood/account_test.go -v