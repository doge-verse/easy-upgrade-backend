package conf

import "github.com/spf13/viper"

// Database .
type Database struct {
	Dialect     string
	User        string
	Password    string
	Host        string
	Dbname      string
	Conf        string
	MaxIDLEConn int
	MaxOpenConn int
	Port        uint
	Debug       bool
	AutoMigrate bool
}

// GetDatabase .
func GetDatabase() *Database {
	return &Database{
		Dialect:     viper.GetString("database.dialect"),
		User:        viper.GetString("database.user"),
		Password:    viper.GetString("database.password"),
		Host:        viper.GetString("database.host"),
		Dbname:      viper.GetString("database.dbname"),
		Conf:        viper.GetString("database.config"),
		Port:        viper.GetUint("database.port"),
		MaxIDLEConn: viper.GetInt("database.max_idle_conns"),
		MaxOpenConn: viper.GetInt("database.max_open_conns"),
		Debug:       viper.GetBool("database.debug"),
		AutoMigrate: viper.GetBool("database.auto_migrate"),
	}
}
