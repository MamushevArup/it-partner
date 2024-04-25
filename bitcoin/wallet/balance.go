package wallet

func (w *Wallet) Balance() Bitcoin {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.balance
}
