package main

import (
	"sync"
	"testing"
)

// func TestCar(t *testing.T) {
// 	cases := []struct {
// 		name                 string
// 		car                  Car
// 		desired_price_cash   int
// 		desired_price_credit int
// 	}{
// 		{"Base test", Car{ID: 1, Brand: "XXX", Model: "XXX", Year: "XXX", Price: 200000000}, 200000000, 204000000},
// 	}

// 	for _, cas := range cases {
// 		t.Run(cas.name, func(t *testing.T) {
// 			cash_result := cas.car.GetPriceCash()
// 			if cash_result != cas.desired_price_cash {
// 				t.Errorf("cash: got %d, want %d", cash_result, cas.desired_price_cash)
// 			}

// 			credit_result := cas.car.GetPriceCredit()
// 			if credit_result != cas.desired_price_credit {
// 				t.Errorf("credit: got %d, want %d", credit_result, cas.desired_price_credit)
// 			}
// 		})
// 	}
// }

// func TestTradeIn(t *testing.T) {
// 	c := Car{ID: 1, Brand: "", Model: "", Year: "", Price: 200000000}
// 	tit_cash := TradeInTransaction{Car: c, TradeInPrice: 50000000, TransactionType: TransTypeCash}
// 	tit_credit := TradeInTransaction{Car: c, TradeInPrice: 50000000, TransactionType: TransTypeCredit}
// 	tit_cash_error := TradeInTransaction{Car: c, TradeInPrice: 500000000000, TransactionType: TransTypeCash}
// 	tit_credit_error := TradeInTransaction{Car: c, TradeInPrice: 500000000000, TransactionType: TransTypeCredit}

// 	cases := []struct {
// 		name               string
// 		tit                TradeInTransaction
// 		desire_price_error bool
// 	}{
// 		{"Correct cash", tit_cash, false},
// 		{"Correct credit", tit_credit, false},
// 		{"Testing Error cash", tit_cash_error, true},
// 		{"Testing Error credit", tit_credit_error, true},
// 	}

// 	for _, cas := range cases {
// 		t.Run(cas.name, func(t *testing.T) {
// 			total_result, err := cas.tit.GetTotalPrice()
// 			if err != nil && !cas.desire_price_error {
// 				t.Error(err)
// 			} else if cas.desire_price_error {
// 				if ErrNegativePrice != err {
// 					t.Errorf("Error untriggered ErrNegativePrice")
// 				}
// 			}
// 			fmt.Println(cas.name+" succeed got ", total_result)
// 		})
// 	}
// }

// func TestCarPriceDynamic(t *testing.T) {
// 	cases := []struct {
// 		name                  string
// 		car                   Car
// 		desired_dynamic_price int
// 		trans_type            TransactionType
// 	}{
// 		{"Base test cash", Car{ID: 1, Brand: "XXX", Model: "XXX", Year: "XXX", Price: 200000000}, 200000000, TransTypeCash},
// 		{"Base test credit", Car{ID: 1, Brand: "XXX", Model: "XXX", Year: "XXX", Price: 200000000}, 204000000, TransTypeCredit},
// 		{"Base test dynamic", Car{ID: 1, Brand: "XXX", Model: "XXX", Year: "XXX", Price: 200000000}, 210000000, TransTypeLeasing},
// 	}

// 	for _, cas := range cases {
// 		t.Run(cas.name, func(t *testing.T) {
// 			result := cas.car.GetPriceDynamic(PriceRates[cas.trans_type])
// 			if result != cas.desired_dynamic_price {
// 				t.Errorf("cash: got %d, want %d", result, cas.desired_dynamic_price)
// 			}
// 		})
// 	}
// }

// func TestFakeDB(t *testing.T) {

// 	cases := []struct {
// 		name                  string
// 		desired_dynamic_price int
// 		trans_type            TransactionType
// 	}{
// 		{"Base test cash", 200000000, TransTypeCash},
// 		{"Base test credit", 204000000, TransTypeCredit},
// 		{"Base test dynamic", 210000000, TransTypeLeasing},
// 	}

// 	for _, cas := range cases {
// 		t.Run(cas.name, func(t *testing.T) {
// 			srdb := ShowroomDatabase{
// 				table_car:            map[int]Car{},
// 				table_payment_method: map[int]PaymentMethod{},
// 			}

// 			service := Service{
// 				repo: srdb,
// 			}
// 			service.AddCar(1, "Brand", "Model", "Year", 200000000, "Available")
// 			service.AddMethod(1, "Method", PriceRates[cas.trans_type])
// 			price, err := service.QuotePrice(1, 1)
// 			if err != nil {
// 				t.Error(err)
// 			}
// 			if price != cas.desired_dynamic_price {
// 				t.Errorf("cash: got %d, want %d", price, cas.desired_dynamic_price)
// 			}
// 		})
// 	}
// }

func TestMutex(t *testing.T) {

	srdb := ShowroomDatabase{
		table_car:            map[int]Car{},
		table_payment_method: map[int]PaymentMethod{},
	}

	service := Service{
		repo: srdb,
	}

	service.AddCar(1, "Brand", "Model", "Year", 200000000, "Available")

	var wg sync.WaitGroup

	wg.Go(func() {
		for i := range 10000 {
			service.UpdateCarStatus(1, "Sold", i)
		}
	})
	wg.Go(func() {
		for i := range 10000 {
			service.UpdateCarStatus(1, "Reserved", i)
		}
	})
	wg.Go(func() {
		for i := range 10000 {
			service.UpdateCarStatus(1, "Pulled", i)
		}
	})
	wg.Wait()

	result_car, err := service.FindCar(1)
	if err != nil {
		t.Error(err)
	}
	if result_car.times_status_edited != 1 {
		t.Errorf("Edited more than one time")
	}
}
