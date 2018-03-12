package main

import (
	"net/http"

	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/zakiry/golang_api_sample/conf"
	"github.com/zakiry/golang_api_sample/database"
	pb "github.com/zakiry/golang_api_sample/proto"
	"github.com/zakiry/golang_api_sample/service"
)

var (
	Version  string
	Revision string
)

type services struct {
	userService pb.UserService
}

func main() {
	fmt.Println(fmt.Sprintf("AppName: sample api, version: %s, Revision: %s", Version, Revision))

	db := database.NewDatabase(conf.DbConf())
	db.Open()
	defer db.Close()

	services := initServices(db.Db())

	mux := http.NewServeMux()
	mux.Handle(pb.UserServicePathPrefix, pb.NewUserServiceServer(services.userService, nil))

	appConf := conf.AppConf()
	http.ListenAndServe(fmt.Sprintf(":%d", appConf.Port), mux)
}

func initServices(db *gorm.DB) services {
	user := database.NewUser(db)

	userService := service.NewUserService(user)

	return services{
		userService: userService,
	}
}
