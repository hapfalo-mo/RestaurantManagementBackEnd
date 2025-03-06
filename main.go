package main

import (
	"RestuarantBackend/db"
	"RestuarantBackend/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	// Connect Database
	db.Connect()
	defer db.DB.Close()

	// Initialize Router
	router := gin.Default()

	// // CORS Middleware
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8080"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// }))

	// Register User API Routes
	routers.SetRoutesAPI(router)

	// Run Server
	router.Run(":1611")
}
