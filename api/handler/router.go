package handler

import (
	"github.com/doge-verse/easy-upgrade-backend/api/middleware"
	"github.com/doge-verse/easy-upgrade-backend/internal/conf"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// InitRouter .
func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors())

	router.Static("/api/docs", "./docs")

	router.Use(sessions.Sessions("easy-upgrade", conf.GetSessionStore()))

	router.POST("/api/login", login)
	router.POST("/api/logout", logoutUser)
	router.POST("/api/user/register", registerUser)

	initNeedAuthRouter(router)
}

func initNeedAuthRouter(r *gin.Engine) {
	r.Use(auth)
	r.GET("/api/currentUser", currentUser)

	userGroup := r.Group("/api/user")
	{
		// userGroup.GET("/get", getUserByQuery)
		userGroup.POST("/email", updateEmail)
	}
	contractGroup := r.Group("/api/contract")
	{
		contractGroup.GET("", getUserContract)
		contractGroup.POST("", addContract)
		contractGroup.GET("/history", getContractHistory)
		contractGroup.POST("/notifier", addNotifier)
	}
}
