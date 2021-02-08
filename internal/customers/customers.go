package customers

import (
	"github.com/jakewitcher/pos-server/graph/model"
	"strconv"
)

type CustomerEntity struct {
	Id            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	ContactInfoId int    `json:"contact_info_id"`
}

type ContactInfoEntity struct {
	Id           int    `json:"id"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
}

func (c *CustomerEntity) ToDTO(contactInfo *ContactInfoEntity) *model.Customer {
	return &model.Customer{
		ID:          strconv.Itoa(c.Id),
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		ContactInfo: contactInfo.ToDTO(),
	}
}

func (c *ContactInfoEntity) ToDTO() *model.ContactInfo {
	return &model.ContactInfo{
		PhoneNumber:  c.PhoneNumber,
		EmailAddress: c.EmailAddress,
	}
}
