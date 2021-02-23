package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jakewitcher/pos-server/graph/generated"
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/datastore"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomerInput) (*model.Customer, error) {
	return datastore.Customers.CreateCustomer(input)
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error) {
	return datastore.Customers.UpdateCustomer(input)
}

func (r *mutationResolver) DeleteCustomer(ctx context.Context, input string) (*model.Customer, error) {
	return datastore.Customers.DeleteCustomer(input)
}

func (r *mutationResolver) CreateStore(ctx context.Context, input model.NewStoreInput) (*model.Store, error) {
	return datastore.Stores.CreateStore(input)
}

func (r *mutationResolver) UpdateStore(ctx context.Context, input model.StoreInput) (*model.Store, error) {
	return datastore.Stores.UpdateStore(input)
}

func (r *mutationResolver) DeleteStore(ctx context.Context, input string) (*model.Store, error) {
	return datastore.Stores.DeleteStore(input)
}

func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.NewEmployeeInput) (*model.Employee, error) {
	return datastore.Employees.CreateEmployee(input)
}

func (r *mutationResolver) UpdateEmployee(ctx context.Context, input model.EmployeeInput) (*model.Employee, error) {
	return datastore.Employees.UpdateEmployee(input)
}

func (r *mutationResolver) DeleteEmployee(ctx context.Context, input string) (*model.Employee, error) {
	return datastore.Employees.DeleteEmployee(input)
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrderInput) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, input string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateManufacturer(ctx context.Context, input model.NewManufacturerInput) (*model.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateManufacturer(ctx context.Context, input model.ManufacturerInput) (*model.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteManufacturer(ctx context.Context, input string) (*model.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateInventoryItem(ctx context.Context, input model.NewInventoryItemInput) (*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateInventoryItem(ctx context.Context, input model.InventoryItemInput) (*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteInventoryItem(ctx context.Context, input string) (*model.InventoryItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customer(ctx context.Context, input string) (*model.Customer, error) {
	return datastore.Customers.FindCustomerById(input)
}

func (r *queryResolver) Customers(ctx context.Context, input *model.CustomerFilter) ([]*model.Customer, error) {
	return datastore.Customers.FindCustomers(input)
}

func (r *queryResolver) Store(ctx context.Context, input string) (*model.Store, error) {
	return datastore.Stores.FindStoreById(input)
}

func (r *queryResolver) Stores(ctx context.Context, input *model.StoreFilter) ([]*model.Store, error) {
	return datastore.Stores.FindStores(input)
}

func (r *queryResolver) Employee(ctx context.Context, input string) (*model.Employee, error) {
	return datastore.Employees.FindEmployeeById(input)
}

func (r *queryResolver) Employees(ctx context.Context) ([]*model.Employee, error) {
	return datastore.Employees.FindEmployees()
}

func (r *queryResolver) Order(ctx context.Context, input string) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
