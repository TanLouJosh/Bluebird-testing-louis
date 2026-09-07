package main

import (
	"testing"
)

func TestCar(t *testing.T) {
	cases := []struct {
		name                 string
		car                  Car
		desired_price_cash   int
		desired_price_credit int
	}{
		{"Base test", Car{ID: 1, Brand: "XXX", Model: "XXX", Year: "XXX", Price: 200000000}, 200000000, 204000000},
	}

	for _, cas := range cases {
		t.Run(cas.name, func(t *testing.T) {
			cash_result := cas.car.GetPriceCash()
			if cash_result != cas.desired_price_cash {
				t.Errorf("cash: got %d, want %d", cash_result, cas.desired_price_cash)
			}

			credit_result := cas.car.GetPriceCredit()
			if credit_result != cas.desired_price_credit {
				t.Errorf("credit: got %d, want %d", credit_result, cas.desired_price_credit)
			}
		})
	}
}
