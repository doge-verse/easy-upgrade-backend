package conf

import "github.com/spf13/viper"

type RPC struct {
	EthMainnet     string
	PolygonMainnet string
	GoerliTestnet  string
	WallabyTestnet string
}

func GetRPC() *RPC {
	return &RPC{
		EthMainnet:     viper.GetString("rpc.eth_mainnet"),
		PolygonMainnet: viper.GetString("rpc.polygon_mainnet"),
		GoerliTestnet:  viper.GetString("rpc.goerli_testnet"),
		WallabyTestnet: viper.GetString("rpc.wallaby_testnet"),
	}
}
