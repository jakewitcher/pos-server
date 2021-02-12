package sqlite

import (
	"database/sql"
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/stores"
	"log"
	"strconv"
	"strings"
)

type StoreProvider struct {
	db *sql.DB
}

func (p *StoreProvider) CreateStore(newStore model.NewStoreInput) (*model.Store, error) {
	newStoreLocation := newStore.Location
	trx, err := p.db.Begin()

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	storeLocationId, err := p.insertNewStoreLocation(trx, newStoreLocation)

	if err != nil {
		if rollbackErr := trx.Rollback(); rollbackErr != nil {
			log.Println(rollbackErr)
		}

		return nil, err
	}

	storeId, err := p.insertNewStore(trx, newStore, storeLocationId)

	if err != nil {
		if rollbackErr := trx.Rollback(); rollbackErr != nil {
			log.Println(rollbackErr)
		}

		return nil, err
	}

	err = trx.Commit()

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	store := &stores.StoreEntity{
		Id:         storeId,
		Name:       newStore.Name,
		LocationId: storeLocationId,
	}

	location := &stores.StoreLocationEntity{
		Id:      storeLocationId,
		Street:  newStoreLocation.Street,
		City:    newStoreLocation.City,
		State:   newStoreLocation.State,
		ZipCode: newStoreLocation.ZipCode,
	}

	return store.ToDTO(location), nil
}

func (*StoreProvider) insertNewStoreLocation(trx *sql.Tx, newStoreLocation *model.StoreLocationInput) (int64, error) {
	statement, err := trx.Prepare(
		`INSERT INTO StoreLocation(Street, City, State, ZipCode) 
			   VALUES (?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	defer statement.Close()

	result, err := statement.Exec(
		newStoreLocation.Street,
		newStoreLocation.City,
		newStoreLocation.State,
		newStoreLocation.ZipCode)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	storeLocationId, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	return storeLocationId, nil
}

func (*StoreProvider) insertNewStore(trx *sql.Tx, newStore model.NewStoreInput, storeLocationId int64) (int64, error) {
	statement, err := trx.Prepare(
		`INSERT INTO Store(Name, LocationId) 
			   VALUES (?,?)`)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	defer statement.Close()

	result, err := statement.Exec(newStore.Name, storeLocationId)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	storeId, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	return storeId, nil
}

func (p *StoreProvider) UpdateStore(updatedStore model.StoreInput) (*model.Store, error) {
	updatedStoreLocation := updatedStore.Location
	trx, err := p.db.Begin()

	if err != nil {
		log.Print(err)
		return nil, serverError
	}

	storeId, err := strconv.ParseInt(updatedStore.ID, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Store, updatedStore.ID)
	}

	storeLocationId, err := p.getStoreLocationIdByStoreId(storeId)

	if err != nil {
		return nil, err
	}

	err = p.updateStoreLocation(trx, updatedStoreLocation, storeLocationId)

	if err != nil {
		if rollbackErr := trx.Rollback(); rollbackErr != nil {
			log.Println(rollbackErr)
		}

		return nil, err
	}

	err = p.updateStore(trx, updatedStore, storeId)

	if err != nil {
		if rollbackErr := trx.Rollback(); rollbackErr != nil {
			log.Println(rollbackErr)
		}

		return nil, err
	}

	err = trx.Commit()

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	storeLocation := &stores.StoreLocationEntity{
		Id:      storeLocationId,
		Street:  updatedStoreLocation.Street,
		City:    updatedStoreLocation.City,
		State:   updatedStoreLocation.State,
		ZipCode: updatedStoreLocation.ZipCode,
	}

	store := &stores.StoreEntity{
		Id:         storeId,
		LocationId: storeLocationId,
	}

	return store.ToDTO(storeLocation), nil
}

func (*StoreProvider) updateStoreLocation(trx *sql.Tx, updatedStoreLocation *model.StoreLocationInput, storeLocationId int64) error {
	statement, err := trx.Prepare(
		`UPDATE StoreLocation
			   SET Street = ?,
				   City = ?,
				   State = ?,
				   ZipCode = ?
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(
		updatedStoreLocation.Street,
		updatedStoreLocation.City,
		updatedStoreLocation.State,
		updatedStoreLocation.ZipCode,
		storeLocationId)

	if err != nil {
		log.Println(err)
		return serverError
	}

	return nil
}

func (*StoreProvider) updateStore(trx *sql.Tx, updatedStore model.StoreInput, storeId int64) error {
	statement, err := trx.Prepare(
		`UPDATE Store
			   SET Name = ?
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(updatedStore.Name, storeId)

	if err != nil {
		log.Println(err)
		return serverError
	}

	return nil
}

func (p *StoreProvider) getStoreLocationIdByStoreId(storeId int64) (int64, error) {
	statement, err := p.db.Prepare(
		`SELECT LocationId
			   FROM Store
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	defer statement.Close()

	row := statement.QueryRow(storeId)
	var locationId int64

	err = row.Scan(&locationId)

	if err == sql.ErrNoRows {
		return 0, newNotFoundError(Store, storeId)
	}

	if err != nil {
		return 0, serverError
	}

	return locationId, nil
}

func (p *StoreProvider) DeleteStore(storeId string) (*model.Store, error) {
	id, err := strconv.ParseInt(storeId, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Store, storeId)
	}

	trx, err := p.db.Begin()

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	store, storeLocation, err := p.findStoreAndStoreLocationByStoreId(id)

	if err != nil {
		return nil, err
	}

	err = p.deleteStoreLocation(trx, storeLocation.Id)

	if err != nil {
		if rollbackErr := trx.Rollback(); rollbackErr != nil {
			log.Println(rollbackErr)
		}

		return nil, err
	}

	err = p.deleteStore(trx, store.Id)

	if err != nil {
		if rollbackErr := trx.Rollback(); rollbackErr != nil {
			log.Println(rollbackErr)
		}

		return nil, err
	}

	err = trx.Commit()

	if err != nil {
		log.Println(err)
		return nil, serverError
	}

	return store.ToDTO(storeLocation), nil
}

func (*StoreProvider) deleteStoreLocation(trx *sql.Tx, storeLocationId int64) error {
	statement, err := trx.Prepare(
		`DELETE FROM StoreLocation 
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(storeLocationId)

	if err != nil {
		log.Println(err)
		return serverError
	}

	return nil
}

func (*StoreProvider) deleteStore(trx *sql.Tx, storeId int64) error {
	statement, err := trx.Prepare(
		`DELETE FROM Store 
			   WHERE Id = ?`)

	if err != nil {
		log.Println(err)
		return serverError
	}

	defer statement.Close()

	_, err = statement.Exec(storeId)

	if err != nil {
		log.Println(err)
		return serverError
	}

	return nil
}

func (p *StoreProvider) FindStoreById(storeId string) (*model.Store, error) {
	id, err := strconv.ParseInt(storeId, 10, 64)

	if err != nil {
		log.Println(err)
		return nil, newInvalidIdError(Store, storeId)
	}

	store, storeLocation, err := p.findStoreAndStoreLocationByStoreId(id)

	if err != nil {
		return nil, err
	}

	return store.ToDTO(storeLocation), nil
}

func (p *StoreProvider) findStoreAndStoreLocationByStoreId(storeId int64) (*stores.StoreEntity, *stores.StoreLocationEntity, error) {
	statement, err := p.db.Prepare(
		`SELECT S.Id, S.Name, S.LocationId, SL.Id, SL.Street, SL.City, SL.State, SL.ZipCode 
			   FROM Store S INNER JOIN StoreLocation SL 
    		   ON SL.Id = S.LocationId
    		   WHERE S.Id = ?`)

	if err != nil {
		log.Println(err)
		return nil, nil, serverError
	}

	defer statement.Close()

	row := statement.QueryRow(storeId)

	store := &stores.StoreEntity{}
	storeLocation := &stores.StoreLocationEntity{}

	err = row.Scan(
		&store.Id,
		&store.Name,
		&store.LocationId,
		&storeLocation.Id,
		&storeLocation.Street,
		&storeLocation.City,
		&storeLocation.State,
		&storeLocation.ZipCode)

	if err == sql.ErrNoRows {
		log.Println(err)
		return nil, nil, newNotFoundError(Store, storeId)
	}

	if err != nil {
		log.Println(err)
		return nil, nil, serverError
	}

	return store, storeLocation, nil
}

func (p *StoreProvider) FindStores(filter *model.StoreFilter) ([]*model.Store, error) {
	queryBase := `SELECT S.Id, S.Name, S.LocationId, SL.Id, SL.Street, SL.City, SL.State, SL.ZipCode 
			      FROM Store S INNER JOIN StoreLocation SL 
			      ON SL.Id = S.LocationId`

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

	storeModels := make([]*model.Store, 0)

	for rows.Next() {
		store := &stores.StoreEntity{}
		storeLocation := &stores.StoreLocationEntity{}

		err := rows.Scan(
			&store.Id,
			&store.Name,
			&store.LocationId,
			&storeLocation.Id,
			&storeLocation.Street,
			&storeLocation.City,
			&storeLocation.State,
			&storeLocation.ZipCode)

		if err != nil {
			log.Println(err)
			return nil, serverError
		}

		storeModel := store.ToDTO(storeLocation)
		storeModels = append(storeModels, storeModel)
	}

	return storeModels, nil
}

func (p *StoreProvider) buildQuery(base string, filter *model.StoreFilter) (string, []interface{}) {
	columns := make([]string, 0)
	values := make([]interface{}, 0)

	if filter == nil {
		return base, values
	}

	if filter.Name != nil {
		columns = append(columns, "S.Name = ?")
		values = append(values, *filter.Name)
	}

	if filter.City != nil {
		columns = append(columns, "SL.City = ?")
		values = append(values, *filter.City)
	}

	if filter.State != nil {
		columns = append(columns, "SL.State = ?")
		values = append(values, *filter.State)
	}

	if len(columns) == 0 {
		return base, values
	}

	query := base + "\nWHERE "
	query += strings.Join(columns, " AND ")

	return query, values
}

func NewStoreProvider(db *sql.DB) *StoreProvider {
	return &StoreProvider{db: db}
}
