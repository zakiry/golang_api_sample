package conf

import (
	"github.com/jinzhu/configor"
	"github.com/zakiry/golang_api_sample/entity"
)

func AppConf() entity.AppConf {
	appConf := entity.AppConf{}

	configor.Load(&appConf, "conf/app_conf.yml")

	return appConf
}
