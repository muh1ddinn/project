package storage

import model "cars_with_sql/models"

type IStorage interface {
	CloseDB()
	Car() ICarstorage
	Customer() ICustomerStorage
}

type ICarstorage interface {
	GetAllCars(request model.GetAllCarsRequest) (model.GetAllCarsResponse, error)
	UpdateCar(car model.Car) (string, error)
	Deletecar(string) error
	Createcar(model.Car) (string, error)
	GetByIDCar(string) ([]model.Car, error)
}

type ICustomerStorage interface {
	Create(customer model.Customers) (string, error)
	GetAllCustomers(request model.GetAllCustomerRequest) (model.GetAllCustomersResponse, error)
	UpdateCustomer(customer model.Customers) (string, error)
	DeleteCustomer(string) error
	GetByIDCustomer(string) ([]model.Customers, error)
}
