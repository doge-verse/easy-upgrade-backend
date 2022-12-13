package subscriber

import (
	"log"

	"github.com/doge-verse/easy-upgrade-backend/internal/sql"
)

func Init() {
	s := subscriber{
		db: sql.Db,
	}

	contracts, err := s.SelectAllContract()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(contracts)
	// if err = s.SubscribeAllContract(contracts); err != nil {
	// 	log.Fatalln(err)
	// }
}
