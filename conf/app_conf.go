package conf

import (
	"github.com/jinzhu/configor"
	"github.com/zakiry/golang_api_sample/entity"
)

func AppConf() entity.AppConf {
	builder := entity.AppConfBuilder{}

	configor.Load(&builder, "conf/app_conf.yml")

	return builder.Build()
}
