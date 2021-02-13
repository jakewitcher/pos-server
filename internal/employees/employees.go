package employees

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"strconv"
)

type ManagerEntity struct {
	Id        int64  `json:"id"`
	StoreId   int64  `json:"store_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type SalesAssociateEntity struct {
	Id        int64  `json:"id"`
	StoreId   int64  `json:"store_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (m *ManagerEntity) ToDTO() *model.Manager {
	return &model.Manager{
		ID:        strconv.FormatInt(m.Id, 10),
		StoreID:   strconv.FormatInt(m.StoreId, 10),
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}
}

func (s *SalesAssociateEntity) ToDTO() *model.SalesAssociate {
	return &model.SalesAssociate{
		ID:        strconv.FormatInt(s.Id, 10),
		StoreID:   strconv.FormatInt(s.StoreId, 10),
		FirstName: s.FirstName,
		LastName:  s.LastName,
	}
}
