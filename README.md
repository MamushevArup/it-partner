# Installation
To run and test code you need golang installed on you PC
First install go!!!

# Go Modules: Bitcoin Wallet and StringUtil

This repository contains two Go modules: `bitcoin` and `stringutil`.

## Wallet test run
To run test for the wallet, first move to the bitcoin directory
`cd bitcoin`
Then you can run 
`go test ./wallet`

## Bitcoin Usage

The `bitcoin` module provides a concurrent-safe `Wallet` type for managing Bitcoin transactions. It supports depositing and withdrawing bitcoins, and checking the wallet's balance.

```go
w := wallet.New()
err := w.Deposit(10)

err = w.Withdraw(8)

balance := w.Balance()
```

## Reverse and Count Symbol in string test run
To run test for the wallet, first move to the bitcoin directory
`cd stringutil`

Running test related to the reverse string done by this command
`go test ./reverse`

Running test related to the count symbol done by this command
`go test ./count`

Or to run both of them simultaneously
`go test ./...`

## String Util Usage

```go
str := reverse.String("hello")

length := count.Symbol(str)
```


