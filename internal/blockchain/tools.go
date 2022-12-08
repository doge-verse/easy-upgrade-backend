package blockchain

import "github.com/ethereum/go-ethereum/crypto"

func CheckAddr(addr, signature, signData string) bool {
	data := []byte(signData)
	sign := []byte(signature)
	signer, err := crypto.SigToPub(data, sign)
	if err != nil {
		return false
	}
	address := crypto.PubkeyToAddress(*signer)
	return address.Hex() == addr
}
