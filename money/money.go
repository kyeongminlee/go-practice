package money

type MoneyInterface interface {
	Multiple(money interface{}) MoneyInterface
	IsEqualPrice(money interface{}) bool
}

var (
	DollarString = "dollar"
	FrancString  = "franc"
)

type Money struct {
	Price    int
	Currency string
}

func Dollar(price int) *Money {
	return &Money{
		Price:    price,
		Currency: DollarString,
	}
}

func Franc(price int) *Money {
	return &Money{
		Price:    price,
		Currency: FrancString,
	}
}

func (m Money) Add(money Money) *Money {
	if m.Currency == money.Currency && m.Currency != "" {
		return &Money{
			Price:    m.Price + money.Price,
			Currency: m.Currency,
		}
	}

	return nil
}

func (m Money) IsEqualPrice(money Money) bool {
	if m.Currency == money.Currency && m.Currency != "" {
		return m.Price == money.Price
	}

	return false
}
