package contract

import (
	"github.com/doge-verse/easy-upgrade-backend/internal/sql"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Repo repoI

func Init(ethClient, polygonClient, goerliClient, fVMWallabyClient *ethclient.Client) {
	Repo = service{
		db:               sql.Db,
		ethClient:        ethClient,
		polygonClient:    polygonClient,
		goerliClient:     goerliClient,
		fVMWallabyClient: fVMWallabyClient,
	}
}
