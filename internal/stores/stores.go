package stores

import (
	"github.com/jakewitcher/pos-server/graph/model"
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
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	LocationId int64  `json:"location_id"`
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

func (s *StoreEntity) ToDTO(location *StoreLocationEntity) *model.Store {
	return &model.Store{
		ID:       strconv.FormatInt(s.Id, 10),
		Name:     s.Name,
		Location: location.ToDTO(),
	}
}
