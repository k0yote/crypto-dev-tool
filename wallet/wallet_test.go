package wallet

import (
	"testing"

	"github.com/k0yote/backend-wallet/config"
	"github.com/stretchr/testify/require"
)

func TestWallet(t *testing.T) {

	config, err := config.LoadConfig()
	require.NoError(t, err)

	wallet, err := GenerateEmbeddedWallet(config, 32)
	require.NoError(t, err)
	require.NotEmpty(t, wallet.Hdwallet)
	require.NotEmpty(t, wallet.Account)
	require.NotEmpty(t, wallet.AccountID)
}
