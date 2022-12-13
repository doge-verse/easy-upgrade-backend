package contract

import (
	"github.com/doge-verse/easy-upgrade-backend/api/request"
	"github.com/doge-verse/easy-upgrade-backend/internal/blockchain"
	"github.com/doge-verse/easy-upgrade-backend/internal/subscriber"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/ethereum/go-ethereum/ethclient"

	"gorm.io/gorm"
)

type service struct {
	db               *gorm.DB
	ethClient        *ethclient.Client
	polygonClient    *ethclient.Client
	goerliClient     *ethclient.Client
	fVMWallabyClient *ethclient.Client
}

// AddContract .
func (repo service) AddContract(contract *models.Contract) (*models.Contract, error) {
	ownerAddr, err := blockchain.GetProxyOwner(contract.ProxyAddress, contract.Network)
	if err != nil {
		return nil, err
	}
	contract.ProxyOwner = ownerAddr
	if err := repo.db.
		Model(&models.Contract{}).
		Create(contract).Error; err != nil {
		return nil, err
	}

	s := &subscriber.Subscriber{
		Db:                   repo.db,
		EthMainnetClient:     repo.ethClient,
		PolygonMainnetClient: repo.polygonClient,
		GoerliClinet:         repo.goerliClient,
	}

	go s.SubscribeOneContract(*contract)

	go func() {
		historyList, err := blockchain.GetOwnershipTransferredEvent(contract.ProxyAddress, contract.Network)
		if err == nil && len(historyList) > 0 {
			for k, v := range historyList {
				historyList[k].ContractID = contract.ID
				if v.UpdateTime > contract.LastUpdate {
					contract.LastUpdate = v.UpdateTime
				}
			}
			// insert contract history
			repo.db.Model(&models.ContractHistory{}).CreateInBatches(historyList, len(historyList))
			// update contract newest update time
			repo.db.Model(&models.Contract{}).Updates(contract)
		}
	}()
	return contract, nil
}

// PageUserContractArr .
func (repo service) PageUserContractArr(userID uint, pageInfo request.PageInfo) ([]models.Contract, int64, error) {
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

func (repo service) PageContractHistory(contractID uint, pageInfo request.PageInfo) ([]models.ContractHistory, int64, error) {
	var records []models.ContractHistory
	query := CHQuery{
		ContractID: contractID,
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

func count(db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
