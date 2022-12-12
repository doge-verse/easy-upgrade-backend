package blockchain

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func CheckAddr(addr, signature, signData string) bool {
	sig := hexutil.MustDecode(signature)

	msg := accounts.TextHash([]byte(signData))
	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	return strings.EqualFold(addr, recoveredAddr.Hex())
}
