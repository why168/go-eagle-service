package service

import (
	"go-eagle-service/internal/app/model"
	"go-eagle-service/internal/app/service/user"
	"go-eagle-service/internal/app/service/user/impl"
)

var (
	UserRepository user.Repository
)

//Init instantiate the service
func Init() {
	UserRepository = impl.NewMysqlImpl(model.MysqlHandler)
}
