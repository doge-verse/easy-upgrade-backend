package contract

import (
	"github.com/doge-verse/easy-upgrade-backend/models"

	"gorm.io/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

// AddContract .
func (repo sqlRepo) AddContract(contract *models.Contract) (*models.Contract, error) {
	if err := repo.db.Model(&models.Contract{}).Create(contract).Error; err != nil {
		return nil, err
	}
	return contract, nil
}

// GetUserContractArr .
func (repo sqlRepo) GetUserContractArr(userID uint) ([]models.Contract, error) {
	var contractArr []models.Contract
	if err := repo.db.Model(&models.Contract{}).Where("user_id = ?", userID).Find(&contractArr).Error; err != nil {
		return nil, err
	}
	return contractArr, nil
}

func (repo sqlRepo) GetContractRecord(addr string) ([]models.ContractHistory, error) {
	var records []models.ContractHistory
	if err := repo.db.Model(&models.ContractHistory{}).Where("contract_addr = ?", addr).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (repo sqlRepo) AddNotifier(param *models.Notifier) error {
	if err := repo.db.Model(&models.Notifier{}).Create(param).Error; err != nil {
		return err
	}
	// TODO: add get contract update history list
	return nil
}
