package request

type Contract struct {
	UserID       uint `json:"-"`
	ContractAddr string
	ContractName string
	Email        string
	Network      uint
}
