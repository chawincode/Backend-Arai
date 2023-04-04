package routes

import (
	controller "exmaple/Backendasktu/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func AuthRoutes(router *gin.Engine) {
	router.POST("/auth/signup", controller.SignUp())
	router.POST("/auth/login", controller.Login())
}
