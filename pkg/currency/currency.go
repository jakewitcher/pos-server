package currency

type USD struct {
	amount int
}

func (u *USD) AsDollars() float64 {
	return float64(u.amount) / 100
}

func (u *USD) AsCents() float64 {
	return float64(u.amount)
}

func CreateFromCents(amount int) USD {
	return USD{
		amount: amount,
	}
}
