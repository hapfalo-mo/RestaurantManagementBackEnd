package routers

import (
	db "RestuarantBackend/db"
	handlers "RestuarantBackend/handlers"
	middleware "RestuarantBackend/middleware"
	service "RestuarantBackend/service"

	"github.com/gin-gonic/gin"
)

func SetRoutesAPI(r *gin.Engine) {
	db.Connect()
	userService := &service.UserService{}
	bookingService := &service.BookingService{}
	bookingController := handlers.NewBookingController(bookingService)
	userController := handlers.NewUserController(userService)
	v1 := r.Group("api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/signup", userController.Register)
			users.POST("/login", userController.LoginToken)
			users.PUT("/updateUser", middleware.AuthenticateMiddleware, userController.Update)
		}

		bookings := v1.Group("/bookings")
		{
			bookings.POST("/bookTable", middleware.AuthenticateMiddleware, bookingController.BookingTable)
			bookings.GET("/getBooking/:id", middleware.AuthenticateMiddleware, bookingController.PagingBookingList)
		}
	}
}
