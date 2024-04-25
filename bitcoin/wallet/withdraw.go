package wallet

import "fmt"

func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if amount > w.balance || amount <= 0 {
		return fmt.Errorf("%w", ErrInsufficientFunds)
	}

	w.balance -= amount
	return nil
}
