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
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Init .
func Init() {
	logInit()
	conf.Init()

	sql.Init()
	cache.Init()

	user.Init()
	blockchain.Init()
	contract.Init()
	subscriber.Init()

	gin.SetMode(conf.GetGin().Mode)
	r := gin.Default()
	handler.InitRouter(r)
	log.Fatal(r.Run(":8080"))
}

func logInit() {
	logrus.SetReportCaller(true)
}
