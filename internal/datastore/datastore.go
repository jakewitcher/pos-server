package datastore

import (
	"github.com/jakewitcher/pos-server/graph/model"
)

var (
	Customers CustomerProvider
	Stores    StoreProvider
	Employees EmployeeProvider
	Users     UserProvider
)

type CustomerProvider interface {
	CreateCustomer(newCustomer model.NewCustomerInput) (*model.Customer, error)
	UpdateCustomer(updatedCustomer model.CustomerInput) (*model.Customer, error)
	DeleteCustomer(customerId string) (*model.Customer, error)
	FindCustomerById(customerId string) (*model.Customer, error)
	FindCustomers(filter *model.CustomerFilter) ([]*model.Customer, error)
}

type StoreProvider interface {
	CreateStore(newStore model.NewStoreInput) (*model.Store, error)
	UpdateStore(updatedStore model.StoreInput) (*model.Store, error)
	DeleteStore(storeId string) (*model.Store, error)
	FindStoreById(storeId string) (*model.Store, error)
	FindStores(filter *model.StoreFilter) ([]*model.Store, error)
}

type EmployeeProvider interface {
	CreateEmployee(newEmployee model.NewEmployeeInput) (*model.Employee, error)
	UpdateEmployee(updatedEmployee model.EmployeeInput) (*model.Employee, error)
	DeleteEmployee(employeeId string) (*model.Employee, error)
	FindEmployeeById(employeeId string) (*model.Employee, error)
	FindEmployees() ([]*model.Employee, error)
}

type UserProvider interface {
	CreateUser(newUser model.NewUserInput) (*model.User, error)
}
