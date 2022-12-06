package conf

import "github.com/spf13/viper"

// Gin .
type Gin struct {
	Mode string
}

// GetGin .
func GetGin() *Gin {
	return &Gin{
		Mode: viper.GetString("gin.mode"),
	}
}
