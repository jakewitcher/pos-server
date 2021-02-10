package datastore

import (
	"github.com/jakewitcher/pos-server/graph/model"
)

var Customers CustomerProvider

type CustomerProvider interface {
	CreateCustomer(newCustomer model.NewCustomerInput) *model.Customer
	UpdateCustomer(updatedCustomer model.CustomerInput) *model.Customer
	DeleteCustomer(customerId string) *model.Customer
	FindCustomerById(customerId string) *model.Customer
	GetAllCustomers() []*model.Customer
}
