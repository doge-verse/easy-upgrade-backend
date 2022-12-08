package handler

import (
	"fmt"
	"log"

	"github.com/doge-verse/easy-upgrade-backend/internal/blockchain"
	"github.com/doge-verse/easy-upgrade-backend/internal/shared"
	"github.com/doge-verse/easy-upgrade-backend/internal/user"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/spf13/cast"

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
		unLogin(c)
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
		success(c, nil)
		return
	}
	session := sessions.Default(c)
	err := session.Save()
	if err != nil {
		log.Println(err)
	}
	success(c, resp{
		"data": userInfo,
	})
}

// currentLoginUser .
func currentLoginUser(c *gin.Context) {
	session := sessions.Default(c)
	t := session.Get("userID")
	if t == nil {
		success(c, resp{
			"data": false,
		})
		return
	}
	userID := t.(uint)
	userInfo, err := user.Repo.GetUserByQuery(user.Query{UserID: userID})
	if userInfo == nil || err != nil {
		success(c, nil)
		return
	}
	err = session.Save()
	if err != nil {
		log.Println(err)
	}
	success(c, resp{
		"data":  userInfo,
		"login": true,
	})
}

// login .
func login(c *gin.Context) {
	param := &struct {
		Address   string `json:"address" form:"address"`
		Signature string `json:"signature" form:"signature"`
		SignData  string `json:"signData" form:"signData"`
	}{}
	if err := c.ShouldBind(param); err != nil {
		fail(c, fmt.Errorf("param error"))
		return
	}
	if !blockchain.CheckAddr(param.Address, param.Signature, param.SignData) {
		fail(c, fmt.Errorf("signature fail"))
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
			fail(c, err)
			return
		}
	}
	session := sessions.Default(c)
	session.Set("userID", userInfo.ID)
	if err = session.Save(); err != nil {
		log.Println(err)
	}
	success(c, resp{
		"data": userInfo,
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
	success(c, nil)
}

func getUserID(c *gin.Context) int64 {
	return cast.ToInt64(sessions.Default(c).Get("userID"))
}
