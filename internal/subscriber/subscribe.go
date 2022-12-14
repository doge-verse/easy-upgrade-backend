package subscriber

import (
	"context"
	"fmt"
	"math/big"
	"net/smtp"

	"github.com/doge-verse/easy-upgrade-backend/internal/blockchain"
	"github.com/doge-verse/easy-upgrade-backend/internal/conf"
	"github.com/doge-verse/easy-upgrade-backend/internal/shared"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Subscriber struct {
	Db *gorm.DB
}

func (s Subscriber) SelectAllContract() ([]models.Contract, error) {
	var contracts []models.Contract

	if err := s.Db.Select("ProxyAddress", "Network", "Email").Find(&contracts).Error; err != nil {
		logrus.Errorln("Select all contract from db error", err)
		return nil, err
	}

	return contracts, nil
}

func (s Subscriber) SubscribeAllContract(contracts []models.Contract) {
	for _, contract := range contracts {
		go s.SubscribeOneContract(contract)
	}
}

func (s Subscriber) SubscribeOneContract(contract models.Contract) {
	logrus.Infof("%+v", contract)
	contractAddressStr := contract.ProxyAddress
	network := contract.Network
	receiverEmail := contract.Email

	var err error
	client, ok := blockchain.ClientList[network]
	if !ok {
		logrus.Errorln(shared.ErrChainNotInit)
		return
	}

	contractAddress := common.HexToAddress(contractAddressStr)
	topicHash := crypto.Keccak256Hash([]byte(blockchain.OwnershipEvent))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{topicHash}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logrus.Errorln("SubscribeFilterLogs error:", err)
		return
	}
	for {
		select {
		case err := <-sub.Err():
			logrus.Errorln("get log from chan error:", err)
		case currentLog := <-logs:
			oldProxyAddress := common.HexToAddress(currentLog.Topics[1].Hex()).String()
			newProxyAddress := common.HexToAddress(currentLog.Topics[2].Hex()).String()
			historyInfo := models.ContractHistory{
				UpdateBlock:   currentLog.BlockNumber,
				UpdateTX:      currentLog.TxHash.Hex(),
				PreviousOwner: oldProxyAddress,
				NewOwner:      newProxyAddress,
			}
			s.updateContract(contractAddressStr, historyInfo, client)
			sendEmail(receiverEmail, oldProxyAddress, newProxyAddress)
		}
	}
}

func sendEmail(receiverEmail string, oldAdmin string, newAdmin string) {
	config := conf.GetEmailConf()
	authCode := config.AuthCode

	e := email.NewEmail()
	e.From = config.From
	e.To = []string{receiverEmail}
	e.Subject = "Proxy Admin Change"
	e.Text = []byte(fmt.Sprintf("The proxy admin has changed from %s to %s!", oldAdmin, newAdmin))

	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", config.Username, authCode, "smtp.qq.com"))
	if err != nil {
		logrus.Errorln("Send Email error:", err)
	} else {
		logrus.Infoln("Send email successfully!")
	}
}

// updateContract update the contract lastUpdate time & owner & create history record
func (s Subscriber) updateContract(contractAddress string, historyInfo models.ContractHistory, client *ethclient.Client) {
	tx := s.Db.Begin()
	var contract models.Contract
	if err := tx.Model(&models.Contract{}).Where("proxy_address = ?", contractAddress).
		First(&contract).Error; err != nil {
		logrus.Infoln("The contract is not exist in db:", err)
		tx.Rollback()
	}
	historyInfo.ContractID = contract.ID
	var blockTime uint64
	blockInfo, err := client.BlockByNumber(context.Background(), big.NewInt(int64(historyInfo.UpdateBlock)))
	if err == nil {
		blockTime = blockInfo.Time()
	}
	historyInfo.UpdateTime = blockTime
	contract.LastUpdate = blockTime
	contract.ProxyOwner = historyInfo.NewOwner
	if err = tx.Save(&contract).Error; err != nil {
		logrus.Infoln("The contract update fail:", err)
		tx.Rollback()
	}
	if err = tx.Model(&models.ContractHistory{}).Create(&historyInfo).Error; err != nil {
		logrus.Infoln("The contract update history create fail:", err)
		tx.Rollback()
	}
	tx.Commit()
}
