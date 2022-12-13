package subscriber

import (
	"log"

	"github.com/doge-verse/easy-upgrade-backend/internal/sql"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Init(ethClient, polygonClient, goerliClient, fVMWallabyClient *ethclient.Client) {
	s := &Subscriber{
		Db:                   sql.Db,
		EthMainnetClient:     ethClient,
		PolygonMainnetClient: polygonClient,
		GoerliClinet:         goerliClient,
		FVMWallabyClient:     fVMWallabyClient,
	}

	contracts, err := s.SelectAllContract()
	if err != nil {
		log.Fatalln(err)
	}

	s.SubscribeAllContract(contracts)
}
