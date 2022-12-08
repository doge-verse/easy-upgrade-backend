package contract

import "github.com/doge-verse/easy-upgrade-backend/models"

type repoI interface {
	AddContract(contract *models.Contract) (*models.Contract, error)

	GetUserContractArr(userID uint) ([]models.Contract, error)

	GetContractHistory(addr string) ([]models.ContractHistory, error)

	AddNotifier(param *models.Notifier) error
}
