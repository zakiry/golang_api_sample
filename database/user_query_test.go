package database

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/zakiry/golang_api_sample/entity"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUserImpl_GetById(t *testing.T) {
	var db *gorm.DB
	_, mock, err := sqlmock.NewWithDSN("sqlmock_db_model_user")
	if err != nil {
		panic("Got an unexpected error.")
	}
	db, err = gorm.Open("sqlmock", "sqlmock_db_model_user")
	if err != nil {
		panic("Got an unexpected error.")
	}
	defer db.Close()

	records := sqlmock.
		NewRows([]string{"id", "name", "age"}).
		AddRow(1, "piyo piyo", 12)

	userId := int64(1)
	mock.ExpectQuery(
		`SELECT id, name, age ` +
			`FROM "users" ` +
			`WHERE \(id = \?\)`).
		WithArgs(userId).
		WillReturnRows(records)

	user := userImpl{db: db}
	actual := user.GetById(userId)

	assert.Equal(t, entity.User{Id: 1, Name: "piyo piyo", Age: 12}, actual)
}
