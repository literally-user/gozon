package bank

type Card struct {
	CardNumber string
	Month      int
	Year       int
	CVV        int
}

type Adapter interface {
	Withdraw(from Card) error
	Refund(to Card) error
}

type AdapterFactory interface {
	GetBankAdapter(bankName string) (Adapter, error)
}
