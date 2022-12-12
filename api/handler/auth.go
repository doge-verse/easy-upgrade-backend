package handler

import (
	"fmt"
	"log"

	"github.com/doge-verse/easy-upgrade-backend/api/middleware"
	"github.com/doge-verse/easy-upgrade-backend/api/request"
	"github.com/doge-verse/easy-upgrade-backend/api/response"
	"github.com/doge-verse/easy-upgrade-backend/internal/blockchain"
	"github.com/doge-verse/easy-upgrade-backend/internal/shared"
	"github.com/doge-verse/easy-upgrade-backend/internal/user"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getUserIDFromSession(c *gin.Context) uint {
	session := sessions.Default(c)
	t := session.Get("userID")
	if t == nil {
		return 0
	}
	return t.(uint)
}

// auth .
func auth(c *gin.Context) {
	userID := getUserIDFromSession(c)
	userInfo, err := user.Repo.GetUserByQuery(user.Query{UserID: userID})
	if err != nil || userID == 0 {
		response.UnLogin(c)
		c.Abort()
		return
	}
	ctx := shared.WithUser(c.Request.Context(), userInfo)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

// currentUser .
func currentUser(c *gin.Context) {
	userInfo, exist := shared.GetUser(c.Request.Context())
	if !exist {
		response.Success(c, nil)
		return
	}
	session := sessions.Default(c)
	err := session.Save()
	if err != nil {
		log.Println(err)
	}
	response.Success(c, &response.RespResult{
		Data: userInfo,
	})
}

// currentLoginUser .
func currentLoginUser(c *gin.Context) {
	session := sessions.Default(c)
	t := session.Get("userID")
	if t == nil {
		response.Success(c, &response.RespResult{
			Data: false,
		})
		return
	}
	userID := t.(uint)
	userInfo, err := user.Repo.GetUserByQuery(user.Query{UserID: userID})
	if userInfo == nil || err != nil {
		response.Success(c, nil)
		return
	}
	err = session.Save()
	if err != nil {
		log.Println(err)
	}
	response.Success(c, &response.RespResult{
		Data: userInfo,
	})
}

// login .
// @Tags auth
// @Summary User login
// @accept application/json
// @Produce application/json
// @Param data body request.Login true "login param"
// @Success 200 {object} response.RespResult{data=response.UserInfo}
// @Router /login [post]
func login(c *gin.Context) {
	var param request.Login
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, fmt.Errorf("param error"))
		return
	}
	if !blockchain.CheckAddr(param.Address, param.Signature, param.SignData) {
		response.Fail(c, fmt.Errorf("signature fail"))
		return
	}
	userInfo, err := user.Repo.GetUserByQuery(user.Query{
		Address: param.Address,
	})
	if err != nil {
		newUser := &models.User{
			Address: param.Address,
		}
		userInfo, err = user.Repo.UserRegister(newUser)
		if err != nil {
			response.Fail(c, err)
			return
		}
	}
	tokenStr, expiresAt, err := middleware.SignJwt(userInfo.ID)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, &response.RespResult{
		Data: response.UserInfo{
			ID:        userInfo.ID,
			Name:      userInfo.Name,
			Address:   userInfo.Address,
			Email:     userInfo.Email,
			Token:     tokenStr,
			ExpiresAt: expiresAt,
		},
	})
}

// logout .
func logoutUser(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userID")
	err := session.Save()
	if err != nil {
		log.Println(err)
	}
	response.Success(c, &response.RespResult{
		Data: nil,
	})
}

func getUserID(c *gin.Context) uint {
	claims, _ := c.Get("claims")
	return claims.(*request.CustomClaims).UserID
}
