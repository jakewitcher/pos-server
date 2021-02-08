package orders

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/customers"
	"github.com/jakewitcher/pos-server/internal/employees"
	"github.com/jakewitcher/pos-server/pkg/currency"
	"strconv"
)

type OrderEntity struct {
	Id               int `json:"id"`
	CustomerId       int `json:"customer_id"`
	StoreId          int `json:"store_id"`
	SalesAssociateId int `json:"sales_associate_id"`
}

type LineItemEntity struct {
	OrderId     int    `json:"order_id"`
	Description string `json:"description"`
	Retail      int    `json:"retail"`
	Quantity    int    `json:"quantity"`
}

func (o *OrderEntity) ToDTO(
	storeId int,
	customer *customers.CustomerEntity,
	customerContactInfo *customers.ContactInfoEntity,
	salesAssociate *employees.SalesAssociateEntity,
	lineItems []*LineItemEntity,
) *model.Order {

	lineItemDTOs := make([]*model.LineItem, len(lineItems))
	for i, lineItem := range lineItems {
		lineItemDTOs[i] = lineItem.ToDTO()
	}

	return &model.Order{
		ID:             strconv.Itoa(o.Id),
		StoreID:        strconv.Itoa(storeId),
		Customer:       customer.ToDTO(customerContactInfo),
		SalesAssociate: salesAssociate.ToDTO(),
		LineItems:      lineItemDTOs,
	}
}

func (l *LineItemEntity) ToDTO() *model.LineItem {
	retail := currency.CreateFromCents(l.Retail)

	return &model.LineItem{
		Description: l.Description,
		Retail:      retail.AsDollars(),
		Quantity:    l.Quantity,
	}
}
