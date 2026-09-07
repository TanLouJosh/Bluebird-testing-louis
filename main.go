package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNegativePrice = errors.New("trade has a negative price")
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

type TransactionType int

const (
	TransTypeCash TransactionType = iota
	TransTypeCredit
)

type TradeInTransaction struct {
	Car             Car
	TradeInPrice    int
	TransactionType TransactionType
}

func (tit TradeInTransaction) GetTotalPrice() (int, error) {
	price := 0
	if tit.TransactionType == TransTypeCash {
		price = tit.Car.GetPriceCash() - tit.TradeInPrice
	} else {
		price = tit.Car.GetPriceCredit() - tit.TradeInPrice
	}

	if price < 0 {
		return price, ErrNegativePrice
	}

	return price, nil
}

func main() {
	c := Car{ID: 1, Brand: "", Model: "", Year: "", Price: 200000000}

	fmt.Println("Cash ", c.GetPriceCash())
	fmt.Println("Credit ", c.GetPriceCredit())

	tit := TradeInTransaction{Car: c, TradeInPrice: 50000000, TransactionType: TransTypeCash}
	tit_total, err := tit.GetTotalPrice()
	if err != nil {
		fmt.Println("Error tit")
		log.Fatal(err)
		return
	}
	fmt.Println("Trade in Total ", tit_total)
}
