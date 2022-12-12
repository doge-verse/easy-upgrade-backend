package request

type Contract struct {
	UserID        uint `json:"-"`
	ContractAddr  string
	ContractAdmin string
	ContractName  string
	Email         string
	Network       uint
}
