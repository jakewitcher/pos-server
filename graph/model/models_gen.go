// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ContactInfo struct {
	PhoneNumber  string `json:"phoneNumber"`
	EmailAddress string `json:"emailAddress"`
}

type ContactInfoInput struct {
	PhoneNumber  string `json:"phoneNumber"`
	EmailAddress string `json:"emailAddress"`
}

type Customer struct {
	ID          string       `json:"id"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	ContactInfo *ContactInfo `json:"contactInfo"`
}

type CustomerFilter struct {
	LastName     *string `json:"lastName"`
	PhoneNumber  *string `json:"phoneNumber"`
	EmailAddress *string `json:"emailAddress"`
}

type CustomerInput struct {
	ID          string            `json:"id"`
	FirstName   string            `json:"firstName"`
	LastName    string            `json:"lastName"`
	ContactInfo *ContactInfoInput `json:"contactInfo"`
}

type Employee struct {
	ID        string `json:"id"`
	StoreID   string `json:"storeId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      Roles  `json:"role"`
}

type EmployeeInput struct {
	ID        string `json:"id"`
	StoreID   string `json:"storeId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      Roles  `json:"role"`
}

type InventoryItem struct {
	ID           string        `json:"id"`
	Description  string        `json:"description"`
	Cost         float64       `json:"cost"`
	Retail       float64       `json:"retail"`
	Manufacturer *Manufacturer `json:"manufacturer"`
}

type InventoryItemInput struct {
	ID           string             `json:"id"`
	Description  string             `json:"description"`
	Cost         float64            `json:"cost"`
	Retail       float64            `json:"retail"`
	Manufacturer *ManufacturerInput `json:"manufacturer"`
}

type LineItem struct {
	Description string  `json:"description"`
	Retail      float64 `json:"retail"`
	Quantity    int     `json:"quantity"`
}

type LineItemInput struct {
	Description string  `json:"description"`
	Retail      float64 `json:"retail"`
	Quantity    int     `json:"quantity"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Manufacturer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ManufacturerInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewCustomerInput struct {
	FirstName   string            `json:"firstName"`
	LastName    string            `json:"lastName"`
	ContactInfo *ContactInfoInput `json:"contactInfo"`
}

type NewEmployeeInput struct {
	StoreID   string `json:"storeId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      Roles  `json:"role"`
}

type NewInventoryItemInput struct {
	Description  string             `json:"description"`
	Cost         float64            `json:"cost"`
	Retail       float64            `json:"retail"`
	Manufacturer *ManufacturerInput `json:"manufacturer"`
}

type NewManufacturerInput struct {
	Name string `json:"name"`
}

type NewOrderInput struct {
	CustomerID       string           `json:"customerId"`
	StoreID          string           `json:"storeId"`
	SalesAssociateID string           `json:"salesAssociateId"`
	LineItems        []*LineItemInput `json:"lineItems"`
}

type NewStoreInput struct {
	Name     string              `json:"name"`
	Location *StoreLocationInput `json:"location"`
}

type NewUserInput struct {
	EmployeeID string `json:"employeeId"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Order struct {
	ID             string      `json:"id"`
	StoreID        string      `json:"storeId"`
	Customer       *Customer   `json:"customer"`
	SalesAssociate *Employee   `json:"salesAssociate"`
	LineItems      []*LineItem `json:"lineItems"`
}

type OrderInput struct {
	ID               string           `json:"id"`
	CustomerID       string           `json:"customerId"`
	StoreID          string           `json:"storeId"`
	SalesAssociateID string           `json:"salesAssociateId"`
	LineItems        []*LineItemInput `json:"lineItems"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Store struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Location *StoreLocation `json:"location"`
}

type StoreFilter struct {
	Name  *string `json:"name"`
	City  *string `json:"city"`
	State *string `json:"state"`
}

type StoreInput struct {
	ID       string              `json:"id"`
	Name     string              `json:"name"`
	Location *StoreLocationInput `json:"location"`
}

type StoreLocation struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

type StoreLocationInput struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

type User struct {
	ID         string `json:"id"`
	EmployeeID string `json:"employeeId"`
	Username   string `json:"username"`
}

type UserInput struct {
	ID              string  `json:"id"`
	EmployeeID      string  `json:"employeeId"`
	Username        string  `json:"username"`
	CurrentPassword string  `json:"currentPassword"`
	NewPassword     *string `json:"newPassword"`
}

type Roles string

const (
	RolesManager        Roles = "MANAGER"
	RolesSalesAssociate Roles = "SALES_ASSOCIATE"
)

var AllRoles = []Roles{
	RolesManager,
	RolesSalesAssociate,
}

func (e Roles) IsValid() bool {
	switch e {
	case RolesManager, RolesSalesAssociate:
		return true
	}
	return false
}

func (e Roles) String() string {
	return string(e)
}

func (e *Roles) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Roles(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Roles", str)
	}
	return nil
}

func (e Roles) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
