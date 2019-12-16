package main

import (
	"github.com/gin-gonic/gin"

	"github.com/cildhdi/In-charge/auth"
	_ "github.com/cildhdi/In-charge/models"
	"github.com/cildhdi/In-charge/router"
	api "github.com/cildhdi/In-charge/router/api"
	"github.com/cildhdi/In-charge/utils"
)

func main() {
	mainRouter := gin.Default()

	mainRouter.GET("/status", router.Status)
	apiGroup := mainRouter.Group("/api")

	apiGroup.POST("/login", api.Login)
	apiGroup.POST("/admin-register", api.SuperRegister)
	apiGroup.POST("/register", api.Register)
	apiGroup.POST("/send-code", api.SendVerificationCode)

	authMiddleware := auth.GetMiddleware()
	authGroup := apiGroup.Group("/auth")
	authGroup.Use(authMiddleware.MiddlewareFunc())
	authGroup.GET("/reachable", func(ctx *gin.Context) {
		utils.Success(ctx, nil)
	})

	authGroup.GET("/unreachable", func(ctx *gin.Context) {
		utils.Success(ctx, nil)
	})

	mainRouter.Run()
}
