package wallet

import (
	"encoding/hex"

	"github.com/hashicorp/vault/shamir"
)

func GenerateSharedSecret(wallet Wallet) ([]string, error) {
	// 2 of 3 secret sharing from private key which hdwallet derivepath index 0
	privateKey, err := wallet.Hdwallet.PrivateKeyBytes(wallet.Account)
	if err != nil {
		return nil, err
	}

	byteShares, err := shamir.Split(privateKey, 3, 2)
	if err != nil {
		return nil, err
	}

	var strShares []string
	for _, byteShare := range byteShares {
		strShares = append(strShares, hex.EncodeToString(byteShare))
	}

	return strShares, nil
}
