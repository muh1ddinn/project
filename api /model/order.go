package model

type GetOrder struct {
	Id string `json:"id"`
	//Car        Car       `json:"car"`
	//Customers  Customers `json:"cudtomer"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
	Status     string `json:"status"`
	Paid       bool   `json:"payment_status"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateOrder struct {
	Id         string `json:"id"`
	CarId      string `json:"car_id"`
	CustomerId string `json:"customer_id"`
	FromDate   string `json:"from_date"`
	ToDate     string `json:"to_date"`
	Status     string `json:"status"`
	Paid       bool   `json:"payment_status"`
	Amount     int    `json:"amount"`
}

type GetAllOrder struct {
	Orders []GetOrder `json:"orders"`
	Count  int64      `json:"count of orders"`
}
