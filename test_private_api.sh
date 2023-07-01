#!/bin/bash

## Run All Tests
# godotenv -f ./.env go test -cover ./...

## Run Specific Test
# godotenv -f ./.env go test -cover ./... ./client/account_test
godotenv -f ./.env go test -run TestGetAccounts ./client/account_test.go -v