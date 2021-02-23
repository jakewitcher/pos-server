package sqlite

import (
	"database/sql"
	"github.com/jakewitcher/pos-server/graph/model"
	"github.com/jakewitcher/pos-server/internal/users"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

type UserProvider struct {
	db *sql.DB
}

func (p *UserProvider) CreateUser(newUser model.NewUserInput) (*model.User, error) {
	employeeId, err := strconv.ParseInt(newUser.EmployeeID, 10, 64)

	if err != nil {
		return nil, newInvalidIdError(Employee, newUser.EmployeeID)
	}

	userId, err := p.insertNewUser(newUser, employeeId)

	user := &users.UserEntity{
		Id: userId,
		EmployeeId: employeeId,
		Username: newUser.Username,
	}

	return user.ToDTO(), nil
}

func (p *UserProvider) insertNewUser(newUser model.NewUserInput, employeeId int64) (int64, error) {
	hashedPassword, err := HashPassword(newUser.Password)

	if err != nil {
		return 0, serverError
	}

	statement, err := p.db.Prepare(
		`INSERT INTO User(EmployeeId, Username, Password) 
			   VALUES (?,?,?)`)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	defer statement.Close()

	result, err := statement.Exec(employeeId, newUser.Username, hashedPassword)

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	userId, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		return 0, serverError
	}

	return userId, nil
}

func NewUserProvider(db *sql.DB) *UserProvider {
	return &UserProvider{db: db}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}