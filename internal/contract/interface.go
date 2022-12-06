package contract

import "github.com/doge-verse/easy-upgrade-backend/pkg"

type repoI interface {
	AddContract(contract *pkg.Contract) (*pkg.Contract, error)

	GetUserContractArr(userID uint) ([]pkg.Contract, error)

	GetContractRecord(addr string) ([]pkg.ContractUpdateRecord, error)
}
