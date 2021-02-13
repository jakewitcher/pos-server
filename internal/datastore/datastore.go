package datastore

import (
	"github.com/jakewitcher/pos-server/graph/model"
)

var (
	Customers CustomerProvider
	Stores    StoreProvider
	Employees EmployeeProvider
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
	CreateManager(newManager model.NewManagerInput) (*model.Manager, error)
	UpdateManager(updatedManager model.ManagerInput) (*model.Manager, error)
	DeleteManager(managerId string) (*model.Manager, error)
	FindManagerById(managerId string) (*model.Manager, error)
}
