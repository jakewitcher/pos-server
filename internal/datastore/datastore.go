package datastore

import (
	"github.com/jakewitcher/pos-server/graph/model"
)

var Customers CustomerProvider

type CustomerProvider interface {
	FindCustomerById(id string) *model.Customer
	GetAllCustomers() []*model.Customer
}
