package employees

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"strconv"
)

type Role string

const (
	Manager        Role = "MANAGER"
	SalesAssociate Role = "SALES_ASSOCIATE"
)

type EmployeeEntity struct {
	Id        int64  `json:"id"`
	StoreId   int64  `json:"store_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      Role   `json:"role"`
}

func (m *EmployeeEntity) ToDTO() *model.Employee {
	return &model.Employee{
		ID:        strconv.FormatInt(m.Id, 10),
		StoreID:   strconv.FormatInt(m.StoreId, 10),
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Role:      model.Roles(m.Role),
	}
}
