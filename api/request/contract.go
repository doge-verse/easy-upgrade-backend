package request

type Contract struct {
	UserID       uint   `json:"-"`
	ProxyAddress string `json:"proxyAddress"`
	ContractName string `json:"contractName"`
	Email        string `json:"email"`
	Network      uint   `json:"network"` // chain id
}
