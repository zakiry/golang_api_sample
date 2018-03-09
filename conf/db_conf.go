package conf

import (
	"github.com/jinzhu/configor"
	"github.com/zakiry/golang_api_sample/entity"
)

func DbConf() entity.DbConf {
	builder := entity.DbConfBuilder{}

	configor.Load(&builder, "conf/db_conf.yml")

	return builder.Build()
}
