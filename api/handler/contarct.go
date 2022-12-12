package handler

import (
	"fmt"

	"github.com/doge-verse/easy-upgrade-backend/api/request"
	response "github.com/doge-verse/easy-upgrade-backend/api/respone"
	"github.com/doge-verse/easy-upgrade-backend/internal/contract"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/doge-verse/easy-upgrade-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// addContract .
// @Tags notify
// @Summary create notify event
// @accept application/json
// @Produce application/json
// @Param data body request.Contract true "add notifier param"
// @Success 200 {object} respResult{data=models.Contract}
// @Router /notifier [post]
func addContract(c *gin.Context) {
	param := request.Contract{}
	if err := c.ShouldBind(&param); err != nil {
		fail(c, err)
		return
	}
	if len(param.ProxyAddress) < 32 {
		fail(c, fmt.Errorf("address wrong"))
		return
	}
	param.UserID = getUserID(c)
	// FIXME:
	param.UserID = 1
	var contractEntity models.Contract
	_ = copier.Copy(&contractEntity, &param)

	result, err := contract.Repo.AddContract(&contractEntity)
	if err != nil {
		fail(c, err)
		return
	}
	success(c, &respResult{
		Data: result,
	})
}

// getUserContract .
// @Tags notify
// @Summary page query notify event
// @accept application/json
// @Produce application/json
// @Param pageNum query int false "page number"
// @Param pageSize query int false "page size"
// @Success 200 {object} respResult{data=response.PageResult{list=[]models.Contract}}
// @Router /notifier [get]
func getUserContract(c *gin.Context) {
	userID := getUserID(c)
	// FIXME:
	userID = 1
	pageInfo := request.PageInfo{
		PageNum:  models.DefaultPageNum(c.Query("pageNum")),
		PageSize: models.DefaultPageSize(c.Query("pageSize")),
	}
	contractArr, total, err := contract.Repo.PageUserContractArr(userID, pageInfo)
	if err != nil {
		fail(c, err)
		return
	}
	success(c, &respResult{
		Data: response.PageResult{
			List:     contractArr,
			Total:    total,
			PageNum:  pageInfo.PageNum,
			PageSize: pageInfo.PageSize,
		},
	})
}

// getContractHistory .
// @Tags notify
// @Summary page query update history
// @accept application/json
// @Produce application/json
// @Param contractID query int true "contract id"
// @Param pageNum query int false "page number"
// @Param pageSize query int false "page size"
// @Success 200 {object} respResult{data=models.Contract}
// @Router /notifier/history [get]
func getContractHistory(c *gin.Context) {
	contractID, _ := util.ParseUint(c.Query("contractID"))
	if contractID == 0 {
		fail(c, fmt.Errorf("the contract id can not be empty"))
		return
	}
	pageInfo := request.PageInfo{
		PageNum:  models.DefaultPageNum(c.Query("pageNum")),
		PageSize: models.DefaultPageSize(c.Query("pageSize")),
	}

	records, total, err := contract.Repo.PageContractHistory(contractID, pageInfo)
	if err != nil {
		fail(c, err)
		return
	}
	success(c, &respResult{
		Data: response.PageResult{
			List:     records,
			Total:    total,
			PageNum:  pageInfo.PageNum,
			PageSize: pageInfo.PageSize,
		},
	})
}
