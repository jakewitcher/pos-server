package inventory

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/manufacturers"
	"github.com/jakewitcher/pos-server/pkg/currency"
	"strconv"
)

type ItemEntity struct {
	Id             int    `json:"id"`
	Description    string `json:"description"`
	Cost           int    `json:"cost"`
	Retail         int    `json:"retail"`
	ManufacturerId int    `json:"manufacturer_id"`
}

type Inventory struct {
	StoreId  int `json:"store_id"`
	ItemId   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

func (i *ItemEntity) ToDTO(manufacturer *manufacturers.ManufacturerEntity) *model.InventoryItem {
	cost := currency.CreateFromCents(i.Cost)
	retail := currency.CreateFromCents(i.Retail)

	return &model.InventoryItem{
		ID:           strconv.Itoa(i.Id),
		Description:  i.Description,
		Cost:         cost.AsDollars(),
		Retail:       retail.AsDollars(),
		Manufacturer: manufacturer.ToDTO(),
	}
}
