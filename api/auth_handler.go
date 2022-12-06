package api

import (
	"fmt"
	"log"

	"github.com/doge-verse/easy-upgrade-backend/internal/shared"
	"github.com/doge-verse/easy-upgrade-backend/internal/user"

	"github.com/doge-verse/easy-upgrade-backend/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getUserIDFromSession(c *gin.Context) uint {
	session := sessions.Default(c)
	t := session.Get("userID")
	if t == nil {
		unLogin(c)
		// c.Abort()
		return 0
	}
	return t.(uint)
}

// auth .
func auth(c *gin.Context) {
	userID := getUserIDFromSession(c)
	userInfo, err := user.Repo.GetUserByQuery(user.Query{UserID: userID})
	if err != nil {
		log.Println("err", err)
		// unLogin(c)
		// c.Abort()
	}
	ctx := shared.WithUser(c.Request.Context(), userInfo)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

// currentUser .
func currentUser(c *gin.Context) {
	userInfo, exist := shared.GetUser(c.Request.Context())
	if !exist {
		ok(c, nil)
		return
	}
	session := sessions.Default(c)
	err := session.Save()
	if err != nil {
		log.Println(err)
	}
	ok(c, resp{
		"data": userInfo,
	})
}

// currentLoginUser .
func currentLoginUser(c *gin.Context) {
	session := sessions.Default(c)
	t := session.Get("userID")
	if t == nil {
		ok(c, resp{
			"data": false,
		})
		return
	}
	userID := t.(uint)
	userInfo, err := user.Repo.GetUserByQuery(user.Query{UserID: userID})
	if userInfo == nil || err != nil {
		ok(c, nil)
		return
	}
	err = session.Save()
	if err != nil {
		log.Println(err)
	}
	ok(c, resp{
		"data":  userInfo,
		"login": true,
	})
}

// login .
func login(c *gin.Context) {
	param := &struct {
		Address   string `json:"address" form:"address"`
		Signature string `json:"signature" form:"signature"`
	}{}
	if err := c.ShouldBindQuery(param); err != nil {
		fail(c, fmt.Errorf("login fail"))
		return
	}
	userInfo, err := user.Repo.GetUserByQuery(user.Query{
		Address: param.Address,
	})
	if err != nil {
		fail(c, err)
		return
	}
	token, err := util.Sign(userInfo.ID)
	if err != nil {
		fail(c, err)
		return
	}
	session := sessions.Default(c)
	session.Set("userID", userInfo.ID)
	err = session.Save()
	if err != nil {
		log.Println(err)
	}
	ok(c, resp{
		"data":  userInfo,
		"token": token,
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
	ok(c, nil)
}
