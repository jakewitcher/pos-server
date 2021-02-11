package sqlite

import (
	"database/sql"
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/customers"
	"log"
	"strconv"
	"strings"
)

type CustomerProvider struct {
	db *sql.DB
}

func (p *CustomerProvider) CreateCustomer(newCustomer model.NewCustomerInput) (*model.Customer, error) {
	newContactInfo := newCustomer.ContactInfo

	contactInfoId, err := p.insertNewContactInfo(newContactInfo)

	if err != nil {
		return nil, err
	}

	customerId, err := p.insertNewCustomer(newCustomer, contactInfoId)

	if err != nil {
		return nil, err
	}

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

	return customer.ToDTO(contactInfo), nil
}

func (p *CustomerProvider) insertNewContactInfo(newContactInfo *model.ContactInfoInput) (int64, error) {
	statement, err := p.db.Prepare(
		`INSERT INTO ContactInfo(EmailAddress, PhoneNumber) 
			   VALUES (?,?)`)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	defer statement.Close()

	result, err := statement.Exec(newContactInfo.EmailAddress, newContactInfo.PhoneNumber)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	contactInfoId, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	return contactInfoId, nil
}

func (p *CustomerProvider) insertNewCustomer(newCustomer model.NewCustomerInput, contactInfoId int64) (int64, error) {
	statement, err := p.db.Prepare(
		`INSERT INTO Customer(FirstName, LastName, ContactInfoId) 
			   VALUES (?,?,?)`)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	defer statement.Close()

	result, err := statement.Exec(newCustomer.FirstName, newCustomer.LastName, contactInfoId)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	customerId, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	return customerId, nil
}

func (p *CustomerProvider) UpdateCustomer(updatedCustomer model.CustomerInput) (*model.Customer, error) {
	updatedContactInfo := updatedCustomer.ContactInfo

	customerId, err := strconv.ParseInt(updatedCustomer.ID, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(customer, updatedCustomer.ID)
	}

	contactInfoId, err := p.getContactInfoIdByCustomerId(customerId)

	if err != nil {
		return nil, err
	}

	err = p.updateContactInfo(updatedContactInfo, contactInfoId)

	if err != nil {
		return nil, err
	}

	err = p.updateCustomer(updatedCustomer, customerId)

	if err != nil {
		return nil, err
	}

	contactInfo := &customers.ContactInfoEntity{
		Id:           contactInfoId,
		EmailAddress: updatedContactInfo.EmailAddress,
		PhoneNumber:  updatedContactInfo.PhoneNumber,
	}

	customer := &customers.CustomerEntity{
		Id:            customerId,
		FirstName:     updatedCustomer.FirstName,
		LastName:      updatedCustomer.LastName,
		ContactInfoId: contactInfoId,
	}

	return customer.ToDTO(contactInfo), nil
}

func (p *CustomerProvider) updateContactInfo(updatedContactInfo *model.ContactInfoInput, contactInfoId int64) error {
	statement, err := p.db.Prepare(
		`UPDATE ContactInfo
			   SET EmailAddress = ?,
				   PhoneNumber = ?
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(updatedContactInfo.EmailAddress, updatedContactInfo.PhoneNumber, contactInfoId)

	if err != nil {
		log.Println(err)
		return serverError
	}

	return nil
}

func (p *CustomerProvider) updateCustomer(updatedCustomer model.CustomerInput, customerId int64) error {
	statement, err := p.db.Prepare(
		`UPDATE Customer
			   SET FirstName = ?,
			   	   LastName = ?
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(updatedCustomer.FirstName, updatedCustomer.LastName, customerId)

	if err != nil {
		log.Println(err)
		return serverError
	}

	return nil
}

func (p *CustomerProvider) getContactInfoIdByCustomerId(customerId int64) (int64, error) {
	statement, err := p.db.Prepare(
		`SELECT ContactInfoId 
			   FROM Customer 
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	defer statement.Close()

	row := statement.QueryRow(customerId)
	var contactInfoId int64

	err = row.Scan(&contactInfoId)

	if err == sql.ErrNoRows {
		return 0, newNotFoundError(customer, customerId)
	}

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	return contactInfoId, nil
}

func (p *CustomerProvider) DeleteCustomer(customerId string) (*model.Customer, error) {
	id, err := strconv.ParseInt(customerId, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(customer, customerId)
	}

	customer, contactInfo, err := p.findCustomerAndContactInfoByCustomerId(id)

	if err != nil {
		return nil, err
	}

	err = p.deleteContactInfoById(contactInfo.Id)

	if err != nil {
		return nil, err
	}

	err = p.deleteCustomerById(customer.Id)

	if err != nil {
		return nil, err
	}

	return customer.ToDTO(contactInfo), nil
}

func (p *CustomerProvider) deleteCustomerById(customerId int64) error {
	statement, err := p.db.Prepare(
		`DELETE FROM Customer 
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(customerId)

	if err != nil {
		log.Println(err)
		return serverError
	}

	return nil
}

func (p *CustomerProvider) deleteContactInfoById(contactInfoId int64) error {
	statement, err := p.db.Prepare(
		`DELETE FROM ContactInfo 
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(contactInfoId)

	if err != nil {
		log.Println(err)
		return serverError
	}

	return nil
}

func (p *CustomerProvider) FindCustomerById(customerId string) (*model.Customer, error) {
	id, err := strconv.ParseInt(customerId, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(customer, customerId)
	}

	customer, contactInfo, err := p.findCustomerAndContactInfoByCustomerId(id)

	if err != nil {
		return nil, err
	}

	return customer.ToDTO(contactInfo), nil
}

func (p *CustomerProvider) findCustomerAndContactInfoByCustomerId(customerId int64) (*customers.CustomerEntity, *customers.ContactInfoEntity, error) {
	statement, err := p.db.Prepare(
		`SELECT C.Id, C.FirstName, C.LastName, C.ContactInfoId, CI.Id, CI.EmailAddress, CI.PhoneNumber
			   FROM Customer C INNER JOIN ContactInfo CI 
			   ON C.ContactInfoId = CI.Id
			   WHERE C.Id = ?`)

	if err != nil {
		log.Println(err)
		return nil, nil, serverError
	}

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

	if err != nil {
		log.Println(err)
		return nil, nil, serverError
	}

	return customer, contactInfo, nil
}

func (p *CustomerProvider) FindCustomers(filter *model.CustomerFilter) ([]*model.Customer, error) {
	queryBase := `SELECT C.Id, C.FirstName, C.LastName, CI.EmailAddress, CI.PhoneNumber
			   	  FROM Customer C INNER JOIN ContactInfo CI 
			   	  ON C.ContactInfoId = CI.Id`

	query, queryParameters := p.buildQuery(queryBase, filter)

	statement, err := p.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	defer statement.Close()

	rows, err := statement.Query(queryParameters...)

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	defer rows.Close()

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

		if err != nil {
			log.Println(err)
			return nil, serverError
		}

		customerModel := customer.ToDTO(contactInfo)
		customerModels = append(customerModels, customerModel)
	}

	return customerModels, nil
}

func (p *CustomerProvider) buildQuery(base string, filter *model.CustomerFilter) (string, []interface{}) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)

	if filter == nil {
		return base, values
	}

	if filter.LastName != nil {
		columns = append(columns, "C.LastName = ?")
		values = append(values, *filter.LastName)
	}

	if filter.PhoneNumber != nil {
		columns = append(columns, "CI.PhoneNumber = ?")
		values = append(values, *filter.PhoneNumber)
	}

	if filter.EmailAddress != nil {
		columns = append(columns, "CI.EmailAddress = ?")
		values = append(values, *filter.EmailAddress)
	}

	if len(columns) < 1 {
		return base, values
	}

	query := base + "\nWHERE "
	query += strings.Join(columns, " AND ")

	return query, values
}

func NewCustomerProvider(db *sql.DB) *CustomerProvider {
	return &CustomerProvider{db: db}
}
