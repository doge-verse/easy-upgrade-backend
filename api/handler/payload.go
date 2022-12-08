package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	respOk      = 0 // "OK"
	respNotUser = 1 // "NotUser"
	respFail    = 2 // "FAIL"
	respUnLogin = 3 // "UnLogin"
	respNoAuth  = 4 // "NoAuth"
)

// Resp .
type resp map[string]interface{}

func success(c *gin.Context, resp resp) {
	result := make(map[string]interface{})
	if resp != nil {
		for key, value := range resp {
			if fmt.Sprint(value) != "<nil>" {
				result[key] = value
			}
		}
	}
	result["code"] = respOk
	result["msg"] = "Success"

	c.JSON(http.StatusOK, result)
}

func unLogin(c *gin.Context) {
	c.JSON(http.StatusOK, resp{
		"code": respUnLogin,
		"msg":  "unLogin",
	})
}

// fail
func fail(c *gin.Context, e error) {
	// logError(e)
	c.JSON(http.StatusOK, resp{
		"code": respFail,
		"msg":  e.Error(),
	})
}

func logError(e error) {
	log.Printf("error: 【full】 %+#v ", e)
}
