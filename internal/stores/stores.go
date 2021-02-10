package stores

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/employees"
	"github.com/jakewitcher/pos-server/internal/inventory"
	"github.com/jakewitcher/pos-server/internal/manufacturers"
	"strconv"
)

type StoreLocationEntity struct {
	Id      int64  `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type StoreEntity struct {
	Id         int64 `json:"id"`
	LocationId int   `json:"location_id"`
}

func (l *StoreLocationEntity) ToDTO() *model.StoreLocation {
	return &model.StoreLocation{
		ID:      strconv.FormatInt(l.Id, 10),
		Street:  l.Street,
		City:    l.City,
		State:   l.State,
		ZipCode: l.ZipCode,
	}
}

func (s *StoreEntity) ToDTO(
	location *StoreLocationEntity,
	manager *employees.ManagerEntity,
	salesAssociates []*employees.SalesAssociateEntity,
	inventory []*inventory.ItemEntity,
	manufacturers map[int]*manufacturers.ManufacturerEntity,
) *model.Store {

	salesAssociateDTOs := make([]*model.SalesAssociate, len(salesAssociates))
	for i, salesAssociate := range salesAssociates {
		salesAssociateDTOs[i] = salesAssociate.ToDTO()
	}

	inventoryDTO := make([]*model.InventoryItem, len(inventory))
	for i, item := range inventory {
		manufacturer := manufacturers[item.ManufacturerId]
		inventoryDTO[i] = item.ToDTO(manufacturer)
	}

	return &model.Store{
		ID:              strconv.FormatInt(s.Id, 10),
		Location:        location.ToDTO(),
		Manager:         manager.ToDTO(),
		SalesAssociates: salesAssociateDTOs,
		Inventory:       inventoryDTO,
	}
}
