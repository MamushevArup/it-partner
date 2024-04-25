package wallet

import "fmt"

func (w *Wallet) Deposit(amount Bitcoin) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if amount <= 0 {
		return fmt.Errorf("%w", ErrInvalidAmount)
	}

	if amount > maxFloat {
		return fmt.Errorf("%w %f", ErrMaxAmount, maxFloat)
	}

	w.balance += amount
	return nil
}
