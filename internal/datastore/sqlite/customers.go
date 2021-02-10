package sqlite

import (
	"database/sql"
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/customers"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

type CustomerProvider struct {
	db *sql.DB
}

func (p *CustomerProvider) CreateCustomer(newCustomer model.NewCustomerInput) *model.Customer {
	newContactInfo := newCustomer.ContactInfo
	contactInfoId := p.insertNewContactInfo(newContactInfo)
	customerId := p.insertNewCustomer(newCustomer, contactInfoId)

	customer := &customers.CustomerEntity{
		Id:            customerId,
		FirstName:     newCustomer.FirstName,
		LastName:      newCustomer.LastName,
		ContactInfoId: contactInfoId,
	}

	contactInfo := &customers.ContactInfoEntity{
		Id:           contactInfoId,
		EmailAddress: newContactInfo.EmailAddress,
		PhoneNumber:  newContactInfo.PhoneNumber,
	}

	return customer.ToDTO(contactInfo)
}

func (p *CustomerProvider) insertNewCustomer(newCustomer model.NewCustomerInput, contactInfoId int64) int64 {
	statement, err := p.db.Prepare(
		`INSERT INTO Customer(FirstName, LastName, ContactInfoId) 
			   VALUES (?,?,?)`)
	checkError(err)

	defer statement.Close()

	result, err := statement.Exec(newCustomer.FirstName, newCustomer.LastName, contactInfoId)
	checkError(err)

	customerId, err := result.LastInsertId()
	checkError(err)

	return customerId
}

func (p *CustomerProvider) insertNewContactInfo(newContactInfo *model.ContactInfoInput) int64 {
	statement, err := p.db.Prepare(
		`INSERT INTO ContactInfo(EmailAddress, PhoneNumber) 
			   VALUES (?,?)`)
	checkError(err)

	defer statement.Close()

	result, err := statement.Exec(newContactInfo.EmailAddress, newContactInfo.PhoneNumber)
	checkError(err)

	contactInfoId, err := result.LastInsertId()
	checkError(err)

	return contactInfoId
}

func (p *CustomerProvider) UpdateCustomer(updatedCustomer model.CustomerInput) *model.Customer {
	updatedContactInfo := updatedCustomer.ContactInfo
	contactInfoId := p.getContactInfoIdByCustomerId(updatedCustomer)

	p.updateContactInfo(updatedContactInfo, contactInfoId)
	p.updateCustomer(updatedCustomer)

	contactInfo := &customers.ContactInfoEntity{
		Id: contactInfoId,
		EmailAddress: updatedContactInfo.EmailAddress,
		PhoneNumber: updatedContactInfo.PhoneNumber,
	}

	customerId, err := strconv.Atoi(updatedCustomer.ID)
	checkError(err)

	customer := &customers.CustomerEntity{
		Id: int64(customerId),
		FirstName: updatedCustomer.FirstName,
		LastName: updatedCustomer.LastName,
		ContactInfoId: contactInfoId,
	}

	return customer.ToDTO(contactInfo)
}

func (p *CustomerProvider) updateCustomer(updatedCustomer model.CustomerInput) {
	statement, err := p.db.Prepare(
		`UPDATE Customer
			   SET FirstName = ?,
			   	   LastName = ?
			   WHERE Id = ?`)
	checkError(err)

	defer statement.Close()

	customerId, err := strconv.Atoi(updatedCustomer.ID)
	checkError(err)

	_, err = statement.Exec(updatedCustomer.FirstName, updatedCustomer.LastName, customerId)
	checkError(err)
}

func (p *CustomerProvider) updateContactInfo(updatedContactInfo *model.ContactInfoInput, contactInfoId int64) {
	statement, err := p.db.Prepare(
		`UPDATE ContactInfo
			   SET EmailAddress = ?,
				   PhoneNumber = ?
			   WHERE Id = ?`)
	checkError(err)

	defer statement.Close()
	
	_, err = statement.Exec(updatedContactInfo.EmailAddress, updatedContactInfo.PhoneNumber, contactInfoId)
	checkError(err)
}

func (p *CustomerProvider) getContactInfoIdByCustomerId(updatedCustomer model.CustomerInput) int64 {
	statement, err := p.db.Prepare(
		`SELECT ContactInfoId FROM Customer WHERE Id = ?`)
	checkError(err)

	defer statement.Close()

	row := statement.QueryRow(updatedCustomer.ID)
	var contactInfoId int64

	err = row.Scan(&contactInfoId)
	checkError(err)
	
	return contactInfoId
}

func (p *CustomerProvider) DeleteCustomer(customerId string) *model.Customer {
	customer, contactInfo := p.findCustomerAndContactInfoByCustomerId(customerId)

	p.deleteContactInfoById(contactInfo.Id)
	p.deleteCustomerById(customer.Id)

	return customer.ToDTO(contactInfo)
}

func (p *CustomerProvider) deleteCustomerById(customerId int64) {
	statement, err := p.db.Prepare(
		`DELETE FROM Customer WHERE Id = ?`)
	checkError(err)

	defer statement.Close()

	_, err = statement.Exec(customerId)
	checkError(err)
}

func (p *CustomerProvider) deleteContactInfoById(contactInfoId int64) {
	statement, err := p.db.Prepare(
		`DELETE FROM ContactInfo WHERE Id = ?`)
	checkError(err)

	defer statement.Close()

	_, err = statement.Exec(contactInfoId)
	checkError(err)
}

func (p *CustomerProvider) FindCustomerById(customerId string) *model.Customer {
	customer, contactInfo := p.findCustomerAndContactInfoByCustomerId(customerId)
	return customer.ToDTO(contactInfo)
}

func (p *CustomerProvider) findCustomerAndContactInfoByCustomerId(customerId string) (*customers.CustomerEntity, *customers.ContactInfoEntity) {
	statement, err := p.db.Prepare(
		`SELECT C.Id, C.FirstName, C.LastName, C.ContactInfoId, CI.Id, CI.EmailAddress, CI.PhoneNumber
			   FROM Customer C INNER JOIN ContactInfo CI 
			   ON C.ContactInfoId = CI.Id
			   WHERE C.Id = ?`)
	checkError(err)

	defer statement.Close()

	row := statement.QueryRow(customerId)

	customer := &customers.CustomerEntity{}
	contactInfo := &customers.ContactInfoEntity{}

	err = row.Scan(
		&customer.Id,
		&customer.FirstName,
		&customer.LastName,
		&customer.ContactInfoId,
		&contactInfo.Id,
		&contactInfo.EmailAddress,
		&contactInfo.PhoneNumber)
	checkError(err)

	return customer, contactInfo
}

func (p *CustomerProvider) GetAllCustomers() []*model.Customer {
	statement, err := p.db.Prepare(
		`SELECT C.Id, C.FirstName, C.LastName, CI.EmailAddress, CI.PhoneNumber
			   FROM Customer C INNER JOIN ContactInfo CI 
			   ON C.ContactInfoId = CI.Id`)
	checkError(err)

	rows, err := statement.Query()
	checkError(err)

	customerModels := make([]*model.Customer, 0)

	for rows.Next() {
		customer := &customers.CustomerEntity{}
		contactInfo := &customers.ContactInfoEntity{}

		err := rows.Scan(
			&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&contactInfo.EmailAddress,
			&contactInfo.PhoneNumber)
		checkError(err)

		customerModel := customer.ToDTO(contactInfo)
		customerModels = append(customerModels, customerModel)
	}

	return customerModels
}

func NewCustomerProvider(db *sql.DB) *CustomerProvider {
	return &CustomerProvider{db: db}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}