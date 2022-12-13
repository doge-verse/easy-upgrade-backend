package subscriber

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/smtp"

	"github.com/doge-verse/easy-upgrade-backend/internal/blockchain"
	"github.com/doge-verse/easy-upgrade-backend/internal/conf"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jordan-wright/email"

	"gorm.io/gorm"
)

type subscriber struct {
	db *gorm.DB
}

func (s subscriber) SelectAllContract() ([]models.Contract, error) {
	var contracts []models.Contract

	if err := s.db.Select("ProxyAddress", "Network", "Email").Find(&contracts).Error; err != nil {
		log.Fatal("Select all contract from db error", err)
		return nil, err
	}

	return contracts, nil
}

func (s subscriber) SubscribeAllContract(contracts []models.Contract) error {
	for _, contract := range contracts {
		contractAddressStr := contract.ProxyAddress
		network := contract.Network
		receiverEmail := contract.Email

		var client *ethclient.Client
		var err error
		switch network {
		case blockchain.EthMainnet:
			client, err = ethclient.Dial(conf.GetRPC().EthMainnt)
			if err != nil {
				log.Fatal("eth client dial eth mainnet error:", err)
				return err
			}
		case blockchain.PolygonMainnet:
			client, err = ethclient.Dial(conf.GetRPC().PolygoMainnet)
			if err != nil {
				log.Fatal("polygon client dial polygon mainnet error:", err)
				return err
			}
		case blockchain.GoerilTestNet:
			client, err = ethclient.Dial(conf.GetRPC().GoerliTestnet)
			if err != nil {
				log.Fatal("goeril client dial goerli testnet error:", err)
				return err
			}
		}

		contractAddress := common.HexToAddress(contractAddressStr)
		topicHash := crypto.Keccak256Hash([]byte("OwnershipTransferred(address,address)"))
		log.Printf("%s", topicHash.String())
		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddress},
			Topics:    [][]common.Hash{{topicHash}},
		}

		logs := make(chan types.Log)
		sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			log.Fatal("SubscribeFilterLogs error:", err)
		}

		go func(ethereum.Subscription, chan types.Log, string, *ethclient.Client) {
			for {
				select {
				case err := <-sub.Err():
					log.Fatal("get log from chan error:", err)
				case currentLog := <-logs:
					oldProxyAddress := common.BytesToAddress(currentLog.Data[:32])
					newProxyAddress := common.BytesToAddress(currentLog.Data[32:])
					historyInfo := models.ContractHistory{
						UpdateBlock:   currentLog.BlockNumber,
						UpdateTX:      currentLog.TxHash.Hex(),
						PreviousOwner: oldProxyAddress.String(),
						NewOwner:      newProxyAddress.String(),
					}
					s.updateContract(contractAddressStr, historyInfo, client)
					sendEmail(receiverEmail, oldProxyAddress, newProxyAddress)
				}
			}
		}(sub, logs, contractAddressStr, client)
	}
	return nil
}

func sendEmail(receiverEmail string, oldAdmin common.Address, newAdmin common.Address) {
	config := conf.GetEmailConf()
	authCode := config.AuthCode

	e := email.NewEmail()
	e.From = config.From
	e.To = []string{receiverEmail}
	e.Subject = "Proxy Admin Change"
	e.Text = []byte(fmt.Sprintf("The proxy admin has changed from %s to %s!", oldAdmin.String(), newAdmin.String()))

	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", config.Username, authCode, "smtp.qq.com"))
	if err != nil {
		log.Fatal("Send Email error:", err)
	} else {
		log.Println("Send email successfully!")
	}
}

// updateContract update the contract lastUpdate time & owner & create history record
func (s subscriber) updateContract(contractAddress string, historyInfo models.ContractHistory, client *ethclient.Client) {
	tx := s.db.Begin()
	var contract models.Contract
	if err := tx.Model(&models.Contract{}).Where("proxy_address = ?", contractAddress).
		First(&contract).Error; err != nil {
		log.Println("The contract is not exist in db:", err)
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
		log.Println("The contract update fail:", err)
		tx.Rollback()
	}
	if err = tx.Model(&models.ContractHistory{}).Create(&historyInfo).Error; err != nil {
		log.Println("The contract update history create fail:", err)
		tx.Rollback()
	}
	tx.Commit()
}
