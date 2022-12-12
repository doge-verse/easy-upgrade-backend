package handler

import (
	"github.com/doge-verse/easy-upgrade-backend/api/middleware"
	"github.com/doge-verse/easy-upgrade-backend/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors())

	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.POST("/api/login", login)
	router.POST("/api/logout", logoutUser)
	router.POST("/api/user/register", registerUser)

	initNeedAuthRouter(router)
}

func initNeedAuthRouter(r *gin.Engine) {
	r.Use(middleware.JWTAuth())
	r.GET("/api/currentUser", currentUser)

	userGroup := r.Group("/api/user")
	{
		// userGroup.GET("/get", getUserByQuery)
		userGroup.POST("/email", updateEmail)
	}
	contractGroup := r.Group("/api/notifier")
	{
		contractGroup.GET("", getUserContract)
		contractGroup.POST("", addContract)
		contractGroup.GET("/history", getContractHistory)
	}
}
