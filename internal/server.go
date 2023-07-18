package internal

import (
	"github.com/gavrishp/coreapiservicetest/internal/common/db"
	"github.com/gavrishp/coreapiservicetest/internal/controllers"
	"github.com/gavrishp/coreapiservicetest/internal/repository"
	"github.com/gavrishp/coreapiservicetest/internal/router"
	"github.com/gavrishp/coreapiservicetest/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func MakeServer() *gin.Engine {
	viper.SetConfigFile("./internal/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	db := db.Init(dbUrl)

	usersRepository := repository.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepository)
	usersController := controllers.NewUsersController(usersService)

	router.RegisterRoutes(r, usersController)

	r.Run(port)
	return r
}
