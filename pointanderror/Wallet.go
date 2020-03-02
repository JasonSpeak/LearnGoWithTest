package pointanderror

//Wallet defines a wallet
type Wallet struct {
	blance BitCoin
}

//Deposit :Add money into wallet
func (w *Wallet) Deposit(money BitCoin) {
	w.blance += money
}

//Balance :return blance of the wallet
func (w *Wallet) Balance() BitCoin {
	return w.blance
}

//Withdraw : withdraw money from wallet
func (w *Wallet) Withdraw(money BitCoin) error {
	if money > w.blance {
		return InsufficientFundsError
	}

	w.blance -= money
	return nil
}
