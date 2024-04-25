package wallet

import (
	"errors"
	"math"
	"sync"
)

var ErrInsufficientFunds = errors.New("cannot withdraw: insufficient funds")
var ErrInvalidAmount = errors.New("amount must be greater than 0")
var ErrMaxAmount = errors.New("amount must be less than")

const maxFloat = math.MaxFloat64

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
	// use RWMutex because only read and write operations happened
	// and no need to block read operations
	mu sync.RWMutex
}

// New creates a new Wallet with a balance of 0
func New() *Wallet {
	return &Wallet{
		balance: 0,
	}
}
