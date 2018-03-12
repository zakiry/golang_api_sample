package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/zakiry/golang_api_sample/entity"
)

var sampleDb *gorm.DB

type Database interface {
	Open() *gorm.DB
	Close()
	Db() *gorm.DB
}

type databaseImpl struct {
	conf entity.DbConf
}

// This function is called once when server start.
func (s *databaseImpl) Open() *gorm.DB {
	options := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		s.conf.User,
		s.conf.Password,
		s.conf.Address,
		s.conf.Port,
		s.conf.Dbname,
	)

	db, err := gorm.Open(s.conf.Dialect, options)
	if err != nil {
		log.Fatalf("Can't open : %v", err)
		panic(err.Error())
	}

	sampleDb = db
	return sampleDb
}

// This function is called once when server end.
func (s *databaseImpl) Close() {
	s.Close()
}

func (s *databaseImpl) Db() *gorm.DB {
	return sampleDb
}

func NewDatabase(conf entity.DbConf) Database {
	return &databaseImpl{conf}
}
