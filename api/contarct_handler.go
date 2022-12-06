package api

import (
	"fmt"

	"github.com/doge-verse/easy-upgrade-backend/internal/contract"
	"github.com/doge-verse/easy-upgrade-backend/pkg"
	"github.com/doge-verse/easy-upgrade-backend/util"

	"github.com/gin-gonic/gin"
)

func addContract(c *gin.Context) {
	param := pkg.Contract{}
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
	ok(c, resp{
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

func getContractRecord(c *gin.Context) {
	addr := c.Query("addr")
	if addr == "" {
		fail(c, fmt.Errorf("The addr can not be empty."))
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
