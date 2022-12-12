package blockchain

import (
	"context"
	"errors"
	"log"
	"math/big"

	"github.com/doge-verse/easy-upgrade-backend/internal/blockchain/abi"
	"github.com/doge-verse/easy-upgrade-backend/internal/conf"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	EthMainnet     uint = 1
	PolygonMainnet uint = 137
	GoerilTestNet  uint = 5
)

var (
	ownershipEvent         = "OwnershipTransferred(address,address)"
	ownershipEventTopicLen = 3
)

var (
	ethClient     *ethclient.Client
	polygonClient *ethclient.Client
	goerliClient  *ethclient.Client
)

func Init() {
	var err error
	// create client
	ethClient, err = ethclient.Dial(conf.GetRPC().EthMainnt)
	if err != nil {
		log.Fatalln(err)
		return
	}
	polygonClient, err = ethclient.Dial(conf.GetRPC().PolygoMainnet)
	if err != nil {
		log.Fatalln(err)
		return
	}
	goerliClient, err = ethclient.Dial(conf.GetRPC().GoerliTestnet)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func GetOwnershipTransferredEvent(addr string, network uint) ([]models.ContractHistory, error) {
	var client *ethclient.Client
	switch network {
	case EthMainnet:
		client = ethClient
	case PolygonMainnet:
		client = polygonClient
	case GoerilTestNet:
		client = goerliClient
	}

	number, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(addr)
	// hash event
	eventSignature := []byte(ownershipEvent)
	topicHash := crypto.Keccak256Hash(eventSignature)

	fromBlock := big.NewInt(int64(number) - 1000)
	toBlock := big.NewInt(int64(number))
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{
			{topicHash},
		},
	})
	if err != nil {
		return nil, err
	}
	var results []models.ContractHistory
	for _, eventLog := range logs {
		var topics []string
		for i := range eventLog.Topics {
			topics = append(topics, eventLog.Topics[i].Hex())
		}

		if topics[0] == topicHash.Hex() && len(topics) == ownershipEventTopicLen {
			var blockTime uint64
			blockInfo, err := client.BlockByNumber(context.Background(), big.NewInt(int64(eventLog.BlockNumber)))
			if err == nil {
				blockTime = blockInfo.Time()
			}
			results = append(results, models.ContractHistory{
				UpdateBlock:   uint(eventLog.BlockNumber),
				Network:       network,
				UpdateTime:    blockTime,
				UpdateTX:      eventLog.TxHash.Hex(),
				PreviousOwner: common.HexToAddress(topics[1]).String(),
				NewOwner:      common.HexToAddress(topics[2]).String(),
			})
		}
	}
	return results, errors.New("not found")
}

func GetProxyOwner(addr string, network uint) (string, error) {
	var client *ethclient.Client
	switch network {
	case EthMainnet:
		client = ethClient
	case PolygonMainnet:
		client = polygonClient
	case GoerilTestNet:
		client = goerliClient
	}

	contractAddr := common.HexToAddress(addr)

	instance, err := abi.NewProxyAdmin(contractAddr, client)
	if err != nil {
		return "", err
	}

	owner, err := instance.Owner(&bind.CallOpts{})
	if err != nil {
		return "", err
	}

	return owner.String(), nil
}
