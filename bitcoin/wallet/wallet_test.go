package wallet

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	tests := map[string]struct {
		depositAmount         Bitcoin
		withdrawAmount        Bitcoin
		expectedDepositError  error
		expectedWithdrawError error
		expectedBalance       Bitcoin
	}{
		"deposit positive amount": {
			depositAmount:        10,
			expectedDepositError: nil,
			expectedBalance:      10,
		},
		"deposit negative amount": {
			depositAmount:        -10,
			expectedDepositError: ErrInvalidAmount,
			expectedBalance:      0,
		},
		"deposit zero amount": {
			depositAmount:        0,
			expectedDepositError: ErrInvalidAmount,
			expectedBalance:      0,
		},
		"withdraw positive amount": {
			depositAmount:         10,
			withdrawAmount:        5,
			expectedWithdrawError: nil,
			expectedBalance:       5,
		},
		"withdraw negative amount": {
			depositAmount:         10,
			withdrawAmount:        -5,
			expectedWithdrawError: ErrInsufficientFunds,
			expectedBalance:       10,
		},
		"withdraw zero amount": {
			depositAmount:         10,
			withdrawAmount:        0,
			expectedWithdrawError: ErrInsufficientFunds,
			expectedBalance:       10,
		},
		"withdraw more than balance": {
			depositAmount:         10,
			withdrawAmount:        15,
			expectedWithdrawError: ErrInsufficientFunds,
			expectedBalance:       10,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			w := New()

			if test.depositAmount != 0 {
				err := w.Deposit(test.depositAmount)
				if !errors.Is(err, test.expectedDepositError) {
					t.Errorf("Deposit error: got %v, want %v", err, test.expectedDepositError)
				}
			}

			if test.withdrawAmount != 0 {
				err := w.Withdraw(test.withdrawAmount)
				if !errors.Is(err, test.expectedWithdrawError) {
					t.Errorf("Withdraw error: got %v, want %v", err, test.expectedWithdrawError)
				}
			}

			got := w.Balance()
			if got != test.expectedBalance {
				t.Errorf("Balance: got %v, want %v", got, test.expectedBalance)
			}
		})
	}
}
