package sqlite

import "fmt"

const (
	Customer = "Customer"
	Store    = "Store"
)

var (
	serverError = internalServerError{}
)

type invalidIdError struct {
	id     string
	entity string
}

func (e invalidIdError) Error() string {
	return fmt.Sprintf("%s id must be an integer, %s is not a valid id", e.entity, e.id)
}

func newInvalidIdError(entity, id string) invalidIdError {
	return invalidIdError{id: id, entity: entity}
}

type notFoundError struct {
	id     int64
	entity string
}

func (e notFoundError) Error() string {
	return fmt.Sprintf("no %s found with id %d", e.entity, e.id)
}

func newNotFoundError(entity string, id int64) notFoundError {
	return notFoundError{id: id, entity: entity}
}

type internalServerError struct{}

func (e internalServerError) Error() string {
	return "internal server error"
}
