package api

import (
	"log"

	"github.com/doge-verse/easy-upgrade-backend/api/handler"
	"github.com/doge-verse/easy-upgrade-backend/internal/blockchain"
	"github.com/doge-verse/easy-upgrade-backend/internal/cache"
	"github.com/doge-verse/easy-upgrade-backend/internal/conf"
	"github.com/doge-verse/easy-upgrade-backend/internal/contract"
	"github.com/doge-verse/easy-upgrade-backend/internal/sql"
	"github.com/doge-verse/easy-upgrade-backend/internal/subscriber"
	"github.com/doge-verse/easy-upgrade-backend/internal/user"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gin-gonic/gin"
)

// Init .
func Init() {
	conf.Init()

	// create client from the begining
	ethClient, err := ethclient.Dial(conf.GetRPC().EthMainnet)
	if err != nil {
		log.Fatalln("EthMainnet Dial err:", err)
	}
	polygonClient, err := ethclient.Dial(conf.GetRPC().PolygonMainnet)
	if err != nil {
		log.Fatalln("PolygonMainnet Dial err:", err)
	}
	goerliClient, err := ethclient.Dial(conf.GetRPC().GoerliTestnet)
	if err != nil {
		log.Fatalln("GoerilTestNet Dial err:", err)
	}
	fVMWallabyClient, err := ethclient.Dial(conf.GetRPC().WallabyTestnet)
	if err != nil {
		log.Fatalln("fVMWallabyClient Dial err:", err)
	}

	sql.Init()
	cache.Init()

	user.Init()
	contract.Init(ethClient, polygonClient, goerliClient, fVMWallabyClient)
	blockchain.Init(ethClient, polygonClient, goerliClient, fVMWallabyClient)
	subscriber.Init(ethClient, polygonClient, goerliClient, fVMWallabyClient)

	gin.SetMode(conf.GetGin().Mode)
	r := gin.Default()
	handler.InitRouter(r)
	log.Fatal(r.Run(":8080"))
}
