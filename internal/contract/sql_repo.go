package contract

import (
	"github.com/doge-verse/easy-upgrade-backend/pkg"

	"gorm.io/gorm"
)

type sqlRepo struct {
	db *gorm.DB
}

// AddContract .
func (repo sqlRepo) AddContract(contract *pkg.Contract) (*pkg.Contract, error) {
	if err := repo.db.Model(&pkg.Contract{}).Create(contract).Error; err != nil {
		return nil, err
	}
	return contract, nil
}

// GetUserContractArr .
func (repo sqlRepo) GetUserContractArr(userID uint) ([]pkg.Contract, error) {
	var contractArr []pkg.Contract
	if err := repo.db.Model(&pkg.Contract{}).Where("user_id = ?", userID).Find(&contractArr).Error; err != nil {
		return nil, err
	}
	return contractArr, nil
}

func (repo sqlRepo) GetContractRecord(addr string) ([]pkg.ContractUpdateRecord, error) {
	var records []pkg.ContractUpdateRecord
	if err := repo.db.Model(&pkg.ContractUpdateRecord{}).Where("contract_addr = ?", addr).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}
