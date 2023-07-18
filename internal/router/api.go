package router

import (
	"github.com/gavrishp/coreapiservicetest/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, usersController *controllers.UsersController) {

	routes := r.Group("/api/v1/users")
	routes.POST("/", usersController.AddUser)
	routes.GET("/", usersController.GetUsers)
	routes.GET("/id", usersController.GetUser)
	routes.PUT("/id", usersController.UpdateUser)
	routes.DELETE("/id", usersController.DeleteUser)
}
