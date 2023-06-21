package routers

import (
	database "RESTful-API-Test-Joe_Allen_Butarbutar/configs"
	"RESTful-API-Test-Joe_Allen_Butarbutar/controllers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	route := gin.Default()
	api := &controllers.APIEnv{
		DB: database.GetDB(),
	}

	r := route.Group("/users")
	{
		r.GET("", api.GetUsers)
		r.POST("", api.CreateUser)
		r.GET("/:id", api.GetUserByID)
		r.PUT("/:id", api.UpdateUser)
		r.DELETE("/:id", api.DeleteUser)
	}

	return route
}