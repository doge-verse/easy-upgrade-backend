package subscriber

import (
	"context"
	"fmt"
	"log"
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
				log.Fatal("ethclient dial eth mainnet error:", err)
				return err
			}
		case blockchain.PolygonMainnet:
			client, err = ethclient.Dial(conf.GetRPC().PolygoMainnet)
			if err != nil {
				log.Fatal("ethclient dial polygon mainnet error:", err)
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

		for {
			select {
			case err := <-sub.Err():
				log.Fatal("get log from chan error:", err)
			case currentLog := <-logs:
				oldProxyAddress := common.BytesToAddress(currentLog.Data[:32])
				newProxyAddress := common.BytesToAddress(currentLog.Data[32:])
				sendEmail(receiverEmail, oldProxyAddress, newProxyAddress)
			}
		}
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

	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", config.From, authCode, "smtp.qq.com"))
	if err != nil {
		log.Fatal("Send Email error:", err)
	} else {
		log.Println("Send email successfully!")
	}
}
