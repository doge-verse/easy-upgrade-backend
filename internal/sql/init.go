package sql

import (
	"fmt"
	"log"

	"github.com/doge-verse/easy-upgrade-backend/internal/conf"
	"github.com/doge-verse/easy-upgrade-backend/models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Init .
func Init() {
	dbConf := conf.GetDatabase()

	// dsn := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":" + cast.ToString(dbConf.Port) + ")/" + dbConf.Dbname + "?" + dbConf.Dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Dbname)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,  // DSN data source name
		DisableDatetimePrecision:  true, // Datetime precision is disabled. Databases before MySQL 5.6do not support it.
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		log.Fatal("init error", err)
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("sqlDB error", err)
		}
		sqlDB.SetMaxIdleConns(dbConf.MaxIDLEConn)
		sqlDB.SetMaxOpenConns(dbConf.MaxOpenConn)
		Db = db
	}

	if dbConf.Debug {
		Db = Db.Debug()
	}
	if err := Db.AutoMigrate(
		&models.User{},
		&models.Contract{},
		&models.ContractHistory{},
		&models.Notifier{},
	); err != nil {
		log.Fatal("AutoMigrate error", err)
	}
	log.Print("All table AutoMigrate finish.")
}
