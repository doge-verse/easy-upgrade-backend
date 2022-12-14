package blockchain

import (
	"context"
	"errors"
	"log"
	"math/big"

	"github.com/doge-verse/easy-upgrade-backend/internal/blockchain/abi"
	"github.com/doge-verse/easy-upgrade-backend/internal/conf"
	"github.com/doge-verse/easy-upgrade-backend/internal/shared"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	MainnetEth     uint = 1
	MainnetPolygon uint = 137
	TestnetGoerli  uint = 5
	TestnetWallaby uint = 31415
	TestnetMumbai  uint = 80001
)

var (
	OwnershipEvent         = "OwnershipTransferred(address,address)"
	OwnershipEventTopicLen = 3
	ClientList             map[uint]*ethclient.Client
)

func Init() {
	// create client from the beginning
	clients := make(map[uint]*ethclient.Client)
	clientEth, err := ethclient.Dial(conf.GetRPC().MainnetEth)
	if err != nil {
		log.Fatalln("clientEth Dial err:", err)
	}
	clients[MainnetEth] = clientEth
	clientPolygon, err := ethclient.Dial(conf.GetRPC().MainnetPolygon)
	if err != nil {
		log.Fatalln("clientPolygon Dial err:", err)
	}
	clients[MainnetPolygon] = clientPolygon
	clientGoerli, err := ethclient.Dial(conf.GetRPC().TestnetGoerli)
	if err != nil {
		log.Fatalln("clientGoerli Dial err:", err)
	}
	clients[TestnetGoerli] = clientGoerli
	clientWallaby, err := ethclient.Dial(conf.GetRPC().TestnetWallaby)
	if err != nil {
		log.Fatalln("clientWallaby Dial err:", err)
	}
	clients[TestnetWallaby] = clientWallaby
	clientMumbai, err := ethclient.Dial(conf.GetRPC().TestnetWallaby)
	if err != nil {
		log.Fatalln("clientWallaby Dial err:", err)
	}
	clients[TestnetMumbai] = clientMumbai
	ClientList = clients
}

func GetOwnershipTransferredEvent(addr string, network uint) ([]models.ContractHistory, error) {
	client, ok := ClientList[network]
	if !ok {
		return nil, shared.ErrChainNotInit
	}

	number, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(addr)
	// hash event
	eventSignature := []byte(OwnershipEvent)
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
		var blockTime uint64
		blockInfo, err := client.BlockByNumber(context.Background(), big.NewInt(int64(eventLog.BlockNumber)))
		if err == nil {
			blockTime = blockInfo.Time()
		}
		results = append(results, models.ContractHistory{
			UpdateBlock:   eventLog.BlockNumber,
			UpdateTime:    blockTime,
			UpdateTX:      eventLog.TxHash.Hex(),
			PreviousOwner: common.HexToAddress(topics[1]).String(),
			NewOwner:      common.HexToAddress(topics[2]).String(),
		})
	}
	return results, errors.New("not found")
}

func GetProxyOwner(addr string, network uint) (string, error) {
	client, ok := ClientList[network]
	if !ok {
		return "", shared.ErrChainNotInit
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
