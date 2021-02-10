package datastore

import (
	"github.com/jakewitcher/pos-server/graph/model"
)

var (
	Customers      CustomerProvider
	Stores         StoreProvider
	StoreLocations StoreLocationProvider
)

type CustomerProvider interface {
	CreateCustomer(newCustomer model.NewCustomerInput) *model.Customer
	UpdateCustomer(updatedCustomer model.CustomerInput) *model.Customer
	DeleteCustomer(customerId string) *model.Customer
	FindCustomerById(customerId string) *model.Customer
	FindAllCustomers() []*model.Customer
}

type StoreProvider interface {
	CreateStore(newStore model.NewStoreInput) *model.Store
	UpdateStore(updatedStore model.StoreInput) *model.Store
	DeleteStore(storeId string) *model.Store
	FindStoreById(storeId string) *model.Store
	FindAllStores() []*model.Store
}

type StoreLocationProvider interface {
	CreateStoreLocation(newStoreLocation model.NewStoreLocationInput) *model.StoreLocation
	UpdateStoreLocation(updatedStoreLocation model.StoreLocationInput) *model.StoreLocation
	DeleteStoreLocation(storeLocationId string) *model.StoreLocation
	FindStoreLocationById(storeLocationId string) *model.StoreLocation
	FindAllStoreLocations() []*model.StoreLocation
}
