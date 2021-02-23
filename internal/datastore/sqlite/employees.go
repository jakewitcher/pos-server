package sqlite

import (
	"database/sql"
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/employees"
	"log"
	"strconv"
)

type EmployeeProvider struct {
	db *sql.DB
}

func (p *EmployeeProvider) CreateEmployee(newEmployee model.NewEmployeeInput) (*model.Employee, error) {
	storeId, err := strconv.ParseInt(newEmployee.StoreID, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Store, newEmployee.StoreID)
	}

	employeeId, err := p.insertNewEmployee(newEmployee, storeId)

	employee := &employees.EmployeeEntity{
		Id:        employeeId,
		StoreId:   storeId,
		FirstName: newEmployee.FirstName,
		LastName:  newEmployee.LastName,
	}

	return employee.ToDTO(), nil
}

func (p *EmployeeProvider) insertNewEmployee(newEmployee model.NewEmployeeInput, storeId int64) (int64, error) {
	statement, err := p.db.Prepare(
		`INSERT INTO Employee(StoreId, FirstName, LastName, Role)
			   VALUES (?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	defer statement.Close()

	result, err := statement.Exec(storeId, newEmployee.FirstName, newEmployee.LastName, newEmployee.Role)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	employeeId, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	return employeeId, nil
}

func (p *EmployeeProvider) UpdateEmployee(updatedEmployee model.EmployeeInput) (*model.Employee, error) {
	employeeId, err := strconv.ParseInt(updatedEmployee.ID, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Employee, updatedEmployee.ID)
	}

	storeId, err := strconv.ParseInt(updatedEmployee.StoreID, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Store, updatedEmployee.StoreID)
	}

	err = p.updateEmployee(updatedEmployee, storeId, employeeId)

	if err != nil {
		return nil, err
	}

	employee := &employees.EmployeeEntity{
		Id:        employeeId,
		StoreId:   storeId,
		FirstName: updatedEmployee.FirstName,
		LastName:  updatedEmployee.LastName,
	}

	return employee.ToDTO(), nil
}

func (p *EmployeeProvider) updateEmployee(updatedEmployee model.EmployeeInput, storeId int64, employeeId int64) error {
	statement, err := p.db.Prepare(
		`UPDATE Employee
			   SET StoreId = ?,
			       FirstName = ?,
			   	   LastName = ?
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(storeId, updatedEmployee.FirstName, updatedEmployee.LastName, employeeId)

	if err != nil {
		log.Println(err)
		return serverError
	}
	return nil
}

func (p *EmployeeProvider) DeleteEmployee(employeeId string) (*model.Employee, error) {
	id, err := strconv.ParseInt(employeeId, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Employee, employeeId)
	}

	employee, err := p.findEmployeeById(id)

	if err != nil {
		return nil, err
	}

	err = p.deleteEmployee(id)

	if err != nil {
		return nil, err
	}

	return employee.ToDTO(), nil
}

func (p *EmployeeProvider) deleteEmployee(id int64) error {
	statement, err := p.db.Prepare(
		`DELETE FROM Employee 
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	_, err = statement.Exec(id)

	if err != nil {
		log.Println(err)
		return serverError
	}

	return nil
}

func (p *EmployeeProvider) FindEmployeeById(employeeId string) (*model.Employee, error) {
	id, err := strconv.ParseInt(employeeId, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Employee, employeeId)
	}

	employee, err := p.findEmployeeById(id)

	if err != nil {
		return nil, err
	}

	return employee.ToDTO(), nil
}

func (p *EmployeeProvider) FindEmployees() ([]*model.Employee, error) {
	statement, err := p.db.Prepare(
		`SELECT Id, StoreId, FirstName, LastName 
			   FROM Employee`)

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	defer statement.Close()

	rows, err := statement.Query()

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	defer rows.Close()

	employeeModels := make([]*model.Employee, 0)

	for rows.Next() {
		employee := &employees.EmployeeEntity{}

		err := rows.Scan(
			&employee.Id,
			&employee.StoreId,
			&employee.FirstName,
			&employee.LastName)

		if err != nil {
			log.Println(err)
			return nil, serverError
		}

		employeeModel := employee.ToDTO()
		employeeModels = append(employeeModels, employeeModel)
	}

	return employeeModels, nil
}

func (p *EmployeeProvider) findEmployeeById(employeeId int64) (*employees.EmployeeEntity, error) {
	statement, err := p.db.Prepare(
		`SELECT Id, StoreId, FirstName, LastName
			   FROM Employee
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	row := statement.QueryRow(employeeId)

	employee := &employees.EmployeeEntity{}

	err = row.Scan(
		&employee.Id,
		&employee.StoreId,
		&employee.FirstName,
		&employee.LastName)

	if err == sql.ErrNoRows {
		return nil, newNotFoundError(Employee, employeeId)
	}

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	return employee, nil
}

func NewEmployeeProvider(db *sql.DB) *EmployeeProvider {
	return &EmployeeProvider{db: db}
}
