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

func (c Car) GetPriceDynamic(rate int) int {
	return (c.Price * (rate + 100)) / 100
}

type TransactionType int

const (
	TransTypeCash TransactionType = iota
	TransTypeCredit
	TransTypeLeasing
)

var PriceRates = map[TransactionType]int{
	TransTypeCash:    0,
	TransTypeCredit:  2,
	TransTypeLeasing: 5,
}

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
	fmt.Println("Cash Dynamic ", c.GetPriceDynamic(PriceRates[TransTypeCash]))
	fmt.Println("Credit Dynamic ", c.GetPriceDynamic(PriceRates[TransTypeCredit]))
	fmt.Println("Leasing Dynamic ", c.GetPriceDynamic(PriceRates[TransTypeLeasing]))

	tit := TradeInTransaction{Car: c, TradeInPrice: 50000000, TransactionType: TransTypeCash}
	tit_total, err := tit.GetTotalPrice()
	if err != nil {
		fmt.Println("Error tit")
		log.Fatal(err)
		return
	}
	fmt.Println("Trade in Total ", tit_total)
}
