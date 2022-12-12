package conf

import "github.com/spf13/viper"

type RPC struct {
	EthMainnt     string
	PolygoMainnet string
	GoerliTestnet string
}

func GetRPC() *RPC {
	return &RPC{
		EthMainnt:     viper.GetString("rpc.eth_mainnet"),
		PolygoMainnet: viper.GetString("rpc.polygo_mainnet"),
		GoerliTestnet: viper.GetString("rpc.goerli_testnet"),
	}
}
