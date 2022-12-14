package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	respOk      = 0 // "OK"
	respNotUser = 1 // "NotUser"
	respFail    = 2 // "FAIL"
	respUnLogin = 3 // "UnLogin"
	respNoAuth  = 4 // "NoAuth"
)

type RespResult struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

func Success(c *gin.Context, resp *RespResult) {
	resp.Code = respOk
	resp.Msg = "success"

	c.JSON(http.StatusOK, resp)
}

func UnLogin(c *gin.Context) {
	c.JSON(http.StatusOK, RespResult{
		Code: respUnLogin,
		Msg:  "UnLogin",
	})
}

// Fail .
func Fail(c *gin.Context, e error) {
	// LogError(e)
	c.JSON(http.StatusOK, RespResult{
		Code: respFail,
		Msg:  e.Error(),
	})
}

func LogError(e error) {
	logrus.Infof("error: 【full】 %+#v ", e)
}
