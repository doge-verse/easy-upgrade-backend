package contract

import (
	"github.com/doge-verse/easy-upgrade-backend/api/request"
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

// PageUserContractArr .
func (repo sqlRepo) PageUserContractArr(userID uint, pageInfo request.PageInfo) ([]models.Contract, int64, error) {
	var contractArr []models.Contract

	query := CQuery{
		UserID: userID,
	}

	db := repo.db.Session(&gorm.Session{}).Model(&models.Contract{}).Scopes(query.where())
	total, err := count(db)
	if err != nil {
		return nil, 0, err
	}

	if err = db.Scopes(models.Paginate(pageInfo)).Order("id desc").Find(&contractArr).Error; err != nil {
		return nil, 0, err
	}

	return contractArr, total, nil
}

func (repo sqlRepo) PageContractHistory(addr string, pageInfo request.PageInfo) ([]models.ContractHistory, int64, error) {
	var records []models.ContractHistory
	query := CHQuery{
		ContractAddr: addr,
	}
	db := repo.db.Session(&gorm.Session{}).Model(&models.ContractHistory{}).Scopes(query.where())
	total, err := count(db)
	if err != nil {
		return nil, 0, err
	}
	if err = db.Scopes(models.Paginate(pageInfo)).Order("id desc").Find(&records).Error; err != nil {
		return nil, 0, err
	}
	return records, total, nil
}

// TODO: after update need to modify contract some filed

func (repo sqlRepo) AddNotifier(param *models.Notifier) error {
	if err := repo.db.Model(&models.Notifier{}).Create(param).Error; err != nil {
		return err
	}
	// TODO: add get contract update history list

	return nil
}

func count(db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
