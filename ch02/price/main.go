package price

type Season int

const (
	Peak Season = iota + 1
	Normal
	Off
)

func (s Season) Price(price int) int {
	if s == Peak {
		return price + 200
	}

	return price
}
