package wallet

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/google/uuid"
	"github.com/k0yote/backend-wallet/config"
	hdwallet "github.com/k0yote/backend-wallet/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	Hdwallet  *hdwallet.Wallet
	Account   accounts.Account
	AccountID string
}

func GenerateEmbeddedWallet(config config.Config, size int) (Wallet, error) {

	entropy, err := entropy(config, int32(size))
	if err != nil {
		return Wallet{}, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return Wallet{}, err
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return Wallet{}, err
	}

	path := hdwallet.MustParseDerivationPath(hdwallet.DefaultBaseDerivationPath.String())

	account, err := wallet.Derive(path, false)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{
		Hdwallet:  wallet,
		Account:   account,
		AccountID: uuid.New().String(),
	}, nil
}
