package db

type Item struct {
	Name  string
	Price float64
}

func LoadItem(id int) *Item {
	return &Item{
		Name:  "mochi",
		Price: 6.10,
	}
}
