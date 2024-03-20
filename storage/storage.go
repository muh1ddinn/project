package storage

import "cars_with_sql/api /model"

type IStorage interface {
	CloseDB()
	Car() ICarstorage
	Customer() ICustomerStorage
	Order() Iorderstorage
}

type ICarstorage interface {
	GetAllCars(request model.GetAllCarsRequest) (model.GetAllCarsResponse, error)
	//UpdateCar(car model.Car) (string, error)
	Deletecar(string) error
	Createcar(model.Car) (string, error)
	GetByidcar(string) ([]model.Car, error)
}

type ICustomerStorage interface {
	CreateCus(customer model.Customers) (string, error)
	GetAllCustomers(string) (model.GetAllCustomersResponse, error)
	//UpdateCustomer(customer model.Customers) (string, error)
	DeleteCustomer(string) error
	GetByIDCustomer(string) (model.Customers, error)
}

type Iorderstorage interface {
	Create(order model.CreateOrder) (string, error)
	GetAll(string) (model.GetAllOrder, error)
	GetByID(string) (model.GetOrder, error)
}
