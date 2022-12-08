package blockchain

import (
	"context"
	"errors"
	"math/big"

	"github.com/doge-verse/easy-upgrade-backend/internal/conf"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	EthMainnet     uint = 1
	PolygonMainnet uint = 137
)

var (
	ownershipEvent         = "OwnershipTransferred(address,address)"
	ownershipEventTopicLen = 3
)

func GetOwnershipTransferredEvent(addr string, network uint) ([]models.ContractHistory, error) {
	var rpcURL string
	switch network {
	case EthMainnet:
		rpcURL = conf.GetRPC().EthMainnt
	case PolygonMainnet:
		rpcURL = conf.GetRPC().PolygoMainnet
	}

	// create client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
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
