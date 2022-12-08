package handler

import (
	"github.com/doge-verse/easy-upgrade-backend/internal/user"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/doge-verse/easy-upgrade-backend/util"

	"github.com/gin-gonic/gin"
)

// getUserByQuery
func getUserByQuery(c *gin.Context) {
	userID, err := util.ParseUint(c.Query("userID"))
	address := c.Query("address")
	if err != nil {
		fail(c, err)
		return
	}
	query := user.Query{
		UserID:  userID,
		Address: address,
	}
	userInfo, err := user.Repo.GetUserByQuery(query)
	if err != nil {
		fail(c, err)
		return
	}
	success(c, &respResult{
		Data: userInfo,
	})
}

func updateEmail(c *gin.Context) {
	param := struct {
		Email  string `json:"email"`
		UserID uint   `json:"userID"`
	}{}
	if err := c.ShouldBindQuery(&param); err != nil {
		fail(c, err)
		return
	}
	if err := user.Repo.UpdateEmail(param.UserID, param.Email); err != nil {
		fail(c, err)
		return
	}
	success(c, &respResult{
		Data: nil,
	})
}

// registerUser
func registerUser(c *gin.Context) {
	param := models.User{}
	if err := c.ShouldBindQuery(&param); err != nil {
		fail(c, err)
		return
	}
	userInfo, err := user.Repo.UserRegister(&param)
	if err != nil {
		fail(c, err)
		return
	}
	success(c, &respResult{
		Data: userInfo,
	})
}
