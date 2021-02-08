package manufacturers

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"strconv"
)

type ManufacturerEntity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (m *ManufacturerEntity) ToDTO() *model.Manufacturer {
	return &model.Manufacturer{
		ID:   strconv.Itoa(m.Id),
		Name: m.Name,
	}
}
