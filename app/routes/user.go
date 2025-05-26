package routes

import (
	"FP-DevOps/config"
	"FP-DevOps/controller"
	"FP-DevOps/middleware"

	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService config.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("/login", userController.Login)
		routes.POST("/register", userController.Register)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
	}
}
