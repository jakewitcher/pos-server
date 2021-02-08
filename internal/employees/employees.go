package employees

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"strconv"
)

type ManagerEntity struct {
	Id        int    `json:"id"`
	StoreId   int    `json:"store_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type SalesAssociateEntity struct {
	Id        int    `json:"id"`
	StoreId   int    `json:"store_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (m *ManagerEntity) ToDTO() *model.Manager {
	return &model.Manager{
		ID:        strconv.Itoa(m.Id),
		StoreID:   strconv.Itoa(m.StoreId),
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
}

func (s *SalesAssociateEntity) ToDTO() *model.SalesAssociate {
	return &model.SalesAssociate{
		ID:        strconv.Itoa(s.Id),
		StoreID:   strconv.Itoa(s.StoreId),
		FirstName: s.FirstName,
		LastName:  s.LastName,
	}
}
