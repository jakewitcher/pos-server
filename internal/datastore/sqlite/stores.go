package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/stores"
	"strconv"
)

type StoreProvider struct {
	db *sql.DB
}

func (p *StoreProvider) CreateStore(newStore model.NewStoreInput) *model.Store {
	panic(fmt.Errorf("not implemented"))
}

func (p *StoreProvider) UpdateStore(updatedStore model.StoreInput) *model.Store {
	panic(fmt.Errorf("not implemented"))
}

func (p *StoreProvider) DeleteStore(storeId string) *model.Store {
	panic(fmt.Errorf("not implemented"))
}

func (p *StoreProvider) FindStoreById(storeId string) *model.Store {
	panic(fmt.Errorf("not implemented"))
}

func (p *StoreProvider) FindAllStores() []*model.Store {
	panic(fmt.Errorf("not implemented"))
}

func NewStoreProvider(db *sql.DB) *StoreProvider {
	return &StoreProvider{db: db}
}

type StoreLocationProvider struct {
	db *sql.DB
}

func (p *StoreLocationProvider) CreateStoreLocation(newStoreLocation model.NewStoreLocationInput) *model.StoreLocation {
	storeLocationId := p.insertNewStoreLocation(newStoreLocation)

	return &model.StoreLocation{
		ID:      strconv.FormatInt(storeLocationId, 10),
		Street:  newStoreLocation.Street,
		City:    newStoreLocation.City,
		State:   newStoreLocation.State,
		ZipCode: newStoreLocation.ZipCode,
	}
}

func (p *StoreLocationProvider) insertNewStoreLocation(newStoreLocation model.NewStoreLocationInput) int64 {
	statement, err := p.db.Prepare(
		`INSERT INTO StoreLocation(Street, City, State, ZipCode) VALUES (?,?,?,?)`)
	checkError(err)

	defer statement.Close()

	result, err := statement.Exec(
		newStoreLocation.Street,
		newStoreLocation.City,
		newStoreLocation.State,
		newStoreLocation.ZipCode)
	checkError(err)

	storeLocationId, err := result.LastInsertId()
	checkError(err)

	return storeLocationId
}

func (p *StoreLocationProvider) UpdateStoreLocation(updatedStoreLocation model.StoreLocationInput) *model.StoreLocation {
	storeLocationId, err := strconv.ParseInt(updatedStoreLocation.ID, 10, 64)
	checkError(err)

	p.updateStoreLocation(updatedStoreLocation, storeLocationId)

	storeLocation := &stores.StoreLocationEntity{
		Id:      storeLocationId,
		Street:  updatedStoreLocation.Street,
		City:    updatedStoreLocation.City,
		State:   updatedStoreLocation.State,
		ZipCode: updatedStoreLocation.ZipCode,
	}

	return storeLocation.ToDTO()
}

func (p *StoreLocationProvider) updateStoreLocation(updatedStoreLocation model.StoreLocationInput, storeLocationId int64) {
	statement, err := p.db.Prepare(
		`UPDATE StoreLocation
			   SET Street = ?,
				   City = ?,
				   State = ?,
				   ZipCode = ?
			   WHERE Id = ?`)
	checkError(err)

	defer statement.Close()

	_, err = statement.Exec(
		updatedStoreLocation.Street,
		updatedStoreLocation.City,
		updatedStoreLocation.State,
		updatedStoreLocation.ZipCode,
		storeLocationId)
	checkError(err)
}

func (p *StoreLocationProvider) DeleteStoreLocation(storeLocationId string) *model.StoreLocation {
	storeLocation := p.findStoreLocationById(storeLocationId)

	statement, err := p.db.Prepare(
		`DELETE FROM StoreLocation WHERE Id = ?`)
	checkError(err)

	defer statement.Close()

	_, err = statement.Exec(storeLocation.Id)
	checkError(err)

	return storeLocation.ToDTO()
}

func (p *StoreLocationProvider) FindStoreLocationById(storeLocationId string) *model.StoreLocation {
	return p.findStoreLocationById(storeLocationId).ToDTO()
}

func (p *StoreLocationProvider) findStoreLocationById(storeLocationId string) *stores.StoreLocationEntity {
	id, err := strconv.ParseInt(storeLocationId, 10, 64)
	checkError(err)

	statement, err := p.db.Prepare(
		`SELECT Id, Street, City, State, ZipCode 
				   FROM StoreLocation 
				   WHERE Id = ?`)
	checkError(err)

	defer statement.Close()

	row := statement.QueryRow(id)

	storeLocation := &stores.StoreLocationEntity{}

	err = row.Scan(
		&storeLocation.Id,
		&storeLocation.Street,
		&storeLocation.City,
		&storeLocation.State,
		&storeLocation.ZipCode)
	checkError(err)

	return storeLocation
}

func (p *StoreLocationProvider) FindAllStoreLocations() []*model.StoreLocation {
	statement, err := p.db.Prepare(
		`SELECT Id, Street, City, State, ZipCode 
			   FROM StoreLocation`)
	checkError(err)

	defer statement.Close()

	rows, err := statement.Query()
	checkError(err)

	storeLocationModels := make([]*model.StoreLocation, 0)

	for rows.Next() {
		storeLocation := &stores.StoreLocationEntity{}

		err := rows.Scan(
			&storeLocation.Id,
			&storeLocation.Street,
			&storeLocation.City,
			&storeLocation.State,
			&storeLocation.ZipCode)
		checkError(err)

		storeLocationModel := storeLocation.ToDTO()
		storeLocationModels = append(storeLocationModels, storeLocationModel)
	}

	return storeLocationModels
}

func NewStoreLocationProvider(db *sql.DB) *StoreLocationProvider {
	return &StoreLocationProvider{db: db}
}
