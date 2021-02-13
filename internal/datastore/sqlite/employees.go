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

func (p *EmployeeProvider) CreateManager(newManager model.NewManagerInput) (*model.Manager, error) {
	storeId, err := strconv.ParseInt(newManager.StoreID, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Store, newManager.StoreID)
	}

	managerId, err := p.insertNewManager(newManager, storeId)

	manager := &employees.ManagerEntity{
		Id:        managerId,
		StoreId:   storeId,
		FirstName: newManager.FirstName,
		LastName:  newManager.LastName,
	}

	return manager.ToDTO(), nil
}

func (p *EmployeeProvider) insertNewManager(newManager model.NewManagerInput, storeId int64) (int64, error) {
	statement, err := p.db.Prepare(
		`INSERT INTO Manager(StoreId, FirstName, LastName, Password)
			   VALUES (?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	defer statement.Close()

	result, err := statement.Exec(storeId, newManager.FirstName, newManager.LastName, newManager.Password)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	managerId, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	return managerId, nil
}

func (p *EmployeeProvider) UpdateManager(updatedManager model.ManagerInput) (*model.Manager, error) {
	managerId, err := strconv.ParseInt(updatedManager.ID, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Manager, updatedManager.ID)
	}

	storeId, err := strconv.ParseInt(updatedManager.StoreID, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Store, updatedManager.StoreID)
	}

	err = p.updateManager(updatedManager, storeId, managerId)

	if err != nil {
		return nil, err
	}

	manager := &employees.ManagerEntity{
		Id:        managerId,
		StoreId:   storeId,
		FirstName: updatedManager.FirstName,
		LastName:  updatedManager.LastName,
	}

	return manager.ToDTO(), nil
}

func (p *EmployeeProvider) updateManager(updatedManager model.ManagerInput, storeId int64, managerId int64) error {
	statement, err := p.db.Prepare(
		`UPDATE Manager
			   SET StoreId = ?,
			       FirstName = ?,
			   	   LastName = ?
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(storeId, updatedManager.FirstName, updatedManager.LastName, managerId)

	if err != nil {
		log.Println(err)
		return serverError
	}
	return nil
}

func NewEmployeeProvider(db *sql.DB) *EmployeeProvider {
	return &EmployeeProvider{db: db}
}
