package dollar

type Dollar struct {
	Price int
}

func (d *Dollar) Multiple(dollar *Dollar) *Dollar {
	return &Dollar{
		Price: d.Price * dollar.Price,
	}
}

func (d *Dollar) IsEqualPrice(dollar *Dollar) bool {
	if d.Price == dollar.Price {
		return true
	}
	return false
}
