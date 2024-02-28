package wallet

import (
	"testing"

	"github.com/k0yote/backend-wallet/config"
	"github.com/stretchr/testify/require"
)

func TestShamirShareSecret(t *testing.T) {

	config, err := config.LoadConfig()
	require.NoError(t, err)

	wallet, err := GenerateEmbeddedWallet(config, 32)
	require.NoError(t, err)

	shared, err := GenerateSharedSecret(wallet)
	require.NoError(t, err)
	require.NotEmpty(t, shared)
	require.Equal(t, 3, len(shared))
}
