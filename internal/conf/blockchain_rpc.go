package conf

import "github.com/spf13/viper"

type RPC struct {
	MainnetEth     string
	MainnetPolygon string
	TestnetGoerli  string
	TestnetWallaby string
	TestnetMumbai  string
}

func GetRPC() *RPC {
	return &RPC{
		MainnetEth:     viper.GetString("rpc.eth_mainnet"),
		MainnetPolygon: viper.GetString("rpc.polygon_mainnet"),
		TestnetGoerli:  viper.GetString("rpc.goerli_testnet"),
		TestnetWallaby: viper.GetString("rpc.wallaby_testnet"),
		TestnetMumbai:  viper.GetString("rpc.mumbai_testnet"),
	}
}
