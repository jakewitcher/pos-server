package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/jakewitcher/pos-server/internal/datastore"

	"github.com/jakewitcher/pos-server/graph/generated"
	"github.com/jakewitcher/pos-server/graph/model"
)

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.NewCustomerInput) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCustomer(ctx context.Context, input string) (*model.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateManager(ctx context.Context, input model.NewManagerInput) (*model.Manager, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateManager(ctx context.Context, input model.ManagerInput) (*model.Manager, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteManager(ctx context.Context, input string) (*model.Manager, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSalesAssociate(ctx context.Context, input model.NewSalesAssociateInput) (*model.SalesAssociate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSalesAssociate(ctx context.Context, input model.SalesAssociateInput) (*model.SalesAssociate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSalesAssociate(ctx context.Context, input string) (*model.SalesAssociate, error) {
	panic(fmt.Errorf("not implemented"))
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

func (r *mutationResolver) CreateStore(ctx context.Context, input model.NewStoreInput) (*model.Store, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStore(ctx context.Context, input model.StoreInput) (*model.Store, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteStore(ctx context.Context, input string) (*model.Store, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Customer(ctx context.Context, input string) (*model.Customer, error) {
	return datastore.Customers.FindCustomerById(input), nil
}

func (r *queryResolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	return datastore.Customers.GetAllCustomers(), nil
}

func (r *queryResolver) Employee(ctx context.Context, input string) (model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Employees(ctx context.Context) ([]model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Store(ctx context.Context, input string) (*model.Store, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Stores(ctx context.Context) ([]*model.Store, error) {
	panic(fmt.Errorf("not implemented"))
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
