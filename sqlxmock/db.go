package sqlxmock

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type DBTemplate struct {
	mock.Mock
}

func (db DBTemplate) MustBegin() *sqlx.Tx {
	argsMock := db.Called()

	return argsMock.Get(0).(*sqlx.Tx)
}

func (db DBTemplate) Preparex(query string) (*sqlx.Stmt, error) {
	argsMock := db.Called(query)

	return argsMock.Get(0).(*sqlx.Stmt), argsMock.Error(1)
}

func (db DBTemplate) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	argsMock := db.Called(query, args)

	return argsMock.Get(0).(*sqlx.Rows), argsMock.Error(1)
}

func (db DBTemplate) Get(dest interface{}, query string, param ...interface{}) error {
	argsMock := db.Called(dest, query)

	return argsMock.Error(0)
}

func (db DBTemplate) Beginx() (*sqlx.Tx, error) {
	argsMock := db.Called()

	return argsMock.Get(0).(*sqlx.Tx), argsMock.Error(1)
}
