package routers

import (
	db "RestuarantBackend/db"
	handlers "RestuarantBackend/handlers"
	service "RestuarantBackend/service"

	"github.com/gin-gonic/gin"
)

func SetRoutesAPI(r *gin.Engine) {
	db.Connect()
	userService := &service.UserService{}
	userController := handlers.NewUserController(userService)
	v1 := r.Group("api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/signup", userController.Register)
			users.POST("/login", userController.Login)
		}
	}
}
