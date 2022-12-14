package conf

import (
	"fmt"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/spf13/viper"
)

// GetSessionStore .
func GetSessionStore() cookie.Store {
	size := viper.GetInt("session.size")
	network := viper.GetString("session.network")
	address := viper.GetString("session.address")
	password := viper.GetString("session.password")
	keyPairs := viper.GetString("session.keyPairs")
	store, err := redis.NewStore(size, network, address, password, []byte(keyPairs))
	if err != nil {
		panic(fmt.Sprintf("init session cookie store fail:size:%d; network:%s;address:%s;password:%s;keyPairs:%s;", size, network, address, password, keyPairs))
	}
	return store
}

func GetTokenExpire() int {
	return viper.GetInt("session.expires-time")
}
