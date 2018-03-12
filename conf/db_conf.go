package conf

import (
	"github.com/jinzhu/configor"
	"github.com/zakiry/golang_api_sample/entity"
)

func DbConf() entity.DbConf {
	dbConf := entity.DbConf{}

	configor.Load(&dbConf, "conf/db_conf.yml")

	return dbConf
}
