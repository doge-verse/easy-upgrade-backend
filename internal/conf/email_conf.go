package conf

import "github.com/spf13/viper"

type EmailConf struct {
	AuthCode string
	From     string
	Username string
}

func GetEmailConf() *EmailConf {
	return &EmailConf{
		AuthCode: viper.GetString("email.authCode"),
		From:     viper.GetString("email.from"),
		Username: viper.GetString("email.username"),
	}
}
