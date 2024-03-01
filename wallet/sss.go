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

	return shareSecret(privateKey)
}

func GenerateSharedSecretBytes(privateKey []byte) ([]string, error) {
	return shareSecret(privateKey)
}

func CombineThresholdShares(strShares []string) ([]byte, error) {
	byteShares := [][]byte{}
	for _, strShare := range strShares {
		if strShare == "" {
			continue
		}

		byteShare, err := hex.DecodeString(strShare)
		if err != nil {
			return nil, err
		}

		byteShares = append(byteShares, byteShare)
	}

	return shamir.Combine(byteShares)
}

func shareSecret(privateKey []byte) ([]string, error) {
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
