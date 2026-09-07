package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

var (
	ErrNegativePrice = errors.New("trade has a negative price")
	ItemDoesntExist  = errors.New("Item Doesnt Exist")
	CarNotAvailable  = errors.New("CarNotAvailable")
)

type Car struct {
	ID                  int
	Brand               string
	Model               string
	Year                string
	Price               int
	status              string
	times_status_edited int
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

// DB related
type PaymentMethod struct {
	ID     int
	Method string
	Rate   int
}

type ShowroomRepository interface {
	AddCar(ID int, Brand string, Model string, Year string, Price int, Status string) error
	AddMethod(ID int, Method string, Rate int) error

	FindOneCar(ID int) (Car, error)
	UpdateCarStatus(ID int, new_status string, status_update_id int) error
	FindOneMethod(ID int) (PaymentMethod, error)
}

type Service struct {
	mu   sync.Mutex
	repo ShowroomRepository
}

func NewService(repo ShowroomRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddCar(ID int, Brand string, Model string, Year string, Price int, Status string) error {
	s.repo.AddCar(ID, Brand, Model, Year, Price, Status)
	return nil
}

func (s *Service) AddMethod(ID int, Method string, Rate int) error {
	s.repo.AddMethod(ID, Method, Rate)
	return nil
}

func (s *Service) QuotePrice(car_id int, pay_method_id int) (int, error) {
	car, err := s.repo.FindOneCar(car_id)
	if err != nil {
		return 0, fmt.Errorf("car price: %w", err)
	}
	pay_method, err := s.repo.FindOneMethod(pay_method_id)
	if err != nil {
		return 0, fmt.Errorf("payment rate: %w", err)
	}
	return car.GetPriceDynamic(pay_method.Rate), nil
}

func (s *Service) FindCar(car_id int) (Car, error) {
	car, err := s.repo.FindOneCar(car_id)
	if err != nil {
		return Car{}, fmt.Errorf("car price: %w", err)
	}
	return car, nil
}

func (s *Service) FindCarStatus(car_id int) (string, error) {
	car, err := s.repo.FindOneCar(car_id)
	if err != nil {
		return "", fmt.Errorf("car price: %w", err)
	}
	return car.status, nil
}

func (s *Service) UpdateCarStatus(car_id int, new_status string, status_update_id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.repo.UpdateCarStatus(car_id, new_status, status_update_id)
	if err != nil {
		return fmt.Errorf("car price: %w", err)
	}
	return nil
}

type ShowroomDatabase struct {
	table_car            map[int]Car
	table_payment_method map[int]PaymentMethod
}

func (sd ShowroomDatabase) AddCar(ID int, Brand string, Model string, Year string, Price int, Status string) error {
	sd.table_car[ID] = Car{ID, Brand, Model, Year, Price, Status, 0}

	return nil
}

func (sd ShowroomDatabase) AddMethod(ID int, Method string, Rate int) error {
	sd.table_payment_method[ID] = PaymentMethod{ID, Method, Rate}

	return nil
}

func (sd ShowroomDatabase) FindOneCar(ID int) (Car, error) {
	result, ok := sd.table_car[ID]
	if !ok {
		return Car{}, ItemDoesntExist
	}

	return result, nil
}

func (sd ShowroomDatabase) UpdateCarStatus(ID int, new_status string, status_update_id int) error {
	result, ok := sd.table_car[ID]
	if !ok {
		return ItemDoesntExist
	}
	if result.status != "Available" {
		return CarNotAvailable
	}
	result.status = new_status
	result.times_status_edited = result.times_status_edited + 1
	fmt.Println("Car status updated ", status_update_id, " to ", result.status)

	sd.table_car[ID] = result
	return nil
}

func (sd ShowroomDatabase) FindOneMethod(ID int) (PaymentMethod, error) {
	result, ok := sd.table_payment_method[ID]
	if !ok {
		return PaymentMethod{}, ItemDoesntExist
	}

	return result, nil
}

func main() {
	// c := Car{ID: 1, Brand: "", Model: "", Year: "", Price: 200000000}

	// fmt.Println("Cash ", c.GetPriceCash())
	// fmt.Println("Credit ", c.GetPriceCredit())
	// fmt.Println("Cash Dynamic ", c.GetPriceDynamic(PriceRates[TransTypeCash]))
	// fmt.Println("Credit Dynamic ", c.GetPriceDynamic(PriceRates[TransTypeCredit]))
	// fmt.Println("Leasing Dynamic ", c.GetPriceDynamic(PriceRates[TransTypeLeasing]))

	// tit := TradeInTransaction{Car: c, TradeInPrice: 50000000, TransactionType: TransTypeCash}
	// tit_total, err := tit.GetTotalPrice()
	// if err != nil {
	// 	fmt.Println("Error tit")
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println("Trade in Total ", tit_total)
	UseFakeDb()
}

func UseFakeDb() {
	srdb := ShowroomDatabase{
		table_car:            map[int]Car{},
		table_payment_method: map[int]PaymentMethod{},
	}

	service := Service{
		repo: srdb,
	}

	service.AddCar(1, "Brand", "Model", "Year", 200000000, "Available")
	service.AddMethod(1, "Method", 5)
	price, err := service.QuotePrice(1, 1)
	if err != nil {
		fmt.Println("Error tit")
		log.Fatal(err)
		return
	}

	fmt.Println("Price ", price)

	status, err := service.FindCarStatus(1)
	if err != nil {
		fmt.Println("Error tit")
		log.Fatal(err)
		return
	}
	fmt.Println("Status ", status)

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
	status, err = service.FindCarStatus(1)
	if err != nil {
		fmt.Println("Error tit")
		log.Fatal(err)
		return
	}
	fmt.Println("Status ", status)

}
