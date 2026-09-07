package main

import (
	"fmt"
)

type Car struct {
	ID    int
	Brand string
	Model string
	Year  string
	Price int
}

func (c Car) GetPriceCash() int {
	return c.Price
}

func (c Car) GetPriceCredit() int {
	return (c.Price * 102) / 100
}

func main() {
	c := Car{ID: 1, Brand: "", Model: "", Year: "", Price: 200000000}

	fmt.Println("Cash ", c.GetPriceCash())
	fmt.Println("Credit ", c.GetPriceCredit())
}
