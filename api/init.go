package api

import (
	"log"

	"github.com/doge-verse/easy-upgrade-backend/api/handler"
	"github.com/doge-verse/easy-upgrade-backend/internal/conf"
	"github.com/doge-verse/easy-upgrade-backend/internal/contract"
	"github.com/doge-verse/easy-upgrade-backend/internal/sql"
	"github.com/doge-verse/easy-upgrade-backend/internal/user"

	"github.com/gin-gonic/gin"
)

// Init .
func Init() {

	conf.Init()
	sql.Init()

	user.Init()
	contract.Init()

	gin.SetMode(conf.GetGin().Mode)
	r := gin.Default()
	handler.InitRouter(r)
	log.Fatal(r.Run(":8080"))
}
