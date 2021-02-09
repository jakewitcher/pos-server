package sqlite

import (
	"database/sql"
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/customers"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type CustomerProvider struct {
	db *sql.DB
}

func (p *CustomerProvider) FindCustomerById(id string) *model.Customer {
	statement, err := p.db.Prepare(
		`SELECT C.Id, C.FirstName, C.LastName, CI.EmailAddress, CI.PhoneNumber
			   FROM Customer C INNER JOIN ContactInfo CI 
			   ON C.ContactInfoId = CI.Id
			   WHERE C.Id = ?`)

	if err != nil {
		log.Fatalln(err)
	}

	defer statement.Close()
	row := statement.QueryRow(id)

	customer := &customers.CustomerEntity{}
	contactInfo := &customers.ContactInfoEntity{}

	err = row.Scan(
		&customer.Id,
		&customer.FirstName,
		&customer.LastName,
		&contactInfo.EmailAddress,
		&contactInfo.PhoneNumber)

	if err != nil {
		log.Fatalln(err)
	}

	return customer.ToDTO(contactInfo)
}

func (p *CustomerProvider) GetAllCustomers() []*model.Customer {
	statement, err := p.db.Prepare(
		`SELECT C.Id, C.FirstName, C.LastName, CI.EmailAddress, CI.PhoneNumber
			   FROM Customer C INNER JOIN ContactInfo CI 
			   ON C.ContactInfoId = CI.Id`)

	if err != nil {
		log.Fatalln(err)
	}

	rows, err := statement.Query()
	if err != nil {
		log.Fatalln(err)
	}

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
			log.Fatalln(err)
		}

		customerModel := customer.ToDTO(contactInfo)
		customerModels = append(customerModels, customerModel)
	}

	return customerModels
}

func NewCustomerProvider() *CustomerProvider {
	db, err := sql.Open("sqlite3", "./pos.db")
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &CustomerProvider{db: db}
}
