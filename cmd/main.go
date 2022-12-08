package main

import(
	_ "github.com/doge-verse/easy-upgrade-backend/docs"
	"github.com/doge-verse/easy-upgrade-backend/api"
)

// @title easy-upgrade-backend Swagger
// @version 1.0
// @description
// @BasePath /api
// @query.collection.format multi
func main() {
	api.Init()
}
