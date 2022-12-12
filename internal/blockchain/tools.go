package blockchain

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func CheckAddr(addr, signature, signData string) bool {
	hash := crypto.Keccak256Hash([]byte(signData))
	sig := hexutil.MustDecode(signature)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false
	}
	address := crypto.PubkeyToAddress(*pubKey)
	return address.Hex() == addr
}
