package handler

import (
	"github.com/doge-verse/easy-upgrade-backend/api/request"
	"github.com/doge-verse/easy-upgrade-backend/api/response"
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
		response.Fail(c, err)
		return
	}
	query := user.Query{
		UserID:  userID,
		Address: address,
	}
	userInfo, err := user.Repo.GetUserByQuery(query)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, &response.RespResult{
		Data: userInfo,
	})
}

// updateEmail .
// @Tags auth
// @Summary update user email
// @accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param Authorization header string false "token"
// @Param data body request.UpdateEmail true "update user email"
// @Success 200 {object} response.RespResult
// @Router /user/email [post]
func updateEmail(c *gin.Context) {
	param := request.UpdateEmail{}
	if err := c.ShouldBindQuery(&param); err != nil {
		response.Fail(c, err)
		return
	}
	param.UserID = getUserID(c)
	if err := user.Repo.UpdateEmail(param.UserID, param.Email); err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, &response.RespResult{
		Data: nil,
	})
}

// registerUser
func registerUser(c *gin.Context) {
	param := models.User{}
	if err := c.ShouldBindQuery(&param); err != nil {
		response.Fail(c, err)
		return
	}
	userInfo, err := user.Repo.UserRegister(&param)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, &response.RespResult{
		Data: userInfo,
	})
}
