package stores

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/employees"
	"github.com/jakewitcher/pos-server/internal/inventory"
	"github.com/jakewitcher/pos-server/internal/manufacturers"
	"strconv"
)

type LocationEntity struct {
	Id      int    `json:"id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type StoreEntity struct {
	Id         int `json:"id"`
	LocationId int `json:"location_id"`
}

func (l *LocationEntity) ToDTO() *model.Location {
	return &model.Location{
		Street:  l.Street,
		City:    l.City,
		State:   l.State,
		Zipcode: l.ZipCode,
	}
}

func (s *StoreEntity) ToDTO(
	location *LocationEntity,
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
		ID:              strconv.Itoa(s.Id),
		Location:        location.ToDTO(),
		Manager:         manager.ToDTO(),
		SalesAssociates: salesAssociateDTOs,
		Inventory:       inventoryDTO,
	}
}
