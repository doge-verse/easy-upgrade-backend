package contract

import (
	"github.com/doge-verse/easy-upgrade-backend/api/request"
	"github.com/doge-verse/easy-upgrade-backend/models"
)

type repoI interface {
	AddContract(contract *models.Contract) (*models.Contract, error)

	PageUserContractArr(userID uint, pageInfo request.PageInfo) ([]models.Contract, int64, error)

	PageContractHistory(addr string, pageInfo request.PageInfo) ([]models.ContractHistory, int64, error)
}
