package handler

import (
	"fmt"

	"github.com/doge-verse/easy-upgrade-backend/api/request"
	response "github.com/doge-verse/easy-upgrade-backend/api/respone"
	"github.com/doge-verse/easy-upgrade-backend/internal/contract"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/doge-verse/easy-upgrade-backend/util"
	"github.com/gin-gonic/gin"
)

func addContract(c *gin.Context) {
	param := models.Contract{}
	if err := c.ShouldBind(&param); err != nil {
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
	pageInfo := request.PageInfo{
		PageNum:  models.DefaultPageNum(c.Query("pageNum")),
		PageSize: models.DefaultPageSize(c.Query("pageSize")),
	}
	contractArr, total, err := contract.Repo.PageUserContractArr(userID, pageInfo)
	if err != nil {
		fail(c, err)
		return
	}
	success(c, resp{
		"data": response.PageResult{
			List:     contractArr,
			Total:    total,
			PageNum:  pageInfo.PageNum,
			PageSize: pageInfo.PageSize,
		},
	})
}

func getContractHistory(c *gin.Context) {
	addr := c.Query("addr")
	if addr == "" {
		fail(c, fmt.Errorf("the addr can not be empty"))
		return
	}
	pageInfo := request.PageInfo{
		PageNum:  models.DefaultPageNum(c.Query("pageNum")),
		PageSize: models.DefaultPageSize(c.Query("pageSize")),
	}

	records, total, err := contract.Repo.PageContractHistory(addr, pageInfo)
	if err != nil {
		fail(c, err)
		return
	}
	success(c, resp{
		"data": response.PageResult{
			List:     records,
			Total:    total,
			PageNum:  pageInfo.PageNum,
			PageSize: pageInfo.PageSize,
		},
	})
}
