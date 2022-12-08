package handler

import (
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

type respResult struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

func success(c *gin.Context, resp *respResult) {
	resp.Code = respOk
	resp.Msg = "success"

	c.JSON(http.StatusOK, resp)
}

func unLogin(c *gin.Context) {
	c.JSON(http.StatusOK, respResult{
		Code: respUnLogin,
		Msg:  "unLogin",
	})
}

// fail
func fail(c *gin.Context, e error) {
	// logError(e)
	c.JSON(http.StatusOK, respResult{
		Code: respFail,
		Msg:  e.Error(),
	})
}

func logError(e error) {
	log.Printf("error: 【full】 %+#v ", e)
}
