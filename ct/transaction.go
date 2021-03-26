package ct

type TransactionType int

const (
	Credit TransactionType = iota + 1
	Debit
)

type Restriction int

const (
	AllowRead Restriction = iota
	AllowCredit
	AllowDebit
)

type Transaction interface {
	Account() string
	Number() uint
	Amount() uint
	Type() TransactionType
}

func DebitTx(acc string, number uint, amount uint) Transaction {
	return &debit{acc, number, amount}
}

type debit struct {
	acc    string
	tx     uint
	amount uint
}

func (d *debit) Account() string {
	return d.acc
}

func (d *debit) Number() uint {
	return d.tx
}

func (d *debit) Amount() uint {
	return d.amount
}

func (d *debit) Type() TransactionType {
	return Debit
}

func CreditTx(acc string, number uint, amount uint) Transaction {
	return &credit{acc, number, amount}
}

type credit struct {
	acc    string
	tx     uint
	amount uint
}

func (c *credit) Account() string {
	return c.acc
}

func (c *credit) Number() uint {
	return c.tx
}

func (c *credit) Amount() uint {
	return c.amount
}

func (c *credit) Type() TransactionType {
	return Credit
}
