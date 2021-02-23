package sqlite

import "database/sql"

type UserProvider struct {
	db *sql.DB
}

func NewUserProvider(db *sql.DB) *EmployeeProvider {
	return &EmployeeProvider{db: db}
}