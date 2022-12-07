package handler

import (
	"fmt"

	"github.com/doge-verse/easy-upgrade-backend/internal/contract"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/doge-verse/easy-upgrade-backend/util"
	"github.com/gin-gonic/gin"
)

func addContract(c *gin.Context) {
	param := models.Contract{}
	if err := c.ShouldBindQuery(&param); err != nil {
		fail(c, err)
		return
	}
	if len(param.ContractAddr) < 32 {
		fail(c, fmt.Errorf("address wrong"))
		return
	}
	result, err := contract.Repo.AddContract(&param)
	if err != nil {
		fail(c, err)
		return
	}
	success(c, resp{
		"data": result,
	})
}

func getUserContract(c *gin.Context) {
	userID, _ := util.ParseUint(c.Query("userID"))
	if userID == 0 {
		fail(c, fmt.Errorf("have to get userID"))
		return
	}
	contractArr, err := contract.Repo.GetUserContractArr(userID)
	if err != nil {
		fail(c, err)
		return
	}
	okArr(c, resp{
		"data": contractArr,
	})
}

func getContractHistory(c *gin.Context) {
	addr := c.Query("addr")
	if addr == "" {
		fail(c, fmt.Errorf("the addr can not be empty"))
		return
	}
	records, err := contract.Repo.GetContractRecord(addr)
	if err != nil {
		fail(c, err)
		return
	}
	okArr(c, resp{
		"data": records,
	})
}

func addNotifier(c *gin.Context) {
	param := models.Notifier{}
	if err := c.ShouldBindQuery(&param); err != nil {
		fail(c, err)
		return
	}
	param.UserID = getUserID(c)
	if len(param.ContractAddr) < 32 {
		fail(c, fmt.Errorf("address wrong"))
		return
	}
	err := contract.Repo.AddNotifier(&param)
	if err != nil {
		fail(c, err)
		return
	}
	success(c, resp{
		"data": nil,
	})
}
